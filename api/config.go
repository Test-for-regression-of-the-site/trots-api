package api

// Config stores the API settings.
type Config struct {
	// Addr is the serving address.
	Addr string
	// Logger creates the a named logger.
	// Name can be ignored or prefix the log entries.
	Logger func(name string) Logger
}

// Option sets config fields.
type Option func(cfg *Config) error

// WithAddr sets the serving address of the API.
func WithAddr(addr string) Option {
	return func(cfg *Config) error {
		cfg.Addr = addr
		return nil
	}
}
