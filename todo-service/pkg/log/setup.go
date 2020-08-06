package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func init() {
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zap.CombineWriteSyncers(zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}
