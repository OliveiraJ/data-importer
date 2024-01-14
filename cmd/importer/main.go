package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	//"github.com/rivo/tview"
)

func readSample(rs io.ReadSeeker) ([][]string, error) {
	// Skip first row (line)
	row1, err := bufio.NewReader(rs).ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	_, err = rs.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Read remaining rows
	r := csv.NewReader(rs)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func main() {
	file, err := os.Open("../../file/clientes.csv")
	if err != nil {
		fmt.Println(err)
	}
	//reader := csv.NewReader(file)
	//records, _ := reader.ReadAll()
	//record, _ := reader.Read()
	record, err := readSample(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(record)

	//	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	//
	// if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	// panic(err)
	// }
}
