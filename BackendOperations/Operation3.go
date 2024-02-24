package backendoperations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID        int    "json:id"
	FirstName string "json:firstname"
	LastName  string "json:lastname"
	Age       int    "json:age"
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

func Op3() {
	// .../api çağrısında
	apiRoot := "/api"
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{Message: "Apı home"}
		output, err := json.Marshal(message)
		checkError(err)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(output))
	})

	//../api/users user verileri çağrısında

	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FirstName: "Necati Efe", LastName: "Hançer", Age: 21},
			User{ID: 2, FirstName: "Sinan", LastName: "Turan", Age: 19},
			User{ID: 3, FirstName: "Ayhan ", LastName: "Düzen", Age: 22},
			User{ID: 4, FirstName: "Arda Baran", LastName: "Aydemir", Age: 21},
			User{ID: 5, FirstName: "Yusuf Bilal", LastName: "Uğur", Age: 22},
		}
		message := users
		outputFile, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(outputFile))
	})

	//..api/me kendi user verisi çağrıldığında
	http.HandleFunc("/api/me", func(w http.ResponseWriter, r *http.Request) {
		me := User{ID: 1, FirstName: "Necati Efe", LastName: "Hançer", Age: 21}
		message := me
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
		fmt.Println("Aha burası çalıştı")
	})

	http.ListenAndServe(":8080", nil)

}
