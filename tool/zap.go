package tool

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var sugaredLog *zap.SugaredLogger
var zapLog *zap.Logger

func GetSugaredLog() *zap.SugaredLogger {
	return sugaredLog
}
func GetZapLog() *zap.Logger {
	return zapLog
}

func Logger() (err error) {
	var coreArr []zapcore.Core

	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()            //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder       //显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	//日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/info.log", //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    1,                //文件大小限制,单位MB
		MaxBackups: 5,                //最大保留日志文件数量
		MaxAge:     30,               //日志文件保留天数
		Compress:   false,            //是否压缩处理
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	//error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/error.log", //日志文件存放目录
		MaxSize:    1,                 //文件大小限制,单位MB
		MaxBackups: 5,                 //最大保留日志文件数量
		MaxAge:     30,                //日志文件保留天数
		Compress:   false,             //是否压缩处理
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	log := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) //zap.AddCaller()为显示文件名和行号，可省略
	zapLog = log
	sugaredLog = zapLog.Sugar()
	return err
}

func SugaredDebug(args ...interface{}) {
	sugaredLog.Debug(args...)
}

func SugaredDebugf(template string, args ...interface{}) {
	sugaredLog.Debugf(template, args...)
}

func SugaredInfo(args ...interface{}) {
	sugaredLog.Info(args...)
}

func SugaredInfof(template string, args ...interface{}) {
	sugaredLog.Infof(template, args...)
}

func SugaredWarn(args ...interface{}) {
	sugaredLog.Warn(args...)
}

func SugaredWarnf(template string, args ...interface{}) {
	sugaredLog.Warnf(template, args...)
}

func SugaredError(args ...interface{}) {
	sugaredLog.Error(args...)
}

func SugaredErrorf(template string, args ...interface{}) {
	sugaredLog.Errorf(template, args...)
}

func SugaredDPanic(args ...interface{}) {
	sugaredLog.DPanic(args...)
}

func SugaredDPanicf(template string, args ...interface{}) {
	sugaredLog.DPanicf(template, args...)
}

func SugaredPanic(args ...interface{}) {
	sugaredLog.Panic(args...)
}

func SugaredPanicf(template string, args ...interface{}) {
	sugaredLog.Panicf(template, args...)
}

func SugaredFatal(args ...interface{}) {
	sugaredLog.Fatal(args...)
}

func SugaredFatalf(template string, args ...interface{}) {
	sugaredLog.Fatalf(template, args...)
}

func Debug(msg string, fields ...zap.Field) {
	zapLog.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zapLog.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zapLog.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zapLog.Error(msg, fields...)
}
