package models

// Model representing a metric by which heroes are compared
// Fields:
// Name: The name of the metric
// Heroes: A list of heros and their values for this metric
type HeroComparisonMetric struct {
	Name string
	Heroes []HeroComparisonHero
}