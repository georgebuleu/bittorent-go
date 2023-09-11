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

	info, ok := decodedData.(map[string]interface{})["info"].(map[string]interface{})

	if !ok {
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
