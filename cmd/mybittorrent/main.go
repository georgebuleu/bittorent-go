package main

import (
	"crypto/sha1"
	"encoding/hex"

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
		torrent, err := parseTorrentFile(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Tracker URL: %s", torrent.url)
		fmt.Printf("Length: %d", torrent.length)
		data, err := parseInfo(path)
		if err != nil {
			fmt.Printf("error:%v", err)
			return
		}
		encoding, err := bencode(data)
		if err != nil {
			fmt.Printf("error:%v", err)
			return
		}

		hashedBytes := sha1.Sum([]byte(encoding))
		hashedInfo := hex.EncodeToString(hashedBytes[:])
		fmt.Printf("Info Hash: %s", hashedInfo)
		fmt.Printf("Piece Length: %d", torrent.pieceLength)
		nPieces := len(torrent.pieces) / 20
		fmt.Printf("Piece Hashes:\n")
		for i := 0; i < nPieces-1; i++ {
			fmt.Printf("%v", hex.EncodeToString([]byte(torrent.pieces[i*20:(i+1)*20])))
		}

	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
