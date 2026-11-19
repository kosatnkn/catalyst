package profile

// nopServer is returned by NewServer when cfg.Enabled == false.
// It satisfies the Server interface with safe no-ops so main.go never needs
// to branch on whether telemetry is enabled.
type nopServer struct{}

func (n *nopServer) Start() error { return nil }
func (n *nopServer) Stop() error  { return nil }

// Err returns a nil channel, which never fires in a select.
func (n *nopServer) Err() <-chan error { return nil }
