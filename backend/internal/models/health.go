package models

import (
	"github.com/google/uuid"
)

type Health struct {
	Id int8      `json:"id" sorm:"primary key;type:uuid;default:gen_random_uuid()"`
	FK uuid.UUID `json:"fk" sorm:"type:uuid;not null"`
}
