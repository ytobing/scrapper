package html

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
}

func InitHandler() Handler {

	return Handler{}
}

func (a *Handler) ServeHTML(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	t.Execute(w, nil)
}
