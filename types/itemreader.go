package types

import (
	"io/ioutil"
)

type FileReader struct {
	filename string
	filepath string
	mode string
}

//type ItemReader interface {
//	read(reader Reader) <-chan interface{}
//}


func NewReader(s string, i int) interface{} {

	switch s {

	case "file":
		return &FileReader{
			filename:"input",
			filepath:".",
			mode:"line",
		}
	default:
		return nil
	}
}

func (r *FileReader) read() <- chan interface{} {
	out := make(chan interface{})
	b, err := ioutil.ReadFile(r.filename)
	if err != nil {
		panic(err)
	}
	out <- b
	defer close(out)
	return out
}