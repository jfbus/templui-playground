package handlers

import (
	"net/http"
	"strconv"
)

func formValueInt(r http.Request, name string) int {
	v := r.FormValue(name)
	if v == "" {
		return 0
	}
	i, _ := strconv.Atoi(v)
	return i
}
