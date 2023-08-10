package common

type DistanceBaseStrategy struct {
}

func NewDistanceBaseStrategy() *DistanceBaseStrategy {
	return &DistanceBaseStrategy{}
}

func (m *DistanceBaseStrategy) matchSort(i Individual, data []Individual) []Individual {
	return []Individual{}
}

func (m *DistanceBaseStrategy) sortValue(i Individual, datum Individual) Individual {
	datum.setDistance(i.Coord.Distance(datum.Coord))
	return datum
}

func (m *DistanceBaseStrategy) compareTo(data []Individual, i, j int) bool {
	if data[i].Distance < data[j].Distance {
		return true
	} else if data[i].Distance == data[j].Distance {
		return data[i].ID < data[j].ID
	}
	return false
}
