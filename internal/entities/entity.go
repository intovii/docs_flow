package entities

import "database/sql"

//type User struct {
//	ID    int
//	Mail  string
//	Phone string
//	Role  string
//	Token string
//}

type User struct {
	ID int `json:"id"`
	//UserToken  string `json:"user_token"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Login  string `json:"login"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	//IsApproved bool   `json:"is_approved"`
}

func NewNullString(s string) *sql.NullString {
	if len(s) == 0 {
		return &sql.NullString{}
	} else {
		return &sql.NullString{
			String: s,
			Valid:  true,
		}
	}
}