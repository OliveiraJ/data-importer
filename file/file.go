package file

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
)

func ReadFile(registers interface{}) {
	fmt.Println("path: ", viper.GetString("pathFile"))

	file, err := os.Open(viper.GetString("pathFile"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	skipFirstline(r)

	err = gocsv.UnmarshalCSV(r, registers)
	if err != nil {
		panic(err)
	}
}

func skipFirstline(r *csv.Reader) {
	_, err := r.Read()
	if err != nil {
		panic(err)
	}
}
