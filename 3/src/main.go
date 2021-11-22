package main

import (
	"fmt"
	"net/http"
	"log"
)

func greeting(salutation string) string {
	return fmt.Sprintf("<b> %s </b>", salutation);
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%s", greeting("Code.educationRocks"));
    })

	log.Fatal(http.ListenAndServe(":8000", nil))
}
