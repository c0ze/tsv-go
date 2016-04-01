package tsv

import (
	"reflect"
	"testing"
)

const timeFormat = "2006-01-02 3:04:05"

func TestLog(t *testing.T) {
	headers := []string{"temperature", "voltage"}
	syslog := Create(headers, "./test.tsv", timeFormat)

	data := []string{"50.6", "1.122"}
	data2 := []string{"48", "1.3"}
	data3 := []string{"48", "1.4", "ahmet"}

	syslog.Add(data)
	syslog.Add(data2)
	err := syslog.Add(data3)
	if err == nil {
		t.Errorf("adding malformed data failed to raise error")
	}

	readHeaders, readData := syslog.Read()
	if !reflect.DeepEqual(headers, readHeaders[1:]) {
		t.Errorf("reading headers failed from the tsv")
	}

	if !reflect.DeepEqual(data, readData[0][1:]) {
		t.Errorf("reading data failed from the tsv")
	}

	if !reflect.DeepEqual(data2, readData[1][1:]) {
		t.Errorf("reading data2 failed from the tsv")
	}

	syslog.Delete()
}

func TestServerRestart(t *testing.T) {
	headers := []string{"temperature", "voltage"}
	syslog := Create(headers, "./test.tsv", timeFormat)
	syslog = Create(headers, "./test.tsv", timeFormat)

	data := []string{"50.6", "1.122"}
	data2 := []string{"48", "1.3"}

	syslog.Add(data)
	syslog.Add(data2)

	readHeaders, readData := syslog.Read()
	if !reflect.DeepEqual(headers, readHeaders[1:]) {
		t.Errorf("reading headers failed from the tsv")
	}

	if !reflect.DeepEqual(data, readData[0][1:]) {
		t.Errorf("reading data failed from the tsv")
	}

	if !reflect.DeepEqual(data2, readData[1][1:]) {
		t.Errorf("reading data2 failed from the tsv")
	}

	syslog.Delete()
}

func TestLogRotate(t *testing.T) {
	headers := []string{"temperature", "voltage"}
	syslog := Create(headers, "./test.tsv", timeFormat)

	data := []string{"50.6", "1.122"}
	data2 := []string{"48", "1.3"}
	data3 := []string{"48", "1.4", "ahmet"}

	syslog.Add(data)
	syslog.Delete()

	syslog.Add(data2)
	err := syslog.Add(data3)
	if err == nil {
		t.Errorf("adding malformed data failed to raise error")
	}

	readHeaders, readData := syslog.Read()
	if !reflect.DeepEqual(headers, readHeaders[1:]) {
		t.Errorf("reading headers failed from the tsv")
	}

	if !reflect.DeepEqual(data2, readData[0][1:]) {
		t.Errorf("reading data failed from the tsv")
	}

	syslog.Delete()
}
