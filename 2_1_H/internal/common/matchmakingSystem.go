package common

import (
	"github.com/barkimedes/go-deepcopy"
	"sort"
)

type MatchmakingSystem struct {
	BaseStrategy MatchmakingStrategy
}

type IMatchmakingSystem interface {
	matchSort(i Individual, data []Individual) Individual
}

func NewMatchMakingSystem(baseStrategy MatchmakingStrategy) *MatchmakingSystem {
	return &MatchmakingSystem{
		baseStrategy,
	}
}

func (m *MatchmakingSystem) Match(data []Individual) {
	for _, datum := range data {
		individual := m.matchSort(datum, data)
		datum.match(individual)
	}
}

func (m *MatchmakingSystem) matchSort(i Individual, data []Individual) Individual {
	removeSelfIndex := 0
	for index, datum := range data {
		if i.ID == datum.ID {
			removeSelfIndex = index
			continue
		}
		data[index] = m.BaseStrategy.sortValue(i, datum)
	}
	cloneData := deepcopy.MustAnything(data).([]Individual)
	cloneData = append(cloneData[:removeSelfIndex], cloneData[removeSelfIndex+1:]...)
	sort.Slice(cloneData, func(i, j int) bool {
		return m.BaseStrategy.compareTo(cloneData, i, j)
	})
	return cloneData[0]
}
