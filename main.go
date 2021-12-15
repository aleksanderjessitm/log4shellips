package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// The url to the list of IPs in Plain Text
	url := "https://raw.githubusercontent.com/CriticalPathSecurity/Public-Intelligence-Feeds/master/log4j.txt"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Converting the body to strings
		bodyToString := string(body)
		bodyToStringSplit := strings.Split(bodyToString, "\n")

		// Now let's write it all nicely to a csv
		// First, let's create a new file
		csvFile, err := os.Create("ips.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer csvFile.Close()

		// Now let's write it
		csvWriter := csv.NewWriter(csvFile)
		for _, row := range bodyToStringSplit {
			err = csvWriter.Write([]string{row})
			if err != nil {
				log.Fatal("Something went wrong while writing to a csv", err)
			}
		}
		csvWriter.Flush()

	}

}
