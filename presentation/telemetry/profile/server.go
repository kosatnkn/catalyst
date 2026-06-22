package profile

import (
	"context"
	"fmt"
	"net/http"
	"time"

	// Blank import registers all pprof HTTP handlers on the default ServeMux:
	//   /debug/pprof/
	//   /debug/pprof/cmdline
	//   /debug/pprof/profile   (CPU — use ?seconds=N)
	//   /debug/pprof/symbol
	//   /debug/pprof/trace
	//   /debug/pprof/heap
	//   /debug/pprof/goroutine
	//   /debug/pprof/allocs
	//   /debug/pprof/block
	//   /debug/pprof/mutex
	_ "net/http/pprof"
	"runtime"

	"github.com/kosatnkn/catalyst/v3/infra"
)

// stopTimeout is how long Stop() waits for in-flight requests (e.g. a 30s CPU
// profile) to finish before forcibly closing the server.
const stopTimeout = 65 * time.Second

// server is the concrete implementation of Server.
type server struct {
	cfg     infra.ProfilerConfig
	log     infra.Logger
	httpSrv *http.Server
}

// NewServer creates a new telemetry server.
// If cfg.Enabled is false it returns a no-op server so callers never need to
// branch — Start/Stop are always safe to call.
func NewServer(cfg infra.ProfilerConfig, ctr *infra.Container) (Server, error) {
	if !cfg.Enabled {
		return &nopServer{}, nil
	}
	if cfg.Port <= 0 {
		return nil, fmt.Errorf("telemetry: invalid port %d", cfg.Port)
	}

	readTimeout := time.Duration(cfg.ReadTimeout) * time.Second
	if readTimeout <= 0 {
		readTimeout = 65 * time.Second // safe default for 60s CPU profiles
	}

	mux := http.NewServeMux()
	registerHandlers(mux)

	httpSrv := &http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.Port),
		Handler:     mux,
		ReadTimeout: readTimeout,
		// No WriteTimeout: pprof profile endpoints stream for N seconds, a
		// write timeout would cancel them mid-capture.
	}

	return &server{
		cfg:     cfg,
		log:     ctr.Logger,
		httpSrv: httpSrv,
	}, nil
}

// Start binds the port and starts serving in the background.
// It blocks briefly to verify the listener is healthy before returning.
func (s *server) Start() error {
	ctx := context.Background()

	s.log.Info(ctx, fmt.Sprintf("starting profiling server on port %d", s.cfg.Port))
	s.log.Debug(ctx, fmt.Sprintf("config values: %+v", s.cfg))
	// Bind the listener eagerly so we can return a port-conflict error
	// immediately rather than swallowing it in a goroutine.
	ln, err := newListener(s.httpSrv.Addr)
	if err != nil {
		return fmt.Errorf("profile: Start, failed to bind to port %s, %w", s.httpSrv.Addr, err)
	}

	go func() {
		if err := s.httpSrv.Serve(ln); err != nil && err != http.ErrServerClosed {
			// Nothing useful to do here other than surface it; the server is
			// typically stopped via Stop(), which causes ErrServerClosed.
			_ = err
		}
	}()

	return nil
}

// Stop gracefully shuts down the HTTP server, giving in-flight profiling
// requests up to stopTimeout to complete.
func (s *server) Stop() error {
	s.log.Info(context.Background(), "stopping profiling server")

	ctx, cancel := context.WithTimeout(context.Background(), stopTimeout)
	defer cancel()

	if err := s.httpSrv.Shutdown(ctx); err != nil {
		return fmt.Errorf("profile: Stop, shutdown error, %w", err)
	}
	return nil
}

// registerHandlers attaches all telemetry endpoints to mux.
// pprof handlers are already registered on http.DefaultServeMux via the blank
// import above; we re-register them explicitly on our own mux so we're not
// sharing DefaultServeMux with anything else in the binary.
func registerHandlers(mux *http.ServeMux) {
	// pprof — re-register from DefaultServeMux so our isolated mux gets them.
	// The blank import populates DefaultServeMux; we proxy through to keep our
	// mux self-contained.
	mux.Handle("/debug/pprof/", http.DefaultServeMux)

	// Runtime stats endpoint — human-readable, no tooling required.
	mux.HandleFunc("/debug/stats", runtimeStatsHandler)
}

// runtimeStatsHandler writes a plain-text snapshot of runtime memory and
// goroutine metrics. Useful for a quick eyeball without running go tool pprof.
//
// Example: curl http://localhost:6060/debug/stats
func runtimeStatsHandler(w http.ResponseWriter, _ *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "goroutines     : %d\n", runtime.NumGoroutine())
	fmt.Fprintf(w, "num_cpu        : %d\n", runtime.NumCPU())
	fmt.Fprintf(w, "gomaxprocs     : %d\n", runtime.GOMAXPROCS(0))
	fmt.Fprintln(w, "---")
	fmt.Fprintf(w, "heap_alloc_mb  : %.2f\n", mbf(m.HeapAlloc))     // live heap objects
	fmt.Fprintf(w, "heap_sys_mb    : %.2f\n", mbf(m.HeapSys))       // reserved from OS
	fmt.Fprintf(w, "heap_inuse_mb  : %.2f\n", mbf(m.HeapInuse))     // spans with objects
	fmt.Fprintf(w, "heap_idle_mb   : %.2f\n", mbf(m.HeapIdle))      // spans returned/waiting
	fmt.Fprintf(w, "heap_released_mb: %.2f\n", mbf(m.HeapReleased)) // returned to OS
	fmt.Fprintf(w, "stack_inuse_mb : %.2f\n", mbf(m.StackInuse))
	fmt.Fprintf(w, "sys_total_mb   : %.2f\n", mbf(m.Sys))
	fmt.Fprintln(w, "---")
	fmt.Fprintf(w, "total_alloc_mb : %.2f\n", mbf(m.TotalAlloc)) // cumulative, ever-growing
	fmt.Fprintf(w, "mallocs        : %d\n", m.Mallocs)
	fmt.Fprintf(w, "frees          : %d\n", m.Frees)
	fmt.Fprintln(w, "---")
	fmt.Fprintf(w, "gc_runs        : %d\n", m.NumGC)
	fmt.Fprintf(w, "gc_forced      : %d\n", m.NumForcedGC)
	if m.NumGC > 0 {
		lastPause := time.Duration(m.PauseNs[(m.NumGC+255)%256])
		fmt.Fprintf(w, "gc_last_pause  : %s\n", lastPause)
	}
	fmt.Fprintf(w, "next_gc_mb     : %.2f\n", mbf(m.NextGC))
}

func mbf(b uint64) float64 { return float64(b) / 1024 / 1024 }
