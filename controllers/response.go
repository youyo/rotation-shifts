package controllers

import "fmt"

type Response struct {
	Response interface{} `json:"response"`
	Error    interface{} `json:"error"`
}

func response(res interface{}, err error) Response {
	r := Response{
		Response: res,
		Error:    err,
	}
	if err != nil {
		r.Error = fmt.Sprintf("%v", err)
	}
	return r
}
