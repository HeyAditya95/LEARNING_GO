// --------task 5 -> working with json file-----------
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

// define library struct
type Library struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

func main() {
	library := Library{
		Name: "maheshpur library",
		Books: []Book{
			{Title: "Golang Mastery", Author: "Alice Brown", Price: 50},
			{Title: "Concurrency in Go", Author: "Bob White", Price: 45},
		},
	}

	//converting struct to json
	jsonData, err := json.MarshalIndent(library, "", "  ")
	if err != nil {
		fmt.Println("Error in Encoding Json", err)
		return
	}

	// save the file/write in file
	err = os.WriteFile("library.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error in writing Json", err)
		return
	}
	fmt.Println("json saved the file")

	//read json from file
	jsonData, err = os.ReadFile("library.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	//decode json back to struct
	var loadedLibrary Library
	err = json.Unmarshal(jsonData, &loadedLibrary)
	if err != nil {
		fmt.Println("Error parsing json:", err)
		return
	}
	fmt.Println("\nDecoded from file:")
	fmt.Println("Library Name:", loadedLibrary.Name)
	for i, book := range loadedLibrary.Books {
		fmt.Printf("Book %d: %s by %s ($%d)\n", i+1, book.Title, book.Author, book.Price)
	}
}

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//// Define Book struct
//type Book struct {
//	Title  string `json:"title"`
//	Author string `json:"author"`
//	Price  int    `json:"price"`
//}
//
//// Define Library struct (which contains Books)
//type Library struct {
//	Name  string `json:"name"`
//	Books []Book `json:"books"` // Slice of books (nested struct)
//}
//
//func main() {
//	// Step 1: Create a library with books
//	library := Library{
//		Name: "City Library",
//		Books: []Book{
//			{Title: "Golang Basics", Author: "John Doe", Price: 30},
//			{Title: "Advanced Go", Author: "Jane Smith", Price: 40},
//		},
//	}
//
//	// Step 2: Convert Library struct to JSON
//	jsonData, err := json.MarshalIndent(library, "", "  ")
//	if err != nil {
//		fmt.Println("Error encoding JSON:", err)
//		return
//	}
//	fmt.Println("Encoded JSON:\n", string(jsonData))
//
//	// Step 3: Convert JSON back to struct
//	var decodedLibrary Library
//	err = json.Unmarshal(jsonData, &decodedLibrary)
//	if err != nil {
//		fmt.Println("Error decoding JSON:", err)
//		return
//	}
//
//	// Step 4: Print the decoded data
//	fmt.Println("\nDecoded Struct:")
//	fmt.Println("Library Name:", decodedLibrary.Name)
//	for i, book := range decodedLibrary.Books {
//		fmt.Printf("Book %d: %s by %s ($%d)\n", i+1, book.Title, book.Author, book.Price)
//	}
//}

//package main
//
////-------------------task 4 -  JSON Encoding & Decoding with Nested Structs-----------
//import (
//	"encoding/json"
//	"fmt"
//)
//
//type Book struct {
//	Title  string `json:"title"`
//	Author string `json:"author"`
//	Price  int    `json:"price"`
//}
//type Library struct {
//	Name  string `json:"name"`
//	Books []Book `json:"books"`
//}
//
//func main() {
//	library := Library{
//		Name: "meethapur library",
//		Books: []Book{
//			{Title: "Go basics", Author: "Michael Stark", Price: 100},
//			{Title: "Adavanced Go", Author: "Michael Stark", Price: 200},
//		},
//	}
//	//converting the struct into json
//	jsonData, err := json.MarshalIndent(library, "", "  ")
//	if err != nil {
//		fmt.Println("here we facing an error to convert the struct into the json", err)
//		return
//	}
//	fmt.Println("the encoded json string is :", string(jsonData))
//
//	var decodedlibrary Library
//	err = json.Unmarshal(jsonData, &decodedlibrary)
//	if err != nil {
//		fmt.Println("here we facing an error to convert the json into the library", err)
//		return
//	}
//	fmt.Println("\nDecoded library is :")
//	fmt.Println("the name of the library is", decodedlibrary.Name)
//	for i, book := range decodedlibrary.Books {
//		fmt.Printf("Book %d: %s by %s ($%d)\n", i+1, book.Title, book.Author, book.Price)
//
//	}
//}

//
//type Book struct {
//	Title  string `json:"title"`
//	Author string `json:"author"`
//	Price  int    `json:"price"`
//}
//
//func main() {
//	jsonString := `{"title" : "aditya" , "author" : "life" , "price" : 30 }`
//	var book Book
//	err := json.Unmarshal([]byte(jsonString), &book)
//	if err != nil {
//		fmt.Println("error converting from JSON : ", err)
//		return
//	}
//	fmt.Println("title of the book is :", book.Title, "and author is :", book.Author, "with the price of", book.Price)
//}

//func main() {
//	Book1 := []Book{
//		{Title: "life", Author: "aditya", Price: 10},
//		{Title: "with aditya", Author: "ananya", Price: 10},
//	}
//	jsonData, err := json.Marshal(Book1)
//	if err != nil {
//		fmt.Println("error converting the JSON:", err)
//		return
//	}
//	fmt.Println(string(jsonData))
//}
