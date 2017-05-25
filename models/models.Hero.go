package models

type Hero struct {
	Name string
	Role string
	Description string
	Difficulty int
	Videos []string
	Abilities []Ability
}