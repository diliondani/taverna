package tavernawebportal

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tavernawebportal/templates/*.gohtml"))
}

//RunWebPortal starts running the dino web portal on address addr
func RunWebPortal(addr string) error {

	http.HandleFunc("/", roothandler)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("tavernawebportal/public"))))
	return http.ListenAndServe(addr, nil)
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
	fmt.Fprintf(w, "Welcome to the Taverna web portal %s", r.RemoteAddr)
}
