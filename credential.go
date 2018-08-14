package main

import (
	"bytes"
	"encoding/json"
	"io"
)

type Credential struct {
	Application Application `json:"application"`
	User        User        `json:"user"`
}

func (c *Credential) ToIoReader() io.Reader {
	requestByte, _ := json.Marshal(c)
	return bytes.NewReader(requestByte)
}
