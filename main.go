/*
Start a server.
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FSItem struct {
	Name string `json:"name"`
}

func main() {
	port := flag.String("port", "8000", "port to run the server on")
	flag.Parse()
	*port = ":" + *port

	handler := http.FileServer(http.Dir("./"))
	http.Handle("/", handler)

	http.HandleFunc("/api/dir", func(w http.ResponseWriter, r *http.Request) {
		dirItems, err := ScanRootItems()
		if err != nil {
			http.Error(w, "an error occured", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string][]FSItem{
			"dir": dirItems,
		})
	})

	log.Println(fmt.Sprintf("started server on port %s", *port))
	log.Fatal(http.ListenAndServe(*port, nil))
}

func ScanRootItems() ([]FSItem, error) {
	fInfos, err := ioutil.ReadDir("./")
	if err != nil {
		return []FSItem{}, err
	}
	var res []FSItem
	for _, f := range fInfos {
		r := FSItem{
			Name: f.Name(),
		}
		res = append(res, r)
	}
	return res, nil
}
