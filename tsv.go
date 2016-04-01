package tsv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"
)

type TsvLog struct {
	headers    []string
	fileName   string
	timeFormat string
	writer     *csv.Writer
	tsvFile    *os.File
}

func (log *TsvLog) getOrCreateFile() {
	_, err := os.Stat(log.fileName)
	creatingNewFile := os.IsNotExist(err)

	log.tsvFile, err = os.OpenFile(log.fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
	}

	log.writer = csv.NewWriter(log.tsvFile)
	log.writer.Comma = '\t'

	if creatingNewFile {
		log.writer.Write(log.headers)
	}
}

func Create(headers []string, path string, format string) *TsvLog {
	headersWithTs := append([]string{"ts"}, headers...) // prepend timestamp header

	log := TsvLog{
		headers:    headersWithTs,
		fileName:   path,
		timeFormat: format}

	return &log
}

func (log *TsvLog) Add(data []string) error {
	var err error

	if len(data) != len(log.headers)-1 {
		err = errors.New("csv data length doesnt match header length")
	} else {
		log.getOrCreateFile()

		log.writer.Write(append([]string{time.Now().Format(log.timeFormat)}, data...))
		log.Close()
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

	file.Close()
	return records[0], records[1:]
}

func (log *TsvLog) Delete() {
	os.Remove(log.fileName)
}

func (log *TsvLog) Close() {
	log.writer.Flush()
	log.tsvFile.Close()
}
