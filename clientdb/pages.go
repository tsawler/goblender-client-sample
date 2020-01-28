package clientdb

import (
	"database/sql"
	"github.com/tsawler/goblender/pkg/models"
)

type PageModel struct {
	DB *sql.DB
}


// AllPages returns slice of pages, for use in admin tool
func (m *PageModel) AllPages() ([]*models.Page, error) {
	stmt := "SELECT id, page_title, active, slug, created_at, updated_at FROM pages ORDER BY page_title"

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pages := []*models.Page{}

	for rows.Next() {
		s := &models.Page{}
		err = rows.Scan(&s.ID, &s.PageTitle, &s.Active, &s.Slug, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets.
		pages = append(pages, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pages, nil
}