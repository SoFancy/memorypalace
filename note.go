// note.go
package main

import "strings"

type Note struct {
	name        string
	description string
	tags        []string
}

func NewNote(name, description string, tags []string) Note {
	return Note{
		name:        name,
		description: description,
		tags:        tags,
	}
}

func (n Note) String() string {
	return n.name + ": " + n.description + " [" + strings.Join(n.tags, ", ") + "]"
}
