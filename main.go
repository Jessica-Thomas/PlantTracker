package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var filename = "plantlist.csv"

// still need a struct

func main() {
	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")

	fmt.Println("Reading single row at a time")
	csvReaderRow()

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	//testing out reading functions.... so we know we can get the csv read in by the Go code.... now we just need to get it to pull up in our web app as a dropdown
	//range over elements to select option element and html template
	//read in CSV and range over into a list?

}

//range over elements to select option element and html template
//read in CSV and range over into

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

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful!!\n")
	name := r.FormValue("name")
	lastWatered := r.FormValue("date")
	fmt.Fprintf(w, "Plant name = %s\n", name)
	fmt.Fprintf(w, "Last watered on = %s\n", lastWatered)
}
