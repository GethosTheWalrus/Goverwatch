package models

// Model representing a watered down hero, for the hero list call
// Fields:
// Name: The name of the hero
// Roles: The role(s) of the hero
// Portrait: The hero's portrait (As displayed on Blizzard's website)
type HeroListHero struct {
	Name string
	Roles string
	Portrait string
}