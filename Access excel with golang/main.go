// make_http_request.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var tmpl = template.Must(template.ParseGlob("Templates/*"))

//Handler handles the function
func Handler(w http.ResponseWriter, r *http.Request) {
	var start string
	var rows, columns int
	var rowerr, colerr error
	if r.Method == "POST" {
		start = r.FormValue("start")
		rows, rowerr = strconv.Atoi(r.FormValue("rows"))
		if rowerr != nil {
			fmt.Println("Please enter valid number")
			w.Write([]byte("Please enter valid number of rows!"))
			return
		}
		columns, colerr = strconv.Atoi(r.FormValue("columns"))
		if colerr != nil {
			fmt.Println("Please enter valid number")
			w.Write([]byte("Please enter valid number of columns!"))
			return
		}
	}
	f, err := excelize.OpenFile("Test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	Columnnumber, starterr := excelize.ColumnNameToNumber(start)
	if starterr != nil {
		w.Write([]byte("Please enter valid character to start!"))
		return
	}
	fmt.Println(Columnnumber)
	totalval := Columnnumber + (rows * columns)

	var Displayvalues []string

	rows2, err := f.Rows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}
	for rows2.Next() {
		columns, _ := rows2.Columns()
		for i := Columnnumber; i < len(columns); i++ {
			if i == totalval {
				break
			} else {
				Displayvalues = append(Displayvalues, columns[i])
			}
		}
		break
	}

	tmpl.ExecuteTemplate(w, "Show", Displayvalues)

}

//Index is the index of the page
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func main() {
	log.Println("Server started on: http://localhost:80")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Handler)
	http.ListenAndServe(":8000", nil)
}
