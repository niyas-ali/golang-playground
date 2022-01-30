package golangplayground

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	port string = ":80"
)

type APIServer struct {
	sm ServiceManager
}

func (w *APIServer) IndexHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./html/index.html")
}
func (w *APIServer) SampleHandler(rw http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("./playground/demo_code.txt")
	fmt.Fprintln(rw, string(data))
}
func (w *APIServer) CompileHandler(rw http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	playground, _ := w.sm.GetPlayground(GetIP(r))
	playground.CopySourceCode(data)
	output, _ := playground.RunPlayground()
	fmt.Fprintln(rw, output)
}
func (w *APIServer) FormatHandler(rw http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	playground, _ := w.sm.GetPlayground(GetIP(r))
	playground.CopySourceCode(data)
	output, _ := playground.FormatSourceCode()
	fmt.Fprintln(rw, output)
}
func (w *APIServer) CleanupHandler(rw http.ResponseWriter, r *http.Request) {

}
func GetIP(r *http.Request) string {
	return strings.Split(r.RemoteAddr, ":")[0]
}
func (server *APIServer) Handler(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.sm.Add(GetIP(r))
		next(w, r)
	})
}
func (w *APIServer) init() {
	http.Handle("/", w.Handler(w.IndexHandler))
	http.Handle("/sample", w.Handler(w.SampleHandler))
	http.Handle("/compile", w.Handler(w.CompileHandler))
	http.Handle("/format", w.Handler(w.FormatHandler))
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
