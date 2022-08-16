package config

// Config ...
//go:generate go run ../../../newc --value
type Config struct {
	Debug bool
}

// DebugConfig ...
type DebugConfig struct {
	Debug bool
}
