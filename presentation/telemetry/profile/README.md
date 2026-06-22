# Profiling Server

An isolated HTTP server exposing Go's `pprof` profiling endpoints and a plain-text runtime stats snapshot. Runs on a dedicated port, separate from the main REST server.

> **Never expose the telemetry port publicly.** Bind it to `localhost` or restrict it at the network level.


## Setup

### 1. Config

Add following to `infra/config.go`

```go
type Telemetry struct {
	Profile ProfilerConfig `mapstructure:"profiler"`
}

type ProfilerConfig struct {
	// Port is the port the pprof/metrics HTTP server listens on.
	// Keep this on a non-public port (e.g. 6060) and never expose it externally.
	Port int `yaml:"port" validate:"min=80,max=65535"`
	// Enabled allows the telemetry server to be turned off entirely (e.g. in prod
	// environments where you rely on an external profiler instead).
	Enabled bool `yaml:"enabled"`
	// ReadTimeoutSeconds is the HTTP read timeout for the telemetry server.
	// Keep it high enough for long-running pprof captures (30s CPU profiles, etc.)
	ReadTimeout int `yaml:"readtimeout"`
}
```

reference `Telemetry` in master config structure.
```go
// Config is the master config struct that holds all other config structs.
type Config struct {
	App        AppConfig         `mapstructure:"app"`
	Log        loggerjson.Config `mapstructure:"log"`
	Telemetry  Telemetry         `mapstructure:"telemetry"` // <-- add
}
```

Add following to `config.yaml` to enable profiler.

```yaml
telemetry:
  profiler:
    enabled: true
    port: 6060
    readtimeout: 60
```

| Field | Default | Description |
|---|---|---|
| `enabled` | `true` | Set to `false` to disable entirely — lifecycle calls are still safe (no-op) |
| `port` | `6060` | Port the profiling server listens on |
| `readtimeout` | `60` | HTTP read timeout — keep above your longest expected CPU capture |

### 2. Lifecycle in `main.go`

```go
// start profiling server
profSrv, err := profile.NewServer(cfg.Telemetry.Profile, ctr)
if err != nil {
  msgNPanic(lctx, ctr.Logger, fmt.Errorf("main: error creating profiling server, %w", err).Error())
}
if err := profSrv.Start(); err != nil {
  msgNPanic(lctx, ctr.Logger, fmt.Errorf("main: error starting profiling server, %w", err).Error())
}
defer func() {
  if err := profSrv.Stop(); err != nil {
    msgNPanic(lctx, ctr.Logger, fmt.Errorf("main: error stopping profiling server, %w", err).Error())
  }
}()
```

`Start()` binds the port eagerly — a port conflict surfaces immediately as an error rather than being swallowed in a goroutine. `Stop()` waits up to 65 seconds for in-flight captures to complete before closing.

## Endpoints

### `GET /debug/pprof/`
Index page listing all available profiles. Start here.
```bash
curl http://localhost:6060/debug/pprof/
```

### `GET /debug/pprof/profile?seconds=N`
CPU profile. Samples CPU usage for `N` seconds (default 30) and returns a profile binary.
```bash
# Capture and open an interactive flame graph
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile?seconds=30
```
Wide bars in the flame graph = hot functions consuming CPU.

### `GET /debug/pprof/heap`
Snapshot of live heap allocations. Use this to find memory leaks.

```bash
# Diff two snapshots taken minutes apart to find what's growing
go tool pprof -http=:8080 -diff_base=heap1.pb.gz http://localhost:6060/debug/pprof/heap
```

### `GET /debug/pprof/goroutine`
Stack traces of all current goroutines. Use this to detect goroutine leaks.

```bash
# Plain text dump — look for goroutines stuck in the same place
curl http://localhost:6060/debug/pprof/goroutine?debug=2
```

A growing goroutine count over time = leak.

### `GET /debug/pprof/allocs`
Allocation profile — samples all past allocations, not just live ones. Use this to diagnose GC pressure.

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/allocs
```

### `GET /debug/pprof/block`
Shows where goroutines block on synchronisation primitives (channels, mutexes). Requires `runtime.SetBlockProfileRate(N)` to be called at startup.

### `GET /debug/pprof/mutex`
Shows contended mutex holders. Requires `runtime.SetMutexProfileFraction(N)` to be called at startup.

### `GET /debug/pprof/trace?seconds=N`
Execution trace — goroutine scheduling, GC events, syscalls. More detailed than a CPU profile.
```bash
curl -o trace.out http://localhost:6060/debug/pprof/trace?seconds=5
go tool trace trace.out
```

### `GET /debug/stats`
Plain-text runtime snapshot. No tooling required — useful for a quick eyeball.

```bash
curl http://localhost:6060/debug/stats
```

```
goroutines     : 14
num_cpu        : 8
gomaxprocs     : 8
---
heap_alloc_mb  : 4.21
heap_sys_mb    : 14.50
heap_inuse_mb  : 5.00
heap_idle_mb   : 9.50
heap_released_mb: 8.25
stack_inuse_mb : 0.34
sys_total_mb   : 18.43
---
total_alloc_mb : 120.77
mallocs        : 981234
frees          : 976110
---
gc_runs        : 42
gc_forced      : 0
gc_last_pause  : 381µs
next_gc_mb     : 8.00
```

Key signals to watch:

| Metric | What it indicates |
|---|---|
| `heap_alloc_mb` growing without bound | Memory leak |
| `goroutines` growing without bound | Goroutine leak |
| `gc_last_pause` in the tens of ms | GC pressure — likely over-allocating |
| `heap_released_mb` stays near 0 | GC not returning memory to OS |

`mallocs`, `frees`, and `total_alloc_mb` are lifetime counters — they only ever increase.
A rising `mallocs` is normal. The meaningful signal is `mallocs - frees` (live object count);
if that delta widens steadily over time, there is a leak.

## Common Workflows

**Investigate a CPU spike**
```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile?seconds=30
```

**Investigate a memory leak**
```bash
# Take two snapshots a few minutes apart, then diff them
curl -o heap1.pb.gz http://localhost:6060/debug/pprof/heap
# ... wait ...
go tool pprof -http=:8080 -diff_base=heap1.pb.gz http://localhost:6060/debug/pprof/heap
```

**Investigate GC-related CPU spikes**
```bash
GODEBUG=gctrace=1 ./your-service
# High frequency + long pause times = GC is thrashing
```
