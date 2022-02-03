package schema

import (
	"entgo.io/bug/sid"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Other("id", sid.ID("")).SchemaType(map[string]string{
			dialect.MySQL:    "bigint",
			dialect.Postgres: "bigint",
			dialect.SQLite:   "integer",
		}).Default(sid.New).Unique().Immutable(),
		field.Int("age"),
		field.String("name"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
