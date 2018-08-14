package main

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

type Application struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	ClientId        string    `json:"client_id"`
	ClientSecret    string    `json:"client_secret"`
	Description     string    `json:"description" xorm:"text"`
	ApplicationLogo string    `json:"application_logo" xorm:"text"`
	CreatedAt       time.Time `json:"created_at" xorm:"created"`
	UpdatedAt       time.Time `json:"updated_at" xorm:"updated"`
}

func (a *Application) ToIoReader() io.Reader {
	requestByte, _ := json.Marshal(a)
	return bytes.NewReader(requestByte)
}
