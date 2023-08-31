package entities

type Segment struct {
	Id   int    `db:"id"`
	Slug string `db:"slug"`
}
