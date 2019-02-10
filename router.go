package stocktables

import (
	"github.com/gorilla/mux"
	"github.com/stocktables/stocktables/handler"
	"net/http"
)

func Router() http.Handler {
	root := mux.NewRouter()
	root.HandleFunc("/", handler.Root)
	return root
}
