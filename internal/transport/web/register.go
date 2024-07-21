package web

import "net/http"

func (wh *WebHandler) Register() *http.ServeMux {
	main := http.NewServeMux()
	fsHandler := http.FileServer(http.Dir("./public"))
	main.Handle("/public/", http.StripPrefix("/public/", fsHandler))
	main.HandleFunc("/", wh.IndexHandler)
	return main
}
