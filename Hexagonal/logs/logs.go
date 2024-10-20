package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1)) //ข้ามไปนึงครั้ง เพราะให้ main ได้ใช้
	if err != nil {
		panic(err)
	}
}
// fields ...zap.Field คือ variadic parameter 
// หมายความว่าคุณสามารถส่งค่า zap.Field เข้าไปหลายๆ ค่า หรือไม่ส่งก็ได้
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// คำสั่ง .(type) ใช้ในการทำ type assertion ซึ่งมักจะใช้ในโครงสร้างการควบคุมที่เรียกว่า type switch 
// เพื่อระบุชนิดของข้อมูลที่จัดเก็บอยู่ในตัวแปรอินเตอร์เฟซ (interface{})
func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}