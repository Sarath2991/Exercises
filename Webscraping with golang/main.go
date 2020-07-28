// make_http_request.go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	strip "github.com/grokify/html-strip-tags-go"
)

//Wordcount tables struct
type Wordcount struct {
	Key   string
	Value int
}

var tmpl = template.Must(template.ParseGlob("Templates/*"))

//Handler handles the function
func Handler(w http.ResponseWriter, r *http.Request) {
	name := ""
	if r.Method == "POST" {
		name = r.FormValue("URL")
	}
	response, err := http.Get(name)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	dataInBytes, err := ioutil.ReadAll(response.Body)
	stripped := strip.StripTags(string(dataInBytes))

	wordcount := WordCount(stripped)
	Wordcarray := []Wordcount{}

	for key, value := range wordcount {
		Wordcsingle := Wordcount{}
		Wordcsingle.Key = key
		Wordcsingle.Value = value
		Wordcarray = append(Wordcarray, Wordcsingle)
	}

	tmpl.ExecuteTemplate(w, "Show", Wordcarray)

}

//Index is the index of the page
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

//WordCount gets words count
func WordCount(s string) map[string]int {

	words := strings.Fields(s)
	wordCountMap := make(map[string]int)

	for _, word := range words {
		wordCountMap[word]++
	}

	return wordCountMap
}

func main() {
	log.Println("Server started on: http://localhost:80")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Handler)
	http.ListenAndServe(":8000", nil)
}
