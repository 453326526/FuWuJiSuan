package server

import (
	"net/http"
	"os"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx)

	n.UseHandler(mx)
	return n
}
func unknownImp(rw http.ResponseWriter, r *http.Request) {
	// have not implement yet
	http.Error(rw, "501 Not Implement", http.StatusNotImplemented)
}
func apihandle(w http.ResponseWriter, req *http.Request) {
	rdr := render.New()
	rdr.JSON(w, http.StatusOK, struct {
		ID      string
		Content string
	}{ID: "lm", Content: "what actions do you want to manupulate?"})
}
func showTable(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rdr := render.New()
	err := rdr.HTML(rw, http.StatusOK, "table", struct {
		Un string
		Pw string
		Eg string
		Cg string
		Mg string
	}{Un: r.Form["username"][0], Pw: r.Form["password"][0], Eg: r.Form["englishgrade"][0], Cg: r.Form["chinesegrade"][0], Mg: r.Form["mathgrade"][0]})
	if err != nil {
		panic("render fail!")
	}
}
func initRoutes(mx *mux.Router) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	mx.HandleFunc("/unknown", unknownImp).Methods("GET")
	mx.HandleFunc("/api", apihandle).Methods("GET")
	mx.HandleFunc("/", showTable).Methods("POST")
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

}
