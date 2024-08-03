package main

import (
	"encoding/json"
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// https://irshadhasmat.medium.com/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968
	// Simple Employee JSON object which we will parse
	empJson := `{
		"id": 11,
		"name": "Irshad",
		"department": "IT",
		"designation": "Product Manager"
	}`

	// Declare an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empJson), &result)

	// Reading each value by it's key
	// fmt.Println("Id: ", result["id"],
	// 	"\nName: ", result["name"],
	// 	"\nDepartment: ", result["department"],
	// 	"\nDesignation: ", result["designation"])
	for key, value := range result {
		fmt.Printf("%s: %v\n", cases.Title(language.English, cases.Compact).String(key), value)
	}

	// https://irshadhasmat.medium.com/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968
	// Parsing Simple Array JSON
	empArray := `[
		{
			"id": 1,
			"name": "Mr. Boss",
			"department": "",
			"designation": "Director"
		},
		{
			"id": 11,
			"name": "Irshad",
			"department": "IT",
			"designation": "Product Manager"
		},
		{
			"id": 12,
			"name": "Pankaj",
			"department": "IT",
			"designation": "Team Lead"
		}
	]`

	// Declare an empty interface of type array
	// Note this one has [] at front of statement for slice/array
	var resultsArr []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface
	json.Unmarshal([]byte(empArray), &resultsArr)

	// Iterate over each array item to print results
	for key, result := range resultsArr {
		fmt.Print("Reading Value for Key: ", key)
		// Reading each value by its key
		fmt.Println(" - Id: ", result["id"],
			" - Name: ", result["name"],
			" - Department: ", result["department"],
			" - Designation: ", result["designation"])
	}
	for index, obj := range resultsArr {
		fmt.Printf("Reading value for Key: %d", index)
		for key, value := range obj {
			fmt.Printf(" - %s: %v", cases.Title(language.English, cases.Compact).String(key), value)
		}
		fmt.Printf("\n")
	}

	// https://irshadhasmat.medium.com/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968
	// Parsing Embedded objects in JSON
	empJsonEmbed := `{
		"id": 11,
		"name": "Irshad",
		"department": "IT",
		"designation": "Product Manager",
		"address": {
			"city": "Mumbai",
			"state": "Maharashtra",
			"country": "India"
		}
	}`

	// Declare an empty interface
	var resultEmbed map[string]interface{}

	// Unmarshal or Decode JSON to the interface
	json.Unmarshal([]byte(empJsonEmbed), &resultEmbed)

	address := resultEmbed["address"].(map[string]interface{})

	// Reading each value by its key
	fmt.Println("Id: ", result["id"],
		"\nName: ", result["name"],
		"\nDepartment: ", result["department"],
		"\nDesignation: ", result["designation"],
		"\nAddress: ", address["city"], address["state"], address["country"])

	//
	// Parsing Embedded object in Array of JSON
	empArrayEmbed := `[
		{
			"id": 1,
			"name": "Mr. Boss",
			"department": "",
			"designation": "Director",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 11,
			"name": "Irshad",
			"department": "IT",
			"designation": "Product Manager",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 12,
			"name": "Pankaj",
			"department": "IT",
			"designation": "Team Lead",
			"address": {
				"city": "Pune",
				"state": "Maharashtra",
				"country": "India"
			}
		}
	]`

	// Declare an empty interface of type Array
	var resultArrEmbed []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface
	json.Unmarshal([]byte(empArrayEmbed), &resultArrEmbed)

	// Run through each item in array and print results
	for key, result := range resultArrEmbed {
		address := result["address"].(map[string]interface{})
		fmt.Printf("Reading Value for Key: %v ", key)
		// Reading each value by its key
		fmt.Println("- Id:", result["id"],
			"- Name:", result["name"],
			"- Department:", result["department"],
			"- Designation:", result["designation"],
			"- Address:", address["city"], address["state"], address["country"])
	}

	fmt.Println("")
}
