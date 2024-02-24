package backendoperations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APİ struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	messageConcat := "Kullanıcı ID " + id
	message := APİ{Message: messageConcat}
	outputFile, err := json.Marshal(message)
	checkError(err)
	fmt.Fprintf(w, string(outputFile))
}

func Op4() {

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/user/{id:[0-9]+}", Hello)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)

}
