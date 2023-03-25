// search.go
package main

import "strings"

func SearchNotes(notes []Note, query string) []Note {
	var results []Note

	for _, note := range notes {
		if strings.Contains(strings.ToLower(note.name), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(note.description), strings.ToLower(query)) {
			results = append(results, note)
			continue
		}

		for _, tag := range note.tags {
			if strings.Contains(strings.ToLower(tag), strings.ToLower(query)) {
				results = append(results, note)
				break
			}
		}
	}

	return results
}
