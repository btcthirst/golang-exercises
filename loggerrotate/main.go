package main

import (
	"fmt"
	"os"
	"time"
)

type RotatingLogger struct {
	basePath string
	maxSize  int64
	file     *os.File
}

func NewRotatingLogger(path string, maxSize int64) (*RotatingLogger, error) {
	logger := &RotatingLogger{
		basePath: path,
		maxSize:  maxSize,
	}
	err := logger.openFile()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func (r *RotatingLogger) openFile() error {
	file, err := os.OpenFile(r.basePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	r.file = file
	return nil
}

func (r *RotatingLogger) rotate() error {
	r.file.Close()

	// Генеруємо нове ім'я для лог-файлу з часовою міткою
	timestamp := time.Now().Format("20060102_150405")
	newName := fmt.Sprintf("%s.%s", r.basePath, timestamp)

	err := os.Rename(r.basePath, newName)
	if err != nil {
		return err
	}
	return r.openFile()
}

func (r *RotatingLogger) Write(p []byte) (n int, err error) {
	// Перевірка розміру
	info, err := r.file.Stat()
	if err != nil {
		return 0, err
	}

	if info.Size()+int64(len(p)) > r.maxSize {
		err = r.rotate()
		if err != nil {
			return 0, err
		}
	}

	return r.file.Write(p)
}

func (r *RotatingLogger) Close() error {
	return r.file.Close()
}

func main() {
	logger, err := NewRotatingLogger("app.log", 1024*10) // 10 KB для тесту
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	for i := 0; i < 1000; i++ {
		msg := fmt.Sprintf("Запис %d: Це приклад логування\n", i)
		_, err := logger.Write([]byte(msg))
		if err != nil {
			fmt.Println("Помилка запису:", err)
			break
		}
	}
}
