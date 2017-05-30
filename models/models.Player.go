package models

// Model representing a player object
// Fields:
// Name: Player Name (As displayed on Blizzard's website)
// Battletag: Player Battle tag
// Portrait: Player's in-game portrait
// Rank: Player's numerical rank along with their rank icon
// FeaturedStats: List of player's featured stats (As displayed on Blizzard's website)
// CareerStats: List of player's career stats (As displayed on Blizzard's website)
type Player struct {
	Name string
	Battletag string
	Portrait string
	Rank [2]string
	FeaturedStats []Stat
	CareerStats []StatCategory
}