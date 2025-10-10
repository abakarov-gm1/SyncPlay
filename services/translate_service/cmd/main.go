package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	text := "Привет, как дела?"

	requestBody, err := json.Marshal(map[string]string{
		"q":      text,
		"source": "ru",
		"target": "en",
		"format": "text",
	})

	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:5000/translate", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}

	fmt.Println("Перевод:", result["translatedText"])

}
