package main

type Torrent struct {
	Announce
	Info
}

type Announce struct {
	url string
}

type Info struct {
	length      int
	name        string
	pieceLength int
	pieces      string
}
