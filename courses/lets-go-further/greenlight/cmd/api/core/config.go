package core

import "flag"

type databaseConfig struct {
	Dsn          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type rateLimiterConfig struct {
	Rps     float64
	Max     int
	Enabled bool
}

type Config struct {
	Port        int
	Env         string
	Db          databaseConfig
	RateLimiter rateLimiterConfig
}

func NewConfig() *Config {

	var cfg Config

	// Port
	flag.IntVar(
		&cfg.Port,
		"port",
		envInt("GREENLIGHT_PORT", 4000),
		"API server port",
	)

	// Environment
	flag.StringVar(
		&cfg.Env,
		"env",
		env("GREENLIGHT_ENV", "development"),
		"Environment (development|staging|production)",
	)

	// Database Source Name
	flag.StringVar(
		&cfg.Db.Dsn,
		"db-dsn",
		env("GREENLIGHT_DB_DSN", "postgres://greenlight:greenlight@localhost:5432/greenlight"),
		"PostgreSQL DSN",
	)

	// Database max open connections
	flag.IntVar(
		&cfg.Db.MaxOpenConns,
		"db-max-open-conns",
		envInt("GREENLIGHT_DB_MAX_OPEN_CONNS", 25),
		"PostgreSQL max open connections",
	)

	// Database max idle connections
	flag.IntVar(
		&cfg.Db.MaxIdleConns,
		"db-max-idle-conns",
		envInt("GREENLIGHT_DB_MAX_IDLE_CONNS", 25),
		"PostgreSQL max idle connections",
	)

	// Database max idle time
	flag.StringVar(
		&cfg.Db.MaxIdleTime,
		"db-max-idle-time",
		env("GREENLIGHT_DB_MAX_IDLE_TIME", "15m"),
		"PostgreSQL max connection idle time",
	)

	// Is rate limiter enabled?
	flag.BoolVar(
		&cfg.RateLimiter.Enabled,
		"limiter-enabled",
		true,
		"Rate limiter enabled",
	)

	// Rate limiter requests/second allowed
	flag.Float64Var(
		&cfg.RateLimiter.Rps,
		"limiter-rps",
		envFloat("GREENLIGHT_RATE_LIMIT_RPS", 2.0),
		"Rate limiter requests/second to be allowed",
	)

	// Rate limiter maximum requests/second allowed in bursts
	flag.IntVar(
		&cfg.RateLimiter.Max,
		"limiter-max",
		envInt("GREENLIGHT_RATE_LIMIT_MAX", 4),
		"Rate limiter maximum requests/second to be allowed in bursts",
	)

	flag.Parse()

	return &cfg
}
