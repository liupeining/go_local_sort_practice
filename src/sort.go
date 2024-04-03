package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"sort"
)

type Record struct {
	Key   [10]byte
	Value [90]byte
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error in opening input file - %v", err)
	}
	defer file.Close()

	var records []Record
	buffer := make([]byte, 100)
	for {
		_, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error in reading file: %v", err)
		}
		var record Record
		copy(record.Key[:], buffer[:10])
		copy(record.Value[:], buffer[10:])
		records = append(records, record)
	}

	sort.Slice(records, func(i, j int) bool {
		return bytes.Compare(records[i].Key[:], records[j].Key[:]) < 0
	})

	output, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalf("Error in creating output file - %v", err)
	}
	defer output.Close()

	for _, record := range records {
		_, err := output.Write(record.Key[:])
		if err != nil {
			log.Fatalf("Error in writing to file - %v", err)
		}
		_, err = output.Write(record.Value[:])
		if err != nil {
			log.Fatalf("Error in writing to file - %v", err)
		}
	}

	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])
}
