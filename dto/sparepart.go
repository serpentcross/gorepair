package dto

import (
	"github.com/jackc/pgtype"
)

type Sparepart struct {
	Id        pgtype.UUID `json:"id"`
	Name      string      `json:"name"`
	Available bool        `json:"available"`
	Artikul   string      `json:"artikul"`
}
