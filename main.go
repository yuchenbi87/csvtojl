package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
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
		fmt.Printf("Error opening input file: %s \n", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(jl_path)
	if err != nil {
		fmt.Printf("Error creating output file: %s \n", err)
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
			fmt.Printf("Fail to reading CSV record: %s", err)
		}
		if len(record) != 7 {
			fmt.Printf("Missing elements in record: %s, skipping...", record)
			continue
		}
		value, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			fmt.Printf("Fail to parse Value to int64: %s, skipping this record...\n", err)
			continue
		}

		income, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Printf("Fail to parse Income to float64: %s, skipping this record...\n", err)
			continue
		}

		age, err := strconv.ParseInt(record[2], 10, 32)
		if err != nil {
			fmt.Printf("Fail to parse Age to int32: %s, skipping this record...\n", err)
			continue
		}

		rooms, err := strconv.ParseInt(record[3], 10, 32)
		if err != nil {
			fmt.Printf("Fail to parse Rooms to int32: %s, skipping this record...\n", err)
			continue
		}

		bedrooms, err := strconv.ParseInt(record[4], 10, 32)
		if err != nil {
			fmt.Printf("Fail to parse Bedrooms to int32: %s, skipping this record...\n", err)
			continue
		}

		pop, err := strconv.ParseInt(record[5], 10, 32)
		if err != nil {
			fmt.Printf("Fail to parse Pop to int32: %s, skipping this record...\n", err)
			continue
		}

		hh, err := strconv.ParseInt(record[6], 10, 32)
		if err != nil {
			fmt.Printf("Fail to parse Hh to int32: %s, skipping this record...\n", err)
			continue
		}

		jsonLine := parseJson(value, income, age, rooms, bedrooms, pop, hh, err)
		if err != nil {
			fmt.Printf("Error converting record to JSON: %s \n", err)
		}

		outputFile.Write(jsonLine)
		outputFile.Write([]byte("\n"))
	}
}

func parseJson(value float64, income float64, age int64, rooms int64, bedrooms int64, pop int64, hh int64, err error) []byte {
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
	return jsonLine
}
