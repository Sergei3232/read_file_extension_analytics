package counter_id

import (
	"io/ioutil"
	"strconv"
)

type TextCounter interface {
	SaveFile(id int) error
	readFile() error
}

type TextCounterStruct struct {
	FilePath string
	Indent   uint64
}

func NewtextCounter(filePath string, indent uint64) (TextCounter, error) {
	counter := &TextCounterStruct{filePath, indent}
	counter.readFile()

	return counter, nil
}

func (t *TextCounterStruct) SaveFile(id int) error {
	strId := strconv.Itoa(id)

	var d = []byte(strId)

	err := ioutil.WriteFile(t.FilePath, d, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (t *TextCounterStruct) readFile() error {
	f, err := ioutil.ReadFile(t.FilePath)
	if err != nil {
		return err
	}

	lastId, err := strconv.ParseUint(string(f), 10, 64)

	if err != nil {
		return err
	}
	t.Indent = lastId
	return nil
}
