package counter_id

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type TextCounter interface {
	SaveFile(id int) error
	ReadFile(filePath string) (int, error)
}

type textCounter struct {
	filePath string
	indent   uint64
}

func NewtextCounter(filePath string, indent uint64) (TextCounter, error) {
	return &textCounter{filePath, indent}, nil
}

func (t *textCounter) SaveFile(id int) error {
	strId := strconv.Itoa(id)

	var d = []byte(strId)

	err := ioutil.WriteFile(t.filePath, d, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (t *textCounter) ReadFile(filePath string) (int, error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read fail", err)
	}

	lastId, err := strconv.Atoi(string(f))

	if err != nil {
		return 0, err
	}

	return lastId, nil
}
