package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("new_time.csv")
	if err != nil {
		log.Fatal("Unable to open CSV file", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to read CSV file", err)
	}

	// Open a text file for writing the generated SQL queries
	outputFile, err := os.Create("time_updated.txt")
	if err != nil {
		log.Fatal("Unable to create output file", err)
	}
	defer outputFile.Close()

	// Loop through the CSV records
	for i, record := range records[1:] { // Skip header row (index 0)
		// Get the SKU from Column A (index 0 in Go arrays)
		sku := record[0]

		// Generate the SQL query
		query := fmt.Sprintf("UPDATE global_object SET timeUpdated = now() WHERE id = (SELECT globalId FROM store_product WHERE sku = '%s');\n", sku)

		// Write the query into the .txt file
		_, err := outputFile.WriteString(query)
		if err != nil {
			log.Printf("Error writing query for row %d: %v", i+2, err)
		}
	}

	fmt.Println("SQL update queries generated successfully in 'time_updated.txt'")
}
