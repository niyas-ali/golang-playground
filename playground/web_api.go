package golangplayground

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	port string = ":80"
)

type APIServer struct {
	sm ServiceManager
}

func (w *APIServer) IndexHandler(rw http.ResponseWriter, r *http.Request) {
	w.sm.Add(r.RemoteAddr)
	http.ServeFile(rw, r, "./html/index.html")
}

func (w *APIServer) CompileHandler(rw http.ResponseWriter, r *http.Request) {
	w.sm.Add(r.RemoteAddr)
	data, _ := ioutil.ReadAll(r.Body)
	playground, _ := w.sm.GetPlayground(r.RemoteAddr)
	playground.CopySourceCode(data)
	output, _ := playground.RunPlayground()
	fmt.Fprintln(rw, output)
}
func (w *APIServer) FormatHandler(rw http.ResponseWriter, r *http.Request) {
	w.sm.Add(r.RemoteAddr)
	data, _ := ioutil.ReadAll(r.Body)
	playground, _ := w.sm.GetPlayground(r.RemoteAddr)
	playground.CopySourceCode(data)
	output, _ := playground.FormatSourceCode()
	fmt.Fprintln(rw, output)
}
func (w *APIServer) CleanupHandler(rw http.ResponseWriter, r *http.Request) {

}
func (w *APIServer) init() {
	http.HandleFunc("/", w.IndexHandler)
	http.HandleFunc("/compile", w.CompileHandler)
	http.HandleFunc("/format", w.FormatHandler)
}
func (w *APIServer) Start() {
	log.Println("Server started and listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
func NewAPIServer() APIServer {
	api := APIServer{}
	api.init()
	api.sm = NewServiceManager()
	return api
}
