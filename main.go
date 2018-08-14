package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Decode body to a struct
func decode(body io.ReadCloser, object interface{}) error {
	decoder := json.NewDecoder(body)
	defer body.Close()

	return decoder.Decode(&object)
}

func getLoggedIn(r *http.Request) *http.Response {
	client := &http.Client{}
	postData := Credential{
		Application: Application{
			ClientId:     "id",
			ClientSecret: "secret",
		},
	}
	req, _ := http.NewRequest("POST", "http://localhost:1234/validate", postData.ToIoReader())
	for _, cookie := range r.Cookies() {
		req.AddCookie(cookie)
	}
	fmt.Printf("req: %v\n", req)
	response, _ := client.Do(req)
	return response

}

func index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getLoggedIn(r)
	metadata := Metadata{}
	decode(response.Body, &metadata)
	for _, cookie := range response.Cookies() {
		http.SetCookie(w, cookie)
	}

	if metadata.Meta.HttpStatus != 200 {
		t, _ := template.ParseFiles("login.html")
		err := t.Execute(w, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		t, _ := template.ParseFiles("index.html")
		err := t.Execute(w, metadata.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.ServeFiles("/css/*filepath", http.Dir("css"))
	router.ServeFiles("/js/*filepath", http.Dir("js"))

	fmt.Println("Starting front-end...")
	http.ListenAndServe(":3123", router)
}
