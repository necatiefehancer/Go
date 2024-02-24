package backendoperations

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func loadFile(filepath string) (string, error) {
	bytesData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(bytesData), nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	pageData, _ := loadFile("./BackendOperations/page.html")
	fmt.Fprintf(w, pageData)
}

func Op5() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)

}
