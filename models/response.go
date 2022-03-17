package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWriter  http.ResponseWriter
}

// Crear un response por defecto
func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWriter:  rw,
		contentType: "application/json",
	}
}

func (resp *Response) NoFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource not found"
}

func (resp *Response) Send() {
	resp.respWriter.Header().Set("Content-Type", resp.contentType)
	resp.respWriter.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWriter, string(output))
}

func (resp *Response) UnproceessableEntity()  {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity"
}

func SendData(rw http.ResponseWriter, data interface{})  {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func SendNotFound(rw http.ResponseWriter)  {
	response := CreateDefaultResponse(rw)
	response.NoFound()
	response.Send()
}

func SendUnprocessableEntity(rw http.ResponseWriter)  {
	response := CreateDefaultResponse(rw)
	response.UnproceessableEntity()
	response.Send()
}

