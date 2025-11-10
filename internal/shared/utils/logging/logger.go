package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func InitLogger(level string, isProd bool, logFilePath string) (*zap.Logger, error) {
	parsedLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		parsedLevel = zapcore.InfoLevel
	}

	devEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	prodEncoderCfg := zap.NewProductionEncoderConfig()
	prodEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	prodEncoder := zapcore.NewJSONEncoder(prodEncoderCfg)

	if !isProd {
		core := zapcore.NewCore(devEncoder, zapcore.AddSync(os.Stdout), parsedLevel)
		return zap.New(core, zap.AddCaller()), nil
	}

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    20, // MB
		MaxBackups: 7,
		MaxAge:     14, // days
		Compress:   true,
	})

	stdoutWriter := zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(prodEncoder, stdoutWriter, parsedLevel),
		zapcore.NewCore(prodEncoder, fileWriter, parsedLevel),
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)), nil
}
