package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

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
		if rune(bencodedString[0]) == 'd' {
			return decodeDictionary(bencodedString)
		}

	}
	return "", 0, fmt.Errorf("unsupported type")
}

func decodeDictionary(s string) (interface{}, int, error) {
	var dic = make(map[string]interface{})
	totalKeys := 0
	dicStr := s[1:]

	for len(dicStr) > 1 {
		if rune(dicStr[0]) == 'e' {
			break
		}
		key, pos, err := decodeBencode(dicStr)
		if err != nil {
			return nil, 0, err
		}
		dicStr = dicStr[pos:]
		totalKeys += pos
		val, pos, err := decodeBencode(dicStr)
		if err != nil {
			return nil, 0, err
		}
		dic[key.(string)] = val
		dicStr = dicStr[pos:]
		totalKeys += pos
	}
	if rune(dicStr[0]) != 'e' {
		return nil, 0, fmt.Errorf("wrong dictionary encoding")
	}
	if len(dic) == 0 {
		return map[string]interface{}{}, 0, nil
	}

	return dic, totalKeys + 2, nil
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
