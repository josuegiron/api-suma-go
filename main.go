package main


import (
    "log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)
type sumando struct {
	sumandoA float32
	sumandoB float32
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/operacion/suma", Suma).Methods("GET")
	

    log.Fatal(http.ListenAndServe(":3002", router))
}

func Suma(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s", err)
		return
}

	sumA := r.Form.Get("sumandoA")
	sumandoA, err := strconv.ParseFloat(sumA, 64)
	if err != nil {
			error := map[string]string{"error": "No se ingreso un numero como sumando..."}
			json.NewEncoder(w).Encode(error)
			return
	}

	sumB := r.Form.Get("sumandoB")
	sumandoB, err := strconv.ParseFloat(sumB, 64)
	if err != nil {
			error := map[string]string{"error" : "No se ingreso un numero como sumando..."}
    		json.NewEncoder(w).Encode(error)
			return
	}

	total := sumandoA + sumandoB;
	suma := map[string]float64{"total" : total}
    json.NewEncoder(w).Encode(suma)

}