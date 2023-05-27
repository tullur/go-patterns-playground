package behavioral

import (
	"log"
)

type Parser interface {
	Parse()
}

type File struct {
	Parser
}

func (f *File) ParseFile(path string) {
	f.openFile()
	f.readFileData()
	f.Parse()
	f.generateReport()
}

func NewParser(parser Parser) *File {
	return &File{parser}
}

func (f *File) openFile() {
	log.Println("File: Open Current File")
}

func (f *File) readFileData() {
	log.Println("File: Read Data")
}

func (f *File) generateReport() {
	log.Println("File: Generate Parse operation report")
}

type CsvFile struct{}

func (csv *CsvFile) Parse() {
	log.Println("File: Parse CSV file")
}

type JsonFile struct{}

func (csv *JsonFile) Parse() {
	log.Println("File: Parse JSON file")
}
