package clienthandlers

import (
	"database/sql"
	"fmt"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/client/clienthandlers/clientmodels"
	"github.com/tsawler/goblender/pkg/config"
	"net/http"
	"time"
)

func NewClientPageHandlers(conn *sql.DB, app config.AppConfig) {
	pageModel = &clientdb.PageModel{DB: conn}
}

func TestDB(w http.ResponseWriter, r *http.Request) {
	results, err := pageModel.AllPages()
	if err != nil {
		fmt.Println(err)
	}
	for _, x := range results {
		fmt.Println("Page:", x.Slug)
	}
}

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