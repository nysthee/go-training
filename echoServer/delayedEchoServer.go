package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Response struct {
	Origin  string              `json:"origin"`
	Method  string              `json:"method"`
	Url     string              `json:"url"`
	Args    map[string][]string `json:"args"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

type echoHandler struct {
	debug bool
}

func (eh *echoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(r.RequestURI)
	params, _ := url.ParseQuery(u.RawQuery)

	// configurable sleep
	delayString := params.Get("delay")
	if delayString != "" {
		delay, err := strconv.Atoi(delayString)
		if err == nil {
			number := rand.Intn(delay * 2)
			time.Sleep(time.Duration(number) * time.Millisecond)
		}
	}

	var body string
	if r.Body != nil {
		asBytes, _ := ioutil.ReadAll(r.Body)
		if asBytes != nil {
			body = string(asBytes)
		}
	}
	resp := Response{
		Origin:  r.RemoteAddr,
		Method:  r.Method,
		Url:     u.Path,
		Args:    params,
		Headers: r.Header,
		Body:    body,
	}

	if eh.debug {
		log.Printf("Server says: %s on %s from %s:%s", r.Method, u.Path, r.RemoteAddr, string(body))
	}

	// Encode json
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// write formatted json body
	var formattedResp bytes.Buffer
	json.Indent(&formattedResp, jsonResp, "", "\t")
	w.Write(formattedResp.Bytes())

}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	mux := http.NewServeMux()

	h := &echoHandler{debug: true}
	mux.Handle("/", h)

	fmt.Printf("Start listening at http://localhost:3000/\n")
	http.ListenAndServe(":3000", mux)
}
