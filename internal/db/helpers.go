package db

import "database/sql"

func NewNullString(s string) sql.NullString {
	valid := false
	if s != "" {
		valid = true
	}
	return sql.NullString{String: s, Valid: valid}
}

func NewNullInt64(n int64) sql.NullInt64 {
	valid := false
	if n != 0 {
		valid = true
	}
	return sql.NullInt64{Int64: n, Valid: valid}
}
