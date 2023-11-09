package examples

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func AcceptStatus(w http.ResponseWriter, r *http.Request) {
	data, _ := httputil.DumpRequest(r, true)
	fmt.Fprint(w, "ok")
	fmt.Println(string(data))
}
