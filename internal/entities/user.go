package entities

type User struct {
	Id   int    `db:"id"`
	Name string `db:"user_name"`
}
