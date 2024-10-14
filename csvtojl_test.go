package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestName(t *testing.T) {
	inputFile, err := os.Open("housesInput.csv")
	if err != nil {
		fmt.Printf("Error opening input file: %s \n", err)
	}
	defer inputFile.Close()
	csvReader := csv.NewReader(inputFile)
	_, err = csvReader.Read()
	record, err := csvReader.Read()
	if len(record) != 7 {
		fmt.Printf("Missing elements in record: %s, skipping...", record)
	}
	value, err := strconv.ParseFloat(record[0], 64)
	if err != nil {
		fmt.Printf("Fail to parse Value to int64: %s, skipping this record...\n", err)
	}

	income, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		fmt.Printf("Fail to parse Income to float64: %s, skipping this record...\n", err)
	}

	age, err := strconv.ParseInt(record[2], 10, 32)
	if err != nil {
		fmt.Printf("Fail to parse Age to int32: %s, skipping this record...\n", err)
	}

	rooms, err := strconv.ParseInt(record[3], 10, 32)
	if err != nil {
		fmt.Printf("Fail to parse Rooms to int32: %s, skipping this record...\n", err)
	}

	bedrooms, err := strconv.ParseInt(record[4], 10, 32)
	if err != nil {
		fmt.Printf("Fail to parse Bedrooms to int32: %s, skipping this record...\n", err)
	}

	pop, err := strconv.ParseInt(record[5], 10, 32)
	if err != nil {
		fmt.Printf("Fail to parse Pop to int32: %s, skipping this record...\n", err)
	}

	hh, err := strconv.ParseInt(record[6], 10, 32)
	if err != nil {
		fmt.Printf("Fail to parse Hh to int32: %s, skipping this record...\n", err)
	}
	jsonLine := parseJson(value, income, age, rooms, bedrooms, pop, hh, err)
	if string(jsonLine) != "{\"value\":452600,\"income\":8.3252,\"age\":41,\"rooms\":880,\"bedrooms\":129,\"pop\":322,\"hh\":126}" {
		t.Errorf("{\"value\":452600,\"income\":8.3252,\"age\":41,\"rooms\":880,\"bedrooms\":129,\"pop\":322,\"hh\":126} expected, but %s got", string(jsonLine))
	}
}
