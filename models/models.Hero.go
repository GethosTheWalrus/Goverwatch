package models

// Model representing an individual hero
// Fields:
// Name: The hero's name
// Role: The hero's role(s)
// Description: The hero's description/lore
// Difficulty: The hero's difficulty (As displayed on Blizzard's website)
// Videos: The hero's idle animation and any other videos that might be of interest
// Abilities: A slice of Ability objects
type Hero struct {
	Name string
	Role string
	Description string
	Difficulty int
	Videos []string
	Abilities []Ability
}