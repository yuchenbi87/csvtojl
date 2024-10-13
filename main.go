package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CSVRecord struct {
	Value    int64   `json:"value"`
	Income   float64 `json:"income"`
	Age      int32   `json:"age"`
	Rooms    int32   `json:"rooms"`
	Bedrooms int32   `json:"bedrooms"`
	Pop      int32   `json:"pop"`
	Hh       int32   `json:"hh"`
}

func main() {
	var args = os.Args
	if len(args) < 3 {
		fmt.Println("Invalid number of arguments. Expecting at least 2")
		panic("Ending the program.")
	}
	csv_path := args[1]
	jl_path := args[2]

	inputFile, err := os.Open(csv_path)
	if err != nil {
		fmt.Println("Error opening input file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(jl_path)
	if err != nil {
		fmt.Println("Error creating output file: %v", err)
		panic("Ending the program.")
	}
	defer outputFile.Close()

	csvReader := csv.NewReader(inputFile)
	_, err = csvReader.Read()
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Fail to reading CSV record: %v", err)
		}
		if len(record) != 7 {
			fmt.Println("Missing elements in record: %v, skipping...", record)
			continue
		}
		value, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			fmt.Println("Fail to parse Value to int64: %v, skipping this record...", err)
			continue
		}

		income, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println("Fail to parse Income to float64: %v, skipping this record...", err)
			continue
		}

		age, err := strconv.ParseInt(record[2], 10, 32)
		if err != nil {
			fmt.Println("Fail to parse Age to int32: %v, skipping this record...", err)
			continue
		}

		rooms, err := strconv.ParseInt(record[3], 10, 32)
		if err != nil {
			fmt.Println("Fail to parse Rooms to int32: %v, skipping this record...", err)
			continue
		}

		bedrooms, err := strconv.ParseInt(record[4], 10, 32)
		if err != nil {
			fmt.Println("Fail to parse Bedrooms to int32: %v, skipping this record...", err)
			continue
		}

		pop, err := strconv.ParseInt(record[5], 10, 32)
		if err != nil {
			fmt.Println("Fail to parse Pop to int32: %v, skipping this record...", err)
			continue
		}

		hh, err := strconv.ParseInt(record[6], 10, 32)
		if err != nil {
			fmt.Println("Fail to parse Hh to int32: %v, skipping this record...", err)
			continue
		}

		r := CSVRecord{
			Value:    int64(value),
			Income:   income,
			Age:      int32(age),
			Rooms:    int32(rooms),
			Bedrooms: int32(bedrooms),
			Pop:      int32(pop),
			Hh:       int32(hh),
		}

		jsonLine, err := json.Marshal(r)
		if err != nil {
			log.Fatalf("Error converting record to JSON: %v", err)
		}

		outputFile.Write(jsonLine)
		outputFile.Write([]byte("\n"))
	}
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}
