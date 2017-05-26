package models

type Player struct {
	Name string
	Battletag string
	Portrait string
	Rank [2]string
	FeaturedStats []Stat
	CareerStats []StatCategory
}