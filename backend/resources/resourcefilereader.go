package resources

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	delimiter = ';'
)

//ResourceFileReader stores the reader of the resource file
type ResourceFileReader struct {
	reader *csv.Reader
}

//SetReaderByPath sets the reader provided a path to a resource file.
func (r *ResourceFileReader) SetReaderByPath(path string) {
	csvFile := openCsvFile(path)
	newReader := csv.NewReader(bufio.NewReader(csvFile))
	newReader.Comma = delimiter
	r.reader = newReader
}

//GetLines collects all lines from the file the reader is set up to read.
func (r *ResourceFileReader) GetLines() ([][]string, error) {
	var lines [][]string
	for {
		line, err := r.reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func openCsvFile(path string) *os.File {
	csvFile, openFileError := os.Open(path)
	if openFileError != nil {
		dir, _ := os.Getwd()
		logrus.Fatalf("Please confirm that the specified path is relative to the current working directory. Current directory: %v", dir)
	}
	return csvFile
}
