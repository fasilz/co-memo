package models

import "fmt"

type Organization struct {
	ID          uint32
	Name        string
	Departments []Department
}

func (org *Organization) String() string {
	return fmt.Sprintf("Team(%s)", org.Name)
}
