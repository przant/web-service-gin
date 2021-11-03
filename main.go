package main

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artits"`
	Price  float64 `json:"Price"`
}
