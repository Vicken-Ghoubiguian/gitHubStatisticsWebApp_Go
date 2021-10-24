//
package serverFunctions

import (
	"fmt"
	"net/http"
)

//
func RootServerFunc(w http.ResponseWriter, r *http.Request) {

	//
	if r.URL.Path != "/" {

		fmt.Fprintf(w, "Error ? Yeah....")

		//
	} else {

		fmt.Fprintf(w, "Hello, world !")
	}
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
