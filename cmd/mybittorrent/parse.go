package main

import (
	"fmt"
	"os"
)

func parseTorrentFile(path string) (Torrent, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Torrent{}, fmt.Errorf("couldn't read from the file")
	}
	decodedData, _, err := decodeBencode(string(data))

	url, ok := decodedData.(map[string]interface{})["announce"].(string)

	if !ok {
		return Torrent{}, fmt.Errorf("invalid 'announce' field")
	}

	info, err := parseInfo(path)

	if err != nil {
		return Torrent{}, fmt.Errorf("invalid 'info' field")
	}

	length, ok := info["length"].(int)

	if !ok {
		return Torrent{}, fmt.Errorf("invalid 'length' field'")
	}

	return Torrent{
		URL:    url,
		length: int64(length),
	}, err

}

func parseInfo(path string) (map[string]interface{}, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't read from the file")
	}
	decodedData, _, err := decodeBencode(string(data))

	info, ok := decodedData.(map[string]interface{})["info"].(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf("invalid 'info' field")
	}

	return info, err
}

func extractInfoEncoding(path string) (string, error) {
	var infoEncoding string

	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("couldn't read from the file")
	}

	for i := 2; i < len(data); i++ {
		if string(data[i:i+4]) == "info" {
			infoEncoding = string(data[i-2:])
			break

		}
	}
	return infoEncoding, nil
}
