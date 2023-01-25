package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	conf := zap.NewProductionConfig()

	// 可以把输出方式改为控制台编码, 更容易阅读
	conf.Encoding = "console"
	// 时间格式自定义
	conf.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
	}
	// 打印路径自定义
	conf.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + caller.TrimmedPath() + "]")
	}
	// 级别显示自定义
	conf.EncoderConfig.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + level.String() + "]")
	}

	logger, _ := conf.Build()
	logger.Info("service start")

	logger.Info("info msg",
		zap.String("name", "掘金"),
		zap.Int("num", 3),
		zap.Duration("timer", time.Minute),
	)
}
