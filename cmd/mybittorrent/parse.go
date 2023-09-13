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

	pieceLength, ok := info["piece length"].(int)

	if !ok {
		return Torrent{}, fmt.Errorf("invalid 'piece length' field'")
	}

	pieces, ok := info["pieces"].(string)

	if !ok {
		return Torrent{}, fmt.Errorf("invalid 'pieces' field'")
	}

	return Torrent{
		Announce: Announce{url: url},
		Info: Info{
			length:      length,
			pieceLength: pieceLength,
			pieces:      pieces,
		},
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
