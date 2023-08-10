package common

type MatchmakingStrategy interface {
	matchSort(i Individual, data []Individual) []Individual
	// template method
	sortValue(i Individual, datum Individual) Individual
	compareTo(data []Individual, i, j int) bool
}
