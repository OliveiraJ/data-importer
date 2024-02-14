package file

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/OliveiraJ/data-importer/internal/schema"
	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
)

// Reads file and unmarshal the data into the provided struct (entity)
func ReadFile[R *[]schema.Customer | *[]schema.Supplier | *[]schema.People](registers R) {
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
