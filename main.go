//
package main

//
import (
	"html/template"
	"net/http"
)

//
func main() {

	http.HandleFunc("/", HelloServer)

	http.ListenAndServe(":80", nil)
}

//
func HelloServer(w http.ResponseWriter, r *http.Request) {

	//
	t := template.New("Label de ma template")

	//
	t = template.Must(t.ParseFiles("tmpl/main.tmpl"))
}
