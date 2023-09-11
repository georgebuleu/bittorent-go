package main

import (
	// Uncomment this line to pass the first stage
	// "encoding/json"
	"encoding/json"
	"fmt"
	"os"
	// bencode "github.com/jackpal/bencode-go" // Available if you need it!
)

// Example:
// - 5:hello -> hello
// - 10:hello12345 -> hello12345

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")

	command := os.Args[1]

	if command == "decode" {

		bencodedValue := os.Args[2]
		decoded, _, err := decodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}
		jsonOutput, err := json.Marshal(decoded)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(jsonOutput))
	} else if command == "info" {
		path := os.Args[2]
		info, err := parseTorrentFile(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Tracker URL:%v\nLength:%v", info.URL, info.length)
		//fmt.Printf("" info.length * 8)
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
