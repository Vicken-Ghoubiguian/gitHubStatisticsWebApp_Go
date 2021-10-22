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

	//
	http.HandleFunc("/", helloServerFunc)
	http.HandleFunc("/user", gitHubUserFunc)
	http.HandleFunc("/org", gitHubOrganizationFunc)
	http.HandleFunc("/repos", gitHubRepositoryFunc)

	//
	http.ListenAndServe(":80", nil)
}

//
func helloServerFunc(w http.ResponseWriter, r *http.Request) {

	//
	//t := template.New("Label de ma template")

	fmt.Fprintf(w, "Hello, world !")

	//
	//t = template.Must(t.ParseFiles("tmpl/main.tmpl"))
}

//
func gitHubUserFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub user function")
}

//
func gitHubOrganizationFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub organization function")
}

//
func gitHubRepositoryFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub repository function")
}
