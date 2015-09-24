package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"strings"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
	roomname string
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathsegments := strings.Split(r.URL.Path, "/")
	
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
		"Room": pathsegments[len(pathsegments) - 1],  //how to pass the proper endpoint here if (room, room2)
	}
	
	//set up go backend to handle chat and socket
	_ , check := openrooms[data["Room"].(string)] 

	// := initializes the variable first
	if !check {newroom := newRoom()
	 http.Handle("/" + pathsegments[len(pathsegments) - 1], newroom)
	 openrooms[pathsegments[len(pathsegments) - 1]] = newroom	

	go newroom.run()
	 }
	
	//Sets up front end at baseURL + unique room name
	t.templ.Execute(w, data)



}

//global variables
var host = flag.String("host", ":8080", "The host of the application.")
var openrooms = make(map[string]*room) 

func main() {

	flag.Parse() // parse the flags

	//room 1
	http.Handle("/chat/", &templateHandler{filename: "chat.html"})

	

	// start the web server
	log.Println("Starting web server on", *host)
	if err := http.ListenAndServe(*host, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
