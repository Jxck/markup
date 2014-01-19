package main

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

const URL string = "https://api.github.com/markdown"

type Message struct {
	Text    string `json:"text"`
	Mode    string `json:"mode"`
	Context string `json:"context"`
}

func Post(text string) string {
	msg := Message{
		Text:    text,
		Mode:    "gfm",
		Context: "github/gollum",
	}

	b, _ := json.Marshal(msg)
	buf := bytes.NewBuffer(b)

	res, _ := http.Post(URL, "text/plain", buf)

	for k, v := range res.Header {
		log.Println(k, v)
	}

	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func MainServer(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./index.html")
}

func RenderServer(ws *websocket.Conn) {
	filename := "./readme.md"
	buf, _ := ioutil.ReadFile(filename)
	text := string(buf)
	rendered := Post(text)
	io.WriteString(ws, rendered)
}

func main() {
	// filename := "./readme.md"
	// buf, _ := ioutil.ReadFile(filename)
	// text := string(buf)
	// rendered := Post(text)
	// fmt.Println(rendered)

	http.HandleFunc("/", MainServer)
	http.Handle("/render", websocket.Handler(RenderServer))
	fmt.Println("server starts at port 3000")
	http.ListenAndServe(":3000", nil)
}
