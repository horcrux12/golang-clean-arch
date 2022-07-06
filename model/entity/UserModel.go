package entity

import "database/sql"

type UserModel struct {
	ID         sql.NullInt64
	Username   sql.NullString
	Password   sql.NullString
	FirstName  sql.NullString
	LastName   sql.NullString
	Locale     sql.NullString
	UserSecret sql.NullString
}
