package funcapisort

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	NewConnection()

	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/conversion", Wrapper(GetConversion, EmptyDecoder)).Methods("GET")

	http.Handle("/", myRoute)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func Wrapper(fn func(interface{}) (interface{}, error), dec func(*http.Request) (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		payload, err := dec(r)
		if err != nil {
			//todo esci
		}

		resp, _ := fn(payload)
		//todo se err ... tornare qualcosa altro

		jsonData, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Scrive il JSON come risposta
		w.Write(jsonData)
	}
}
