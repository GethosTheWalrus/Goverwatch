package models

// Model representing an individual stat
// Fields:
// Name: The name of this specific stat
// Value: A player's value for this specific stat
// Icon: The icon for this specific stat (As displayed on Blizzard's website)
type Stat struct {
	Name string
	Value string
	Icon string
}