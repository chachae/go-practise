package restful

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestRestfulApi(t *testing.T) {
	openApi()
}

func openApi() {
	http.HandleFunc("/ping", pong)
	log.Println("Starting http server ...")
	log.Fatal(http.ListenAndServe(":50052", nil))
}

func pong(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		fmt.Printf("/pong interface has an error, %s\n", err.Error())
		return
	}
}
