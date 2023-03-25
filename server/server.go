package server

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/fl0Lec/portfolio/data"
	"gopkg.in/yaml.v3"
)

var (
	tpl *template.Template
	dd  data.Data
)

func init() {
	tpl = template.Must(template.ParseFiles("template/cv.gohtml"))
}

func ServerRawHTML(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	tpl.Execute(w, dd)
}

func GetMux() *http.ServeMux {
	b, err := os.ReadFile("data/data.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(b, &dd)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/cv", ServerRawHTML)

	return mux
}
