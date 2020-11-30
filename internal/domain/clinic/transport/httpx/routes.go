package httpx

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/schigh-ntwrk/ent-poc/internal/ent"
	"github.com/schigh-ntwrk/ent-poc/internal/httpx"
	"github.com/schigh-ntwrk/ent-poc/internal/services"
)

func Routes(svc services.ClinicService) []httpx.Route {
	return []httpx.Route{
		{
			Path:   "/clinic/{id}",
			Method: http.MethodGet,
			Handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				vars := mux.Vars(r)
				idVar := vars["id"]
				id, parseErr := uuid.Parse(idVar)
				if parseErr != nil {
					_ = httpx.Error(rw, http.StatusUnprocessableEntity, parseErr)
					return
				}
				c, getErr := svc.GetByID(r.Context(), id)
				if ent.IsNotFound(getErr) {
					_ = httpx.Error(rw, http.StatusNotFound, errors.New("the requested item could not be found"))
					return
				}
				if getErr != nil {
					_ = httpx.InternalServerError(rw, getErr)
					return
				}

				_ = httpx.Response(rw).AsJSON().WithInterface(c).Finish()
			}),
		},
		{
			Path:   "/clinic",
			Method: http.MethodPost,
			Handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				data, dataErr := ioutil.ReadAll(r.Body)
				if dataErr != nil {
					_ = httpx.InternalServerError(rw, dataErr)
					return
				}

				var c ent.Clinic
				if jsonErr := json.Unmarshal(data, &c); jsonErr != nil {
					_ = httpx.InternalServerError(rw, jsonErr)
					return
				}

				cNew, saveErr := svc.Save(r.Context(), &c)
				if saveErr != nil {
					_ = httpx.InternalServerError(rw, saveErr)
					return
				}

				_ = httpx.Response(rw).
					AsJSON().
					WithInterface(cNew).
					WithStatusCode(http.StatusCreated).
					Finish()
			}),
		},
		{
			Path:   "/clinic/{id}",
			Method: http.MethodPatch,
			Handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				vars := mux.Vars(r)
				idVar := vars["id"]
				id, parseErr := uuid.Parse(idVar)
				if parseErr != nil {
					_ = httpx.Error(rw, http.StatusUnprocessableEntity, parseErr)
					return
				}

				data, dataErr := ioutil.ReadAll(r.Body)
				if dataErr != nil {
					_ = httpx.InternalServerError(rw, dataErr)
					return
				}

				var c ent.Clinic
				if jsonErr := json.Unmarshal(data, &c); jsonErr != nil {
					_ = httpx.InternalServerError(rw, jsonErr)
					return
				}

				// this is probably redundant, but it's just an example
				c.ID = id

				cUpdated, saveErr := svc.Save(r.Context(), &c)
				if saveErr != nil {
					_ = httpx.InternalServerError(rw, saveErr)
					return
				}

				_ = httpx.Response(rw).AsJSON().WithInterface(cUpdated).Finish()
			}),
		},
	}
}
