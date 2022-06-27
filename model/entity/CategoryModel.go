package entity

import "database/sql"

type CategoryModel struct {
	ID   sql.NullInt64
	Name sql.NullString
}
