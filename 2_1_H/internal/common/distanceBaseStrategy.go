package common

import (
	"github.com/barkimedes/go-deepcopy"
	"sort"
)

type DistanceBaseStrategy struct {
}

func NewDistanceBaseStrategy() *DistanceBaseStrategy {
	return &DistanceBaseStrategy{}
}

func (m *DistanceBaseStrategy) matchSort(i Individual, data []Individual) []Individual {
	removeIndex := 0
	for index, datum := range data {
		if i.ID == datum.ID {
			removeIndex = index
			continue
		}
		data[index].SetDistance(i.Coord.Distance(datum.Coord))
	}
	tmpData := deepcopy.MustAnything(data).([]Individual)
	tmpData = append(tmpData[:removeIndex], tmpData[removeIndex+1:]...)
	sort.Slice(tmpData, func(i, j int) bool {
		return tmpData[i].Distance < tmpData[j].Distance
	})
	return tmpData
}

func (m *DistanceBaseStrategy) sortTemplate(i Individual, datum Individual) Individual {
	datum.SetDistance(i.Coord.Distance(datum.Coord))
	return datum
}

func (m *DistanceBaseStrategy) compareTo(data []Individual, i, j int) bool {
	return data[i].Distance < data[j].Distance
}
