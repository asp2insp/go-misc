package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/asp2insp/go-misc/p2p/netutils"
)

func handler(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s: %s\n", r.URL.Path[1:], buf)

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Serving on %s\n", netutils.GetLocalAddress(8080))
	http.ListenAndServe(":8080", nil)
}
