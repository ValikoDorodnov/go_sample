package entity

type GreetingEntity struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
