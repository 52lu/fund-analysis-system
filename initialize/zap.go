/**
 * @Author Mr.LiuQH
 * @Description 初始化zap日志
 * @Date 2021/7/5 5:54 下午
 **/
package initialize

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/utils"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	// 日志输出格式
	outJson = "json"
)

// 初始化Logger
func initLogger() {
	logConfig := global.GvaConfig.Log
	// 判断日志目录是否存在
	if exist, _ := utils.DirExist(getLogPath()); !exist {
		_ = utils.CreateDir(getLogPath())
	}
	// 设置输出格式
	var encoder zapcore.Encoder
	if logConfig.OutFormat == outJson {
		encoder = zapcore.NewJSONEncoder(getEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(getEncoderConfig())
	}
	// 设置日志文件切割
	writeSyncer := zapcore.AddSync(getLumberjackWriteSyncer())
	// 创建NewCore
	zapCore := zapcore.NewCore(encoder, writeSyncer, getLevel())
	// 创建logger
	logger := zap.New(zapCore,zap.AddCaller(),zap.AddStacktrace(zap.ErrorLevel))
	defer logger.Sync()
	// 赋值给全局变量
	global.GvaLogger = logger
}

// 获取最低记录日志级别
func getLevel() zapcore.Level {
	levelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	if level, ok := levelMap[global.GvaConfig.Log.Level]; ok {
		return level
	}
	// 默认info级别
	return zapcore.InfoLevel
}

// 自定义日志输出字段
func getEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     getEncodeTime, // 自定义输出时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return config
}

// 定义日志输出时间格式
func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// 获取文件切割和归档配置信息
func getLumberjackWriteSyncer() zapcore.WriteSyncer {
	lumberjackConfig := global.GvaConfig.Log.LumberJack
	lumberjackLogger := &lumberjack.Logger{
		Filename:   getLogFile(),                //日志文件
		MaxSize:    lumberjackConfig.MaxSize,    //单文件最大容量(单位MB)
		MaxBackups: lumberjackConfig.MaxBackups, //保留旧文件的最大数量
		MaxAge:     lumberjackConfig.MaxAge,     // 旧文件最多保存几天
		Compress:   lumberjackConfig.Compress,   // 是否压缩/归档旧文件
	}
	// 设置日志文件切割
	return zapcore.AddSync(lumberjackLogger)
}

// 获取日志文件名
func getLogFile() string {
	fileFormat := time.Now().Format(global.GvaConfig.Log.FileFormat)
	fileName := strings.Join([]string{
		global.GvaConfig.Log.FilePrefix,
		fileFormat,
		"log"}, ".")
	return path.Join(getLogPath(), fileName)
}
// 获取日志文件目录
func getLogPath() string {
	_, file, _, _ := runtime.Caller(1)
	return path.Dir(file) + "/../"+global.GvaConfig.Log.Path
}