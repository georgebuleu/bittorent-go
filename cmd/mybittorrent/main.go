package main

import (
	// Uncomment this line to pass the first stage
	// "encoding/json"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	// bencode "github.com/jackpal/bencode-go" // Available if you need it!
)

// Example:
// - 5:hello -> hello
// - 10:hello12345 -> hello12345
func decodeBencode(bencodedString string) (interface{}, int, error) {
	if unicode.IsDigit(rune(bencodedString[0])) {
		return decodeString(bencodedString)

	} else {
		if rune(bencodedString[0]) == 'i' {
			return decodeInteger(bencodedString)
		}
		if rune(bencodedString[0]) == 'l' {
			return decodeList(bencodedString)
		}

	}
	return "", 0, fmt.Errorf("unsupported type")
}

func decodeList(s string) (interface{}, int, error) {
	list := s[1:]
	var res []interface{}
	totalChars := 0
	for len(list) > 1 {
		if rune(list[0]) == 'e' {
			break
		}
		data, index, err := decodeBencode(list)
		if err != nil {
			return nil, 0, err
		}

		res = append(res, data)
		list = list[index:]
		totalChars += index

	}
	if rune(list[0]) != 'e' {
		return nil, 0, fmt.Errorf("wrong list encoding")
	}

	if len(res) == 0 {
		return []interface{}{}, totalChars + 2, nil
	}

	return res, totalChars + 2, nil

}
func decodeInteger(s string) (int, int, error) {
	var buffer strings.Builder
	index := 1
	for rune(s[index]) != 'e' {
		buffer.WriteByte(s[index])
		index++
	}

	res, err := strconv.Atoi(buffer.String())

	return res, len(buffer.String()) + 2, err

}

func decodeString(s string) (string, int, error) {
	var firstColonIndex = indexOfSemicolon(s)

	lengthStr := s[:firstColonIndex]

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", 0, err
	}

	return s[firstColonIndex+1 : firstColonIndex+1+length], length + firstColonIndex + 1, err
}

func indexOfSemicolon(s string) int {
	var index int
	for i := 0; i < len(s); i++ {
		if s[i] == ':' {
			index = i
			break
		}
	}
	return index
}

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

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
