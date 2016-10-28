package handlers

import (
	"app/storage"
	"net/http"
)

func GetKey(db storage.Database) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, res *http.Request) {

		key := res.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key name in query string", http.StatusBadRequest)
			return
		}

		val, err := db.Get(key)

		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	})
}
