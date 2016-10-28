package handlers

import (
	"app/storage"
	"io/ioutil"
	"net/http"
)

func SetKey(db storage.Database) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key in path", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading put body", http.StatusBadRequest)
			return
		}

		db.Set(key, string(val))
		w.WriteHeader(http.StatusOK)
	})
}
