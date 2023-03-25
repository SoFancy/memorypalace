package main

type MemoryPalace struct {
	notes   []Note
	storage Storage
}

func NewMemoryPalace(storage Storage) *MemoryPalace {
	mp := &MemoryPalace{
		notes:   []Note{},
		storage: storage,
	}

	mp.Load()

	return mp
}

func (mp *MemoryPalace) Load() {
	notes, err := mp.storage.LoadNotes()
	if err == nil {
		mp.notes = notes
	}
}

func (mp *MemoryPalace) Save() error {
	return mp.storage.SaveNotes(mp.notes)
}

func (mp *MemoryPalace) Add(name, description string, tags []string) {
	mp.notes = append(mp.notes, NewNote(name, description, tags))
}

func (mp *MemoryPalace) Search(query string) []Note {
	return SearchNotes(mp.notes, query)
}
