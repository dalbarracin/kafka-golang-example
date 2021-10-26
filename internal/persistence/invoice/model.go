package invoice

import "time"

type Invoice struct {
	Id        int
	Message   string
	CreatedAt time.Time
}
