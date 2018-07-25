package main


import (
    "fmt"
    "log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/operacion/suma", Suma)

    log.Fatal(http.ListenAndServe(":3002", router))
}

func Suma(w http.ResponseWriter, r *http.Request) {
	

	vars := mux.Vars(r)
	sumandoA := vars["sumandoA"]
	fmt.Fprintln(w, "Todo show:", sumandoA)

}