package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/bukalapak/packen/response"
	"github.com/julienschmidt/httprouter"
)

const (
	TOKEN_LIFETIME = 1 * time.Hour
	COOKIE_NAME    = "minisso"
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

func login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	credential := Credential{
		Application: Application{
			ClientId:     "id",
			ClientSecret: "secret",
		},
		User: User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("pass"),
		},
	}
	log.Printf("email %s pass %s\n", r.FormValue("email"), r.FormValue("pass"))
	req, _ := http.NewRequest("POST", "http://localhost:1234/login", credential.ToIoReader())
	for _, cookie := range r.Cookies() {
		req.AddCookie(cookie)
	}
	client := &http.Client{}
	response, _ := client.Do(req)
	for _, cookie := range response.Cookies() {
		http.SetCookie(w, cookie)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

type ErrorBody struct {
	Errors []response.ErrorInfo `json:"errors"`
	Meta   response.MetaInfo    `json:"meta"`
}

func signup(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// simulate HTTP request
	credential := Credential{}
	decode(r.Body, &credential)
	req, _ := http.NewRequest("POST", "http://localhost:1234/signup", credential.ToIoReader())
	for _, cookie := range r.Cookies() {
		req.AddCookie(cookie)
	}

	// do request
	client := &http.Client{}
	resp, _ := client.Do(req)

	// simulate HTTP response
	for _, cookie := range resp.Cookies() {
		http.SetCookie(w, cookie)
	}
	if resp.StatusCode == 200 {
		metadata := Metadata{}
		decode(resp.Body, &metadata)
		res := response.BuildSuccess(metadata, response.MetaInfo{HTTPStatus: metadata.Meta.HttpStatus})
		response.Write(w, res, metadata.Meta.HttpStatus)
	} else {
		errorBody := ErrorBody{}
		decode(resp.Body, &errorBody)
		res := response.BuildSuccess(errorBody, response.MetaInfo{HTTPStatus: errorBody.Meta.HTTPStatus})
		response.Write(w, res, errorBody.Meta.HTTPStatus)
	}

}

func logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	http.SetCookie(w, &http.Cookie{
		Name:   COOKIE_NAME,
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.POST("/login", login)
	router.POST("/signup", signup)
	router.POST("/logout", logout)
	router.ServeFiles("/css/*filepath", http.Dir("css"))
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/images/*filepath", http.Dir("images"))

	fmt.Println("Starting front-end...")
	http.ListenAndServe(":3123", router)
}
