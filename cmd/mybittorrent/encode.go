package main

import (
	"fmt"
	"strconv"
	"strings"
)

func bencode(data interface{}) (string, error) {

	switch data.(type) {
	case string:
		return encodeString(data.(string)), nil
	case int:
		return encodeInteger(data.(int)), nil
	case []interface{}:
		return encodeList(data.([]interface{}))
	case map[string]interface{}:
		return encodeMap(data.(map[string]interface{}))
	default:
		return "", fmt.Errorf("incompatible type")

	}

}

func encodeString(data string) string {
	return strconv.Itoa(len(data)) + ":" + data
}

func encodeInteger(data int) string {
	return "i" + strconv.Itoa(data) + "e"
}

func encodeList(data []interface{}) (string, error) {
	var encodedData = strings.Builder{}
	encodedData.WriteString("l")
	for _, val := range data {
		tmp, err := bencode(val)
		if err != nil {
			return "", err

		}
		encodedData.WriteString(tmp)
	}
	encodedData.WriteString("e")
	return encodedData.String(), nil
}

func encodeMap(data map[string]interface{}) (string, error) {
	var encodedData = strings.Builder{}
	encodedData.WriteString("d")
	for key, val := range data {
		{
			tmp, err := bencode(key)
			if err != nil {
				return "", err

			}
			encodedData.WriteString(tmp)

			tmp, err = bencode(val)
			if err != nil {
				return "", err

			}
			encodedData.WriteString(tmp)

		}
	}

	encodedData.WriteString("e")

	return encodedData.String(), nil
}
