package parser

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	m "github.com/ytobing/scrapper/models"
)

type Handler struct {
	Service Src
}

func InitHandler() Handler {

	return Handler{
		Service: InitService(),
	}
}

func (a *Handler) ParseURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := m.ParseURL{}
	json.NewDecoder(r.Body).Decode(&req)

	res, err := a.Service.ReadData(req.URL)
	if err != nil {
		m.ErrResponse(w, 500, err)
	}

	m.HTTPResponse(w, res, 200)

}
