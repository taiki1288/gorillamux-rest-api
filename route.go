package main

import (
	"encoding/json"
)

type Post struct {
	Id    int `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}