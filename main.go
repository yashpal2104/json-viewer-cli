// /*
// Copyright © 2025 NAME HERE <EMAIL ADDRESS>
// */
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	// "log"
	"net/http"
	"os"
	"strings"

	"github.com/yashpal2104/json-viewer-cli/cmd"
)

type Data struct {
	Fruit string `json:"fruit"`
	Size  string `json:"size"`
	Color string `json:"color"`
}

var FileParsedData []Data

func readJSON(path string) ([]byte, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		//URL, fetch the data
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return io.ReadAll(resp.Body)
	} else {
		//  It's a file path, read the file
		return os.ReadFile(path)

	}
}

func main() {
	cmd.Execute()
	// For a local file path
	filePath := "json_file.json" //the file path
	fileData, err := readJSON(filePath)
	key := flag.String("key", "", "Key to filter from the JSON")
	flag.Parse()

	if err != nil {
		fmt.Println("Error reading a file: ", err)
		return
	}

	err = json.Unmarshal(fileData, &FileParsedData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	for i, d := range FileParsedData {
		fmt.Printf("Data from file: %d. %+v\n", i+1, d)
	}

	// Filter key if provided
	if *key != "" {
		for i, d := range FileParsedData {
			switch *key {
			case "fruit":
				fmt.Printf("%d. Fruit: %s\n", i+1, d.Fruit)
			case "size":
				fmt.Printf("%d. Size: %s\n", i+1, d.Size)
			case "color":
				fmt.Printf("%d. Color: %s\n", i+1, d.Color)
			default:
				fmt.Println("Unknown key:", *key)
				return
			}
		}
	} else {
		// Pretty-print all
		prettyJSON, err := json.MarshalIndent(FileParsedData, "", "  ")
		if err != nil {
			fmt.Println("Error formatting JSON:", err)
			return
		}
		fmt.Println(string(prettyJSON))
	}

	// For usage with a URL file path
	// urlPath := "https://dummyjson.com/image/400x200/008080/ffffff?text=Hello+Peter" // Replace with your URL
	// urlData, err := readJSON(urlPath)
	// if err != nil {
	// 	fmt.Println("Error reading URL:", err)
	// 	return
	// }

	// var urlParsedData Data
	// err = json.Unmarshal(urlData, &urlParsedData)
	// if err != nil {
	// 	fmt.Println("Error unmarshaling JSON:", err)
	// 	return
	// }

	// fmt.Printf("Data from URL: %+v\n", urlParsedData)

}

// func getJSONBodyFromURL() {
// 	url := "https://api.github.com"
// 	// resp, err := http.Get(url)
// 	res, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := io.ReadAll(res.Body)
// 	res.Body.Close()
// 	if res.StatusCode > 299 {
// 		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", body)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	data := map[string]interface{}{
// 		"intValue":    1234,
// 		"boolValue":   true,
// 		"stringValue": "hello!",
// 		"objectValue": map[string]interface{}{
// 			"arrayValue": []int{1, 2, 3, 4},
// 		},
// 	}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Printf("could not marshal json: %s\n", err)
// 		return
// 	}

// 	fmt.Printf("json data: %s\n", jsonData)
// }
