package tavernawebportal

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tavernawebportal/templates/*.gohtml"))
}

//RunWebPortal starts running the dino web portal on address addr
func RunWebPortal(addr string) error {

	http.HandleFunc("/", roothandler)
	http.HandleFunc("/public/", fileServer)
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/manifest.json", manifest)
	http.HandleFunc("/browserconfig.xml", browserConfig)
	//http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("tavernawebportal/public"))))
	return http.ListenAndServe(addr, nil)
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age:604800")

	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}

	fmt.Fprintf(w, "Welcome to the Taverna web portal %s", r.RemoteAddr)
}

func fileServer(w http.ResponseWriter, r *http.Request) {

	name := "tavernawebportal" + r.URL.Path
	log.Println(name)
	file, err := os.Open(name)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Unable to open and read file : %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	w.Header().Set("Cache-Control", "public, max-age:604800")
	http.ServeContent(w, r, name, time.Now(), file)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	name := "tavernawebportal/public/favicon.ico"
	file, err := os.Open(name)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Unable to open and read file : %v", err), http.StatusNotFound)
		return
	}
	defer file.Close()
	w.Header().Set("Cache-Control", "public, max-age:31536000")
	http.ServeContent(w, r, name, time.Now(), file)
}

func manifest(w http.ResponseWriter, r *http.Request) {
	name := "tavernawebportal/public/manifest.json"
	file, err := os.Open(name)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Unable to open and read file : %v", err), http.StatusNotFound)
		return
	}
	defer file.Close()
	w.Header().Set("Cache-Control", "public, max-age:31536000")
	http.ServeContent(w, r, name, time.Now(), file)
}

func browserConfig(w http.ResponseWriter, r *http.Request) {
	name := "tavernawebportal/public/browserconfig.xml"
	file, err := os.Open(name)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Unable to open and read file : %v", err), http.StatusNotFound)
		return
	}
	defer file.Close()
	w.Header().Set("Cache-Control", "public, max-age:31536000")
	http.ServeContent(w, r, name, time.Now(), file)
}
