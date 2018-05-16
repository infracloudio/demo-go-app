package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type Instance struct {
	Pod     string
	TS      string
	Version string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	tpl := template.Must(template.New("out").Parse(html))
	i := &Instance{}
	i = newInstance()
	tpl.Execute(w, i)

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	fmt.Fprintf(w, "I'm fine Version : %s ", version)
}

func testoneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	fmt.Fprintf(w, "I'm Test ONE 11111111111")
}

func testtwoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	fmt.Fprintf(w, "I'm Test TWO 2222222222222")
}

func newInstance() *Instance {
	var i = new(Instance)
	i.Pod, _ = os.Hostname()
	i.TS = time.Now().Format(time.UnixDate)
	i.Version = version
	return i
}

const version string = "1.1.1"

func main() {
	fmt.Println("Starting the goapp at port 8000")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/testone", testoneHandler)
	http.HandleFunc("/testtwo", testtwoHandler)
	http.ListenAndServe(":8000", nil)
}
