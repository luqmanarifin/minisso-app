package main

type Meta struct {
	HttpStatus int `json:"http_status"`
}

type Metadata struct {
	Data User `json:"data"`
	Meta Meta `json:"meta"`
}
