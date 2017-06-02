package models

// Model representing a category of stats (e.g. offense, deaths, etc.)
// Fields:
// Name: The name of the category
// Icon: The icon of the category (As displayed on Blizzard's website)
// Stats: A slice of Stat objects
type StatCategory struct { 
	Name string
	Icon string
	Stats []Stat 
}