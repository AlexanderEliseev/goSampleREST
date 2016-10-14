package main

import (
	"fmt"
	"net/http"
	//"strconv"
	"strings"
)

var m map[string]string

func APIValueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		pathArr := strings.Split(r.URL.Path, `/`)
		if len(pathArr) != 4 {
			http.NotFound(w, r)
			fmt.Print(w, `Setting value format: "/value/[name]/[value]"`)
		} else {
			w.WriteHeader(http.StatusCreated)
			m[pathArr[2]] = pathArr[3]
			fmt.Print(w, `Value `+pathArr[2]+` writed succeed.`)
		}
	}
	if r.Method == "GET" {
		pathArr := strings.Split(r.URL.Path, `/`)
		if len(pathArr) != 3 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Print(w, `Getting value format: "/value/[name]"`)
		} else {
			value, exist := m[pathArr[2]]
			if exist {
				//w.WriteHeader(http.StatusOK)
				fmt.Print(w, value)
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Print(w, `Value not found`)
			}
		}
	}
	if r.Method == "DELETE" {
		pathArr := strings.Split(r.URL.Path, `/`)
		if len(pathArr) != 3 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Print(w, `Deleting value format: "/value/[name]"`)
		} else {
			_, exist := m[pathArr[2]]
			if exist {
				delete(m, pathArr[2])
				fmt.Print(w, `Deleted succeed.`)
				w.WriteHeader(http.StatusOK)
			} else {
				fmt.Print(w, `Value not found`)
			}
		}
	}
}

func main() {
	m = make(map[string]string)
	http.HandleFunc("/value/", APIValueHandler)
	http.ListenAndServe(":8080", nil)
}
