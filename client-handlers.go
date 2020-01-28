package clienthandlers

import (
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}


func TestProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Worked!")
}