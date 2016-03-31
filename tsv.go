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
}

func Create(headers []string, path string, format string) *TsvLog {
	tsvFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
	}

	writer := csv.NewWriter(tsvFile)
	writer.Comma = '\t'

	log := TsvLog{
		headers:    append([]string{"ts"}, headers...),
		fileName:   path,
		timeFormat: format}

	if len(headers) > 0 {
		writer.Write(log.headers)
	}

	writer.Flush()
	tsvFile.Close()
	return &log
}

func (log *TsvLog) Add(data []string) error {
	var err error
	var tsvFile *os.File

	if len(data) != len(log.headers)-1 {
		err = errors.New("csv data length doesnt match header length")
	} else {
		_, err := os.Stat(log.fileName)
		creatingNewFile := os.IsNotExist(err)
		tsvFile, err = os.OpenFile(log.fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)

		if err != nil {
			fmt.Println(err)
		} else {

			writer := csv.NewWriter(tsvFile)
			writer.Comma = '\t'

			if creatingNewFile {
				writer.Write(log.headers)
			}
			writer.Write(append([]string{time.Now().Format(log.timeFormat)}, data...))
			writer.Flush()
			tsvFile.Close()
		}
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
