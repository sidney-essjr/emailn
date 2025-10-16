package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type EndpoitFunc func(w http.ResponseWriter, r *http.Request) (any, int, map[string]string)

func HandlerError(endpointFunc EndpoitFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, mapErr := endpointFunc(w, r)

		if len(mapErr) > 0 {
			var err string
			for k, v := range mapErr {
				err = fmt.Sprintf("Resource: %s - %s", k, v)
				if v == internalerrors.ErrInternal.Error() {
					render.Status(r, 500)
				} else {
					render.Status(r, 400)
				}
			}
			render.JSON(w, r, err)
			return
		}

		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
