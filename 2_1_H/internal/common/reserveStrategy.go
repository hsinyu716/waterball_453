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
	result := m.BaseStrategy.matchSort(i, data)
	return []Individual{
		result[len(result)-1],
	}
}

//func (m *ReverseStrategy) sortTemplate(i Individual, datum Individual) Individual {
//	return m.BaseStrategy.sortTemplate(i, datum)
//}
//
//func (m *ReverseStrategy) compareTo(data []Individual, i, j int) bool {
//	return m.BaseStrategy.compareTo(data, j, i)
//}
