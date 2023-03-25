// storage.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Storage interface {
	LoadNotes() ([]Note, error)
	SaveNotes(notes []Note) error
}

type JSONStorage struct {
	filename string
}

func NewJSONStorage(filename string) (*JSONStorage, error) {
	js := &JSONStorage{filename: filename}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := ioutil.WriteFile(filename, []byte("[]"), 0644)
		if err != nil {
			return nil, err
		}
	}

	return js, nil
}

func (js *JSONStorage) LoadNotes() ([]Note, error) {
	data, err := ioutil.ReadFile(js.filename)
	if err != nil {
		return nil, err
	}

	var notes []Note
	err = json.Unmarshal(data, &notes)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (js *JSONStorage) SaveNotes(notes []Note) error {
	data, err := json.Marshal(notes)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(js.filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
