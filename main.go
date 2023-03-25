package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fl0Lec/portfolio/data"
	"github.com/fl0Lec/portfolio/server"
	"gopkg.in/yaml.v3"
)

func main() {
	var d data.Data
	b, _ := os.ReadFile("data/data.yaml")
	yaml.Unmarshal(b, &d)
	fmt.Printf("%+v\n", d)
	for name, list := range d.Skills {
		fmt.Println(name)
		fmt.Println(list)
	}

	mux := server.GetMux()
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(":8081", mux))

}
