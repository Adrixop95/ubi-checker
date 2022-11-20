package main

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
)

func TestCorrectTelegrafMessage(t *testing.T) {
	// UBI9 ID
	msg := TelegrafMessage("615bcf606feffc5384e8452e")
	expected := "ubi9/ubi"

	if !strings.Contains(msg, expected) {
		t.Fatalf(`Output does not contain an expected value: %s, output: %s`, expected, msg)
	}
}

func TestIfJsonOutputTelegrafMessage(t *testing.T) {
	msg := TelegrafMessage("")

	log.Println(msg)

	var js interface{}
	if json.Unmarshal([]byte(msg), &js) != nil {
		t.Fatalf(`OutOutput is not valid JSON. Output: %s`, msg)
	}
}
