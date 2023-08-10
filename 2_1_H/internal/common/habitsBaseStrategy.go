package common

import (
	"github.com/barkimedes/go-deepcopy"
	"sort"
)

type HabitsBaseStrategy struct {
}

func NewHabitsBaseStrategy() *HabitsBaseStrategy {
	return &HabitsBaseStrategy{}
}

func (m *HabitsBaseStrategy) matchSort(i Individual, data []Individual) []Individual {
	removeIndex := 0
	for index, datum := range data {
		if i.ID == datum.ID {
			removeIndex = index
			continue
		}
		data[index].SetHabitsMatchCnt(i.HabitsMatch(datum.Habits))
	}
	tmpData := deepcopy.MustAnything(data).([]Individual)
	tmpData = append(tmpData[:removeIndex], tmpData[removeIndex+1:]...)
	sort.Slice(tmpData, func(i, j int) bool {
		return tmpData[i].HabitsMatchCnt > tmpData[j].HabitsMatchCnt
	})
	return tmpData
}

func (m *HabitsBaseStrategy) sortTemplate(i Individual, datum Individual) Individual {
	datum.SetHabitsMatchCnt(i.HabitsMatch(datum.Habits))
	return datum
}

func (m *HabitsBaseStrategy) compareTo(data []Individual, i, j int) bool {
	return data[i].HabitsMatchCnt > data[j].HabitsMatchCnt
}
