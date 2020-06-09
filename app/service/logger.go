package service

// Logger is an adapter for zap
type Logger struct {}

// ProvideLogger is the provider for the Logger Service
func ProvideLogger() (*Logger, error) {
	return &Logger{}, nil
}
