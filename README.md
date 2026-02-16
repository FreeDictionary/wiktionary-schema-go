# wiktionary-schema-go

Go structures to describe Wiktionary data.

## File Order

Schemas of different languages are stored in
`{lang-code}/schema.go` file.

## Fields

Optional fields are all defined as pointer types (e.g.,
`*string`, `*int`, slice, etc.) with `json:"-,omitempty"` flag.
`string` type means that this field is not expected to be
empty. What's more, fields that can be used as indices in the
database are added with `db:"INDEX"` tags.
