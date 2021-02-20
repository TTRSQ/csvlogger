package csvlogger

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

type csvLogger struct {
	basePath  string
	fileNamme string
}

// NewLogger .. Create logger.
func NewLogger(basePath, fileNamme string) (*csvLogger, error) {
	if basePath == "" {
		basePath = "./"
	}
	if fileNamme == "" {
		return nil, errors.New("fileNamme is required.")
	}

	return &csvLogger{
		basePath:  basePath,
		fileNamme: fileNamme,
	}, nil
}

// Add .. Add row or Create file and write header if not found.
func (c *csvLogger) Add(data interface{}) {
	now := time.Now()
	fullPath := fmt.Sprintf("%s/%s.%s.log", c.basePath, c.fileNamme, now.Format("20060102"))

	head, body := structToList(data)

	csvArray := [][]string{}
	if !fileExists(fullPath) {
		csvArray = append(csvArray, head)
	}

	csvArray = append(csvArray, body)

	f, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(csvArray)

	if err := w.Error(); err != nil {
		log.Println(err)
	}
}

func structToList(data interface{}) ([]string, []string) {
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	head := []string{}
	body := []string{}

	for i := 0; i < size; i++ {
		value := elem.Field(i).Interface()
		field := elem.Type().Field(i).Name
		head = append(head, field)
		body = append(body, fmt.Sprint(value))
	}

	return head, body
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}
