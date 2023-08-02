package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SlackBlock struct {
	Type string `json:"type"`
	Text struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"text"`
}
type SlackBody struct {
	Text   string       `json:"text"`
	Blocks []SlackBlock `json:"blocks"`
}

var SecurityToken = "herpaderp"

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Nothing to see here.\n")
}

func getSlack(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a slack request ", r)

	fmt.Printf("Content length: %d\nBody: %v\n", r.ContentLength, r.Body)

	if r.Method != "POST" || r.ContentLength == 0 || r.Body == nil {
		io.WriteString(w, "Endpoint")
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(" - failed to read body")
		return
	}

	var content SlackBody
	err = json.Unmarshal(bodyBytes, &content)
	if err != nil {
		fmt.Println(" - failed to unmarshal body json")
		return
	}

	fmt.Println("Read body: ", content)
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc(fmt.Sprintf("/slack/%v", SecurityToken), getSlack)

	fmt.Println("Starting listener on port 80")
	err := http.ListenAndServe("127.0.0.1:80", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
