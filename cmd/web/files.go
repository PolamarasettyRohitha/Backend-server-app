package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
)

func readCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func writeCsv(path string, data []string) error {

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	if err := writer.Write(data); err != nil {
		return err
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}

func renameCsv(oldpath, newpath string) error {
	if err := os.Remove(newpath); err != nil {
		return err
	}

	return os.Rename(oldpath, newpath)
}

func display(w http.ResponseWriter, filepath string) {
	records, err := readCsv(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, eachrecord := range records {
		for _, each := range eachrecord {
			w.Write([]byte(fmt.Sprintf("%v \t", each)))
		}
		w.Write([]byte("\n"))
	}
}
