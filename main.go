//
package main

//
import (
	//"html/template"
	"fmt"
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
	//t := template.New("Label de ma template")

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	//
	//t = template.Must(t.ParseFiles("tmpl/main.tmpl"))
}
