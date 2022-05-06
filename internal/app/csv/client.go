package csv

import (
	"encoding/csv"
	"log"
	"os"
)

type HandlersCsv interface {
	ReaderCsv()
	WriterCsv() error
}

type CounterCsv struct {
	ListExpansion map[string]int
	FilePath      string
}

func NewCounterCsv(filePath string) HandlersCsv {
	CounterCsv := &CounterCsv{make(map[string]int), filePath}
	CounterCsv.ReaderCsv()
	return CounterCsv
}

func (c *CounterCsv) ReaderCsv() {
	//f, err := os.Open(file)
	//defer f.Close()
	//
	//if err != nil {
	//	return err
	//}
	//lines, err := csv.NewReader(f).ReadAll()
	//if err != nil {
	//	return err
	//}
	//
	//for nl, line := range lines {
	//	if nl > 0 {
	//		db_csv.InsertDataCSVFile(line, fileName)
	//		fmt.Println(line)
	//	}
	//	nl++
	//}
}

func (c *CounterCsv) WriterCsv() error {
	records := [][]string{{}}

	for val, key := range c.ListExpansion {
		records = append(records, []string{val, string(key)})
	}

	file, errCreate := os.Create(c.FilePath)
	if errCreate != nil {
		log.Panic(errCreate)
	}

	w := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	return nil
}
