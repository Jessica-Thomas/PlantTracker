package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Reading whole file at once")
	csvReaderAll()
	fmt.Println("===================================")
	fmt.Println("Reading single row at a time")
	csvReaderRow()

	//testing out reading functions.... so we know we can get the csv read in by the Go code.... now we just need to get it to pull up in our web app as a dropdown
}

var filename = "plantlist.csv"

func csvReaderAll() {
	// Open the file
	recordFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}

	// Setup the reader
	reader := csv.NewReader(recordFile)

	// Read the records
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}
	fmt.Println(allRecords)

	err = recordFile.Close()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}
}

func csvReaderRow() {
	// Open the file
	recordFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}

	// Setup the reader
	reader := csv.NewReader(recordFile)

	// Read the records
	header, err := reader.Read()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}
	fmt.Printf("Headers : %v \n", header)

	for i := 0; ; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			fmt.Println("An error encountered ::", err)
			return
		}

		fmt.Printf("Row %d : %v \n", i, record)
	}
}

func addPlant() {

}
