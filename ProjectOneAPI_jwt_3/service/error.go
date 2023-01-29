package service

import (
	"fmt"
	"net/http"
)

func BError(w *http.ResponseWriter, str string, c bool) {
	if !c {
		fmt.Fprintf(*w, str)
	}
}

func EError(w *http.ResponseWriter, str string, err error) {
	if err != nil {
		fmt.Fprintf(*w, fmt.Sprintf("%s: %v", str, err))
	}
}
