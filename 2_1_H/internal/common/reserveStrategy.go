package common

type ReverseStrategy struct {
	BaseStrategy MatchmakingStrategy
}

func NewReverseStrategy(baseStrategy MatchmakingStrategy) *ReverseStrategy {
	return &ReverseStrategy{
		baseStrategy,
	}
}

func (m *ReverseStrategy) matchSort(i Individual, data []Individual) []Individual {
	return []Individual{}
}

func (m *ReverseStrategy) sortValue(i Individual, datum Individual) Individual {
	return m.BaseStrategy.sortValue(i, datum)
}

func (m *ReverseStrategy) compareTo(data []Individual, i, j int) bool {
	return m.BaseStrategy.compareTo(data, j, i)
}
