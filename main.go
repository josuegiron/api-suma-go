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
    log.Fatal(http.ListenAndServe(":3001", router))
}

func Suma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := map[string]string{"error": "No se ingreso un numero como sumando..."}
			json.NewEncoder(w).Encode(error)
		return
}

	sumA := r.Form.Get("sumandoA")
	sumandoA, err := strconv.ParseFloat(sumA, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
			error := map[string]string{"error": "No se ingreso un numero como sumando..."}
			json.NewEncoder(w).Encode(error)
			return
	}

	sumB := r.Form.Get("sumandoB")
	sumandoB, err := strconv.ParseFloat(sumB, 64)
	if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			error := map[string]string{"error" : "No se ingreso un numero como sumando..."}
    		json.NewEncoder(w).Encode(error)
			return
	}
	
	w.WriteHeader(http.StatusOK)
	total := sumandoA + sumandoB;
	suma := map[string]float64{"total" : total}
    json.NewEncoder(w).Encode(suma)
}