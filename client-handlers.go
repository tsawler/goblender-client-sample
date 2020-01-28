package clienthandlers

import (
	"fmt"
	"github.com/tsawler/goblender/client/clienthandlers/clientmodels"
	"net/http"
	"time"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}


func TestProtectedHandler(w http.ResponseWriter, r *http.Request) {
	example := clientmodels.Sample{
		ID:        1,
		Name:      "Jack",
		Phone:     "506-444-5555",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	w.Write([]byte(fmt.Sprintf("It worked, and name is %s", example.Name)))
}