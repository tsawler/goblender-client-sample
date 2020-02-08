package clientmodels

import "time"

type Sample struct {
	ID        int
	Name      string
	Phone     string
	CreatedAt time.Time
}
