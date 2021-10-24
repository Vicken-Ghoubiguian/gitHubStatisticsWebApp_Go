//
package serverFunctions

import (
	"fmt"
	"net/http"
)

//
func HelloServerFunc(w http.ResponseWriter, r *http.Request) {

	//
	//t := template.New("Label de ma template")

	fmt.Fprintf(w, "Hello, world !")

	//
	if r.URL.Path != "/" {

		http.NotFound(w, r)

		fmt.Fprintf(w, "Error ? Yeah....")
	}

	//
	//t = template.Must(t.ParseFiles("tmpl/main.tmpl"))
}

//
func GitHubUserFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub user function")
}

//
func GitHubOrganizationFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub organization function")
}

//
func GitHubRepositoryFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub repository function")
}

//
func GitHubCommitfunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub commit function")
}

//
func GitHubLanguagefunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub language function")
}

//
func GitHubIssuefunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub issue function")
}

//
func GitHubLicensefunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub license function")
}

//
func GitHubCommentfunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub comment function")
}

//
func GitHubBranchfunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub branch function")
}

//
func GitHubFilefunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "GitHub file function")
}
