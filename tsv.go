package tsv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"
)

type TsvLog struct {
	tsvFile    *os.File
	writer     *csv.Writer
	headers    []string
	fileName   string
	timeFormat string
}

func Create(headers []string, path string, format string) *TsvLog {
	tsvFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
	}

	writer := csv.NewWriter(tsvFile)

	writer.Comma = '\t' // Use tab-delimited instead of comma <---- here!

	log := TsvLog{
		tsvFile:    tsvFile,
		headers:    append([]string{"ts"}, headers...),
		writer:     writer,
		fileName:   path,
		timeFormat: format}

	if len(headers) > 0 {
		writer.Write(log.headers)
	}

	writer.Flush()

	return &log
}

func (log *TsvLog) Add(data []string) error {
	var err error
	if len(data) != len(log.headers)-1 {
		err = errors.New("csv data length doesnt match header length")
	} else {
		log.writer.Write(append([]string{time.Now().Format(log.timeFormat)}, data...))
		log.writer.Flush()
	}
	return err
}

func (log *TsvLog) Read() ([]string, [][]string) {
	file, err := os.Open(log.fileName)
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	records, err2 := reader.ReadAll()
	if err != nil {
		fmt.Println(err2)
	}

	return records[0], records[1:]
}

func (log *TsvLog) Close() {
	log.writer.Flush()
	log.tsvFile.Close()
}

func (log *TsvLog) Delete() {
	log.writer.Flush()
	log.tsvFile.Close()
	os.Remove(log.fileName)
}
