package main

import (
	"net/http"
	"svar-widgets/grid-backend-go/data"

	"github.com/go-chi/chi"
)

func initRoutes(r chi.Router, dao *data.DAO) {

	r.Get("/films", func(w http.ResponseWriter, r *http.Request) {
		data, err := dao.Films.GetAll()
		if err != nil {
			format.Text(w, 500, err.Error())
		} else {
			format.JSON(w, 200, data)
		}
	})

	r.Post("/films", func(w http.ResponseWriter, r *http.Request) {
		upd := data.FilmUpdate{}
		var row *data.Film

		err := ParseForm(w, r, &upd)
		if err == nil {
			row, err = dao.Films.Add(upd)
		}

		if err != nil {
			format.Text(w, 500, err.Error())
		} else {
			format.JSON(w, 200, row)
		}
	})

	r.Put("/films/{id}", func(w http.ResponseWriter, r *http.Request) {
		upd := data.FilmUpdate{}
		var row *data.Film

		id := NumberParam(r, "id")
		err := ParseForm(w, r, &upd)
		if err == nil {
			row, err = dao.Films.Update(id, upd)
		}

		if err != nil {
			format.Text(w, 500, err.Error())
		} else {
			format.JSON(w, 200, row)
		}
	})

	r.Patch("/films/{id}", func(w http.ResponseWriter, r *http.Request) {
		p := data.Patch{}

		id := NumberParam(r, "id")
		err := ParseForm(w, r, &p)
		if err == nil {
			err = dao.Films.Patch(id, p)
		}

		if err != nil {
			format.Text(w, 500, err.Error())
		} else {
			format.JSON(w, 200, Response{id})
		}
	})

	r.Delete("/films/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := NumberParam(r, "id")
		err := dao.Films.Delete(id)

		if err != nil {
			format.Text(w, 500, err.Error())
		} else {
			format.JSON(w, 200, Response{id})
		}
	})
}
