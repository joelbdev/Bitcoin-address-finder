package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func writeCsv(matches []string, filename string) {
	csvFile, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating the output csv file: %s", err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	err = w.Write(matches)
	if err != nil {
		log.Fatalf("error writing the record to the file, %s", err)
	}
}

func main() {
	var matches []string

	//open the file, TODO: implement directory scanning

	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	//read the csv data

	csvReader := csv.NewReader(f)

	re, _ := regexp.Compile("^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$")
	//"(5[HJK][1-9A-Za-z][^OIl]{48}) | ^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$ | (X[1-9A-HJ-NP-Za-km-z]{33}) | (0x[a-fA-F0-9]{40}) | (r[0-9a-zA-Z]{24,34}) | (D{1}[5-9A-HJ-NP-U]{1}[1-9A-HJ-NP-Za-km-z]{32}) | (A[0-9a-zA-Z]{33}) | (4[0-9AB][0-9a-zA-Z]{93,104})"
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		for _, word := range line {
			match := re.MatchString(word)

			if err != nil {
				log.Fatal(err)
			}
			if match {
				matches = append(matches, line...)
				matches = append(matches, "\n")
			} else {
				continue
			}
		}
		//matches = append(matches, "\n")
	}
	fmt.Println(matches)
	writeCsv(matches, "results.csv")

}
