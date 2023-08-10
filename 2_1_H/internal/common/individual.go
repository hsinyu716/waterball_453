package common

import (
	"fmt"
	"math/rand"
	"time"
)

type Individual struct {
	ID             int
	Gender         string
	Age            int
	Intro          string
	Habits         []string
	Coord          Coord
	Distance       float64
	HabitsMatchCnt int
}

func (i *Individual) SetDistance(d float64) {
	i.Distance = d
}

func (i *Individual) SetHabitsMatchCnt(h int) {
	i.HabitsMatchCnt = h
}

func (i *Individual) HabitsMatch(h []string) int {
	m := make(map[string]bool)
	for _, habit := range i.Habits {
		m[habit] = true
	}
	count := 0
	for _, s := range h {
		if m[s] {
			count++
		}
	}
	return count
}

var (
	genders = []string{"MALE", "FEMALE"}
	hobbies = []string{"打籃球", "煮菜", "玩遊戲", "閱讀", "旅遊", "攝影", "寫作", "唱歌", "跳舞", "登山"}
)

func (i *Individual) GenerateRandomIndividual() []Individual {
	rand.Seed(time.Now().UnixNano())

	var data []Individual
	for j := 1; j <= 10; j++ {
		ID := j
		gender := genders[rand.Intn(len(genders))]
		age := rand.Intn(83) + 18 // 18~100歲
		intro := fmt.Sprintf("這是編號 %d 的用戶自我介紹。", ID)
		numHabits := rand.Intn(4) + 2
		habitList := rand.Perm(len(hobbies))[:numHabits]
		var habits []string
		for _, idx := range habitList {
			habits = append(habits, hobbies[idx])
		}
		coordX := float64(j) //rand.Float64() * 100
		coordY := float64(j)
		coord := Coord{}
		coord.Set(coordX, coordY)

		data = append(data, Individual{
			ID:     ID,
			Gender: gender,
			Age:    age,
			Intro:  intro,
			Habits: habits,
			Coord:  coord,
		})
	}

	return data
}

func (i *Individual) match(person Individual) {
	fmt.Println(fmt.Sprintf(`%d 用戶與 %d 為最佳配對；距離為%v, 興趣分數為%v`, i.ID, person.ID, person.Distance, person.HabitsMatchCnt))
}

type MyIndividual struct {
	Individuals []Individual
}

func (m MyIndividual) Len() int {
	return len(m.Individuals)
}

func (m MyIndividual) Less(i, j int) bool {
	return m.Individuals[i].ID < m.Individuals[j].ID
}

func (m MyIndividual) Swap(i, j int) {
	m.Individuals[i], m.Individuals[j] = m.Individuals[j], m.Individuals[i]
}

type MyIndividualDistance struct {
	MyIndividual
}

func (m MyIndividualDistance) Less(i, j int) bool {
	if m.Individuals[i].Distance < m.Individuals[j].Distance {
		return true
	} else if m.Individuals[i].Distance == m.Individuals[j].Distance {
		return m.Individuals[i].ID < m.Individuals[j].ID
	}
	return false
}

type MyIndividualHabits struct {
	MyIndividual
}

func (m MyIndividualHabits) Less(i, j int) bool {
	if m.Individuals[i].HabitsMatchCnt > m.Individuals[j].HabitsMatchCnt {
		return true
	} else if m.Individuals[i].HabitsMatchCnt == m.Individuals[j].HabitsMatchCnt {
		return m.Individuals[i].ID < m.Individuals[j].ID
	}
	return false
}
