package main

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
)

// 参考 https://zhuanlan.zhihu.com/p/88856378

var logger *zap.Logger

//var sugarLogger *zap.SugaredLogger

func InitLogger() {
	//logger, _ = zap.NewProduction()
	//sugarLogger = logger.Sugar()
	//zap.NewDevelopment()
	//zap.NewExample()

	// 定制logger
	// zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel
	core := zapcore.NewCore(getEncoder(), getLogWriter(), zap.DebugLevel)
	//logger = zap.New(core)
	logger = zap.New(core, zap.AddCaller()) // 在zap.New(..)函数中添加一个Option，添加将调用函数信息记录到日志中的功能
}

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(config)
}

func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.Create("./test.log")
	//return zapcore.AddSync(file)

	lumberLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,
		MaxAge:     30,
		MaxBackups: 5,
		LocalTime:  false,
		Compress:   false,
	}
	wrapper := &MyWrapper{w: lumberLogger}
	return zapcore.AddSync(wrapper)
}

func main() {
	InitLogger()
	defer logger.Sync()
	logger.Info("Test ..")
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("error fetching url...", zap.String("url", url), zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		//sugarLogger.Info("Success..",
		//	zap.String("statusCode", resp.Status),
		//	zap.String("url", url))
		resp.Body.Close()
	}
}

type MyWrapper struct {
	w io.Writer
}

func (myW *MyWrapper) Write(p []byte) (n int, err error) {
	fmt.Println("pre do same thing...", string(p))
	write, err := myW.w.Write(p)
	fmt.Println("behind do same thing...")
	return write, err
}
