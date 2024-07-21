package web

import (
	"context"
	"gostack/static/templates"
	"net/http"
)

func (wh *WebHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	indextemple := templates.Index()
	maintemnple := templates.Layout(indextemple)

	maintemnple.Render(context.Background(), w)
}
