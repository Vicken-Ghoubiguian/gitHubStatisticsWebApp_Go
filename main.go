//
package main

//
import (
	//"html/template"
	"net/http"
	"serverFunctions"
)

//
func main() {

	//
	http.HandleFunc("/", serverFunctions.HelloServerFunc)
	http.HandleFunc("/user", serverFunctions.GitHubUserFunc)
	http.HandleFunc("/org", serverFunctions.GitHubOrganizationFunc)
	http.HandleFunc("/repos", serverFunctions.GitHubRepositoryFunc)
	http.HandleFunc("/commit", serverFunctions.GitHubCommitfunc)
	http.HandleFunc("/lang", serverFunctions.GitHubLanguagefunc)
	http.HandleFunc("/issue", serverFunctions.GitHubIssuefunc)
	http.HandleFunc("/license", serverFunctions.GitHubLicensefunc)
	http.HandleFunc("/comment", serverFunctions.GitHubCommentfunc)
	http.HandleFunc("/branch", serverFunctions.GitHubBranchfunc)
	http.HandleFunc("/file", serverFunctions.GitHubBranchfunc)

	//
	http.ListenAndServe(":80", nil)
}
