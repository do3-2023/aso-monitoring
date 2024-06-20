package handlers

import (
	"apimonitoring/utils"
	"fmt"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func Health(w http.ResponseWriter, _ *http.Request, ps spinhttp.Params) {
	db, err := utils.OpenPostgres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	_, err = db.Query("SELECT 1;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "OK")
}
