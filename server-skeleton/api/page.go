package api

import (
	"net/http"
	pageHandler "server-skeleton/api/page/handler"
)

func handler() {
	http.HandleFunc("/page", pageHandler.Get)
	http.ListenAndServe(":8080", nil)
}
