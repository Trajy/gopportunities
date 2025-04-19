package config

var (
	logger *Logger
)

func GetLogger() *Logger {
	if logger == nil {
		logger = newLogger()
	}
	return logger
}
