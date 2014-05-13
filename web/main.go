package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "strconv"
  "html/template"
  "log"
)

func configLogger() {
  log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err error) {
  http.Error(w, "500: Internal server error", 500)
  log.Print(err)
}

func landingPage(w http.ResponseWriter, r *http.Request) {
  html, err := getHtmlFile("index.html")
  if err != nil {
    http.NotFound(w, r)
    return
  }
  fmt.Fprintf(w, html)
  return
}

func showChatPage(w http.ResponseWriter, r *http.Request) {
  html, err := getHtmlFile("index.html")
  if err != nil {
    http.NotFound(w, r)
    return
  }
  fmt.Fprintf(w, html)
  return
}

func newChatPage(w http.ResponseWriter, r *http.Request) {
  c, err := newChat()
  m := newMessage("Server", "Welcome!")
  c.addMessage(m)
  if err != nil {
    internalServerError(w, r, err)
    return
  }
  path := fmt.Sprintf("/chats/%v", c.Id)
  http.Redirect(w, r, path, 307)
  return
}

func getHtmlFile(filename string) (html string, err error) {
  filename = "views/" + filename
  var fileBytes []byte
  fileBytes, err = ioutil.ReadFile(filename)
  if err != nil {
    fmt.Printf("ERROR couldn't find %s\n", filename)
    return
  }
  html = string(fileBytes)
  return
}

func main() {
  configLogger()
  http.HandleFunc("/", landingPage)
  http.HandleFunc("/chats/", showChatPage)
  http.HandleFunc("/chats/new", newChatPage)
  http.ListenAndServe(":8080", nil)
}
