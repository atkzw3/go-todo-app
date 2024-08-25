package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ログファイルの読み込み
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// ログの書き込み先の指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// ログのフォーマット指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// 出力先を指定
	log.SetOutput(multiLogFile)
}
