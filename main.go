// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
  "os"
)

var port string
var addr *string
var homeTempl = template.Must(template.ParseFiles("views/chat.html"))

func parseCli() {
  if len(os.Args) <= 1 {
    port = "8080"
	} else {
    port = os.Args[1] 
  }
  addr = flag.String("addr", ":" + port, "http service address")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method nod allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}

func sourceHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
  parseCli()
	flag.Parse()
	go h.run()
	http.HandleFunc("/", serveHome)
  http.HandleFunc("/js/", sourceHandler)
  http.HandleFunc("/css/", sourceHandler)
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
