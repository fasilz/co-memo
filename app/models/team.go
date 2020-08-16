package models

import "fmt"

// Team ...
type Team struct {
	ID      uint32
	Name    string
	Members []User
}

func (t *Team) String() string {
	return fmt.Sprintf("Team(%s)", t.Name)
}
