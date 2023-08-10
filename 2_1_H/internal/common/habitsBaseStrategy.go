package common

type HabitsBaseStrategy struct {
}

func NewHabitsBaseStrategy() *HabitsBaseStrategy {
	return &HabitsBaseStrategy{}
}

func (m *HabitsBaseStrategy) matchSort(i Individual, data []Individual) []Individual {
	return []Individual{}
}

func (m *HabitsBaseStrategy) sortValue(i Individual, datum Individual) Individual {
	datum.SetHabitsMatchCnt(i.HabitsMatch(datum.Habits))
	return datum
}

func (m *HabitsBaseStrategy) compareTo(data []Individual, i, j int) bool {
	if data[i].HabitsMatchCnt > data[j].HabitsMatchCnt {
		return true
	} else if data[i].HabitsMatchCnt == data[j].HabitsMatchCnt {
		return data[i].ID < data[j].ID
	}
	return false
}
