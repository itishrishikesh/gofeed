package controller

import "github.com/itishrishikesh/gofeed/internal/database"

// ApiConfig is a type shared between all the handlers.
// This just contains database object.
type ApiConfig struct {
	DB *database.Queries
}
