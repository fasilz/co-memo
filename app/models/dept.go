package models

import "fmt"

type Department struct {
	ID    uint32
	Name  string
	Teams []Team
}

func (d *Department) String() string {
	return fmt.Sprintf("Team(%s)", d.Name)
}
