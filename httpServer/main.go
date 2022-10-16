package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
)

func main() {
	fmt.Println("this is a http server exercise")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/healthz", healthyCheck)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" http server exercise ends")

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// for k, v := range r.Form {
	// 	fmt.Println(k, ":", strings.Join(v, ""))
	// }

	fmt.Fprint(w, "hello visitor\n")
	// header := r.Header
	// fmt.Fprintln(w, header)

	io.WriteString(w, "==Details of the http request header:==\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		var str string
		fmt.Println(k)
		for _, n := range v {
			str += n

		}
		fmt.Println(str)
		//w.Header().Add(k, str)
		w.Header().Set("client_"+k, str)

	}
	w.Header().Set("Go-Version", runtime.Version())

	http_code := 200
	w.WriteHeader(http_code)
	fmt.Printf("client info { IP: %s, HTTP_CODE: %d }\n", r.RemoteAddr, http_code)

}
func healthyCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "StatusCode: %d\n", 200)
}
