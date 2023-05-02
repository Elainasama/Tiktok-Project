package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func InitLogger() {
	//var err error
	//logger, err = zap.NewProduction()
	//sugar = logger.Sugar()
	//if err != nil {
	//	panic(err)
	//}
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugar = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

func Sync() {
	logger.Sync()
}

func Infof(s string, v ...interface{}) {
	sugar.Infof(s, v...)
}

func Infow(s string, v ...interface{}) {
	sugar.Infow(s, v...)
}

func Info(v ...interface{}) {
	sugar.Info(v...)
}

func Debugf(s string, v ...interface{}) {
	sugar.Debugf(s, v...)
}

func Debugw(s string, v ...interface{}) {
	sugar.Debugw(s, v...)
}

func Debug(v ...interface{}) {
	sugar.Debug(v...)
}

func Errorf(s string, v ...interface{}) {
	sugar.Errorf(s, v...)
}

func Errorw(s string, v ...interface{}) {
	sugar.Errorw(s, v...)
}

func Error(v ...interface{}) {
	sugar.Error(v...)
}

func Fatalf(s string, v ...interface{}) {
	sugar.Fatalf(s, v...)
}

func Fatalw(s string, v ...interface{}) {
	sugar.Fatalw(s, v...)
}

func Fatal(v ...interface{}) {
	sugar.Error(v...)
}
