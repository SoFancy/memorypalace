package main

import (
	"fmt"
	"os"
)

func main() {
	storage, err := NewJSONStorage("memory_palace.json")
	if err != nil {
		fmt.Println("Error loading storage:", err)
		os.Exit(1)
	}

	mp := NewMemoryPalace(storage)

	mp.Add("Socrates", "Greek philosopher", []string{"philosophy", "ancient"})
	mp.Add("Plato", "Student of Socrates", []string{"philosophy", "ancient", "student"})

	err = mp.Save()
	if err != nil {
		fmt.Println("Error saving notes:", err)
	}

	searchResults := mp.Search("philosophy")

	for _, note := range searchResults {
		fmt.Println(note)
	}
}
