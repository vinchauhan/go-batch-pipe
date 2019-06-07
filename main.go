package main

import (
	"fmt"
	"github.com/vinchauhan/indexer/reader"
	"github.com/vinchauhan/indexer/types"
)

func main()  {


	//Create a reader and specify where to read from
	r := types.NewReader(reader.File, 3)

	fmt.Printf("Reader is of type %v\n", r)

	rr, ok := r.(*types.FileReader)
	if !ok {
		fmt.Printf("The reader is not a FileReader")
	}
	out := rr.read()

}