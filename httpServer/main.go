package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("this is a http server exercise")
	http.HandleFunc("/", helloHandler)
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
		w.Header().Add(k, str)

	}

}
