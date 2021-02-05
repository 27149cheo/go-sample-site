package main

import (
	"fmt"
	"log"
	"net/http"

	"go-sample-site/version"
)

func myHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(
		w,
		"Version:\t%s\nBuild Number:\t%s\nGit Commit:\t%s",
		version.Version, version.BuildNumber, version.GitCommit,
		)
}

func main() {
	http.HandleFunc("/", myHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
