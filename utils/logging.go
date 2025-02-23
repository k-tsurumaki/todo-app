package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // 適当な権限を与えて読み込み
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile) // 出力先を標準出力とログファイルに指定
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile) // ログのフォーマットを指定
	log.SetOutput(multiLogFile)
}