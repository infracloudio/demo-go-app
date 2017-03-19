package main
 
import(
"fmt"
"net/http"
"os"
"time"
)
 
func indexHandler( w http.ResponseWriter, r *http.Request){
h, _ := os.Hostname()
fmt.Fprintf(w, "hello world, I'm running on pod %s with version %s at %s ", h,version,time.Now().Format(time.UnixDate) )
 
}

func healthHandler( w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "I'm fine Version : %s ", version)

}

const version string = "1.0.1"
 
func main(){
http.HandleFunc("/", indexHandler)
http.HandleFunc("/healthz", healthHandler)
http.ListenAndServe(":8000",nil)
}
