// /*
// Copyright © 2025 NAME HERE <EMAIL ADDRESS>
// */
package main

import (

	// "log"
	"github.com/yashpal2104/json-viewer-cli/cmd"
	// "github.com/yashpal2104/json-viewer-cli/data"

)



func main() {
	cmd.Execute()
	




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
