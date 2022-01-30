package golangplayground

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	port string = ":80"
)

type APIServer struct {
	sm ServiceManager
}

func (w *APIServer) IndexHandler(rw http.ResponseWriter, r *http.Request) {
	w.sm.Add(r.Host)
	http.ServeFile(rw, r, "./html/index.html")
}
func (w *APIServer) SampleHandler(rw http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("./playground/demo_code.txt")
	fmt.Fprintln(rw, string(data))
}
func (w *APIServer) CompileHandler(rw http.ResponseWriter, r *http.Request) {
	w.sm.Add(r.Host)
	data, _ := ioutil.ReadAll(r.Body)
	playground, _ := w.sm.GetPlayground(r.Host)
	playground.CopySourceCode(data)
	output, _ := playground.RunPlayground()
	fmt.Fprintln(rw, output)
}
func (w *APIServer) FormatHandler(rw http.ResponseWriter, r *http.Request) {
	w.sm.Add(r.Host)
	data, _ := ioutil.ReadAll(r.Body)
	playground, _ := w.sm.GetPlayground(r.Host)
	playground.CopySourceCode(data)
	output, _ := playground.FormatSourceCode()
	fmt.Fprintln(rw, output)
}
func (w *APIServer) CleanupHandler(rw http.ResponseWriter, r *http.Request) {

}
func (w *APIServer) init() {
	http.HandleFunc("/", w.IndexHandler)
	http.HandleFunc("/sample", w.SampleHandler)
	http.HandleFunc("/compile", w.CompileHandler)
	http.HandleFunc("/format", w.FormatHandler)
}
func (w *APIServer) Start() {
	port = os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Println("Server started and listening on:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func NewAPIServer() APIServer {
	api := APIServer{}
	api.init()
	api.sm = NewServiceManager()
	return api
}
