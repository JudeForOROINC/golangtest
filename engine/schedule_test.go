package engine

import (
	"testing"
)

// //TODO

// type Schedule struct {
// 	tasks       []Task
// 	jobs        map[int]Job
// 	CurrentTime GameTime
// }

// type Task struct {
// 	job Job
// }

// type Job struct {
// 	name   string
// 	action ActionJob
// }

// type Oserver struct {
// 	subjects []interface{}
// }

// type JobType int

// type ActionJob func(d ...interface{})

// const (
// 	growFood = iota
// )

// func FactoryGetJob(t JobType) *Job {
// 	switch t {
// 	case growFood:
// 		Job := Job{name: "GrowFood", action: func(d ...interface{}) {
// 			// for
// 		}}
// 	}
// }

// func TestGrowFood(t *testing.T) {
// 	// if loc.Ltype != farmLand {
// 	// 	return
// 	// }
// 	// itm := w.CreateItem(food, strPointer("food"), intPointer(1), intPointer(1))
// 	// loc.items.AddItem(itm)
// 	GrowFood()
// }

// func (j Job) GrowFood(d ...interface{}) {
// 	world, ok := d[0].(*World)
// 	if !ok {
// 		fmt.Println("wrong argument type")
// 		return
// 	}
// 	for _, v := range world.locations {
// 		GrowFood(world, v)
// 	}
// 	return
// }

// type GameTime struct {
// 	year  int
// 	month int
// 	week  int
// 	day   int
// 	hour  int
// 	min   int
// 	sec   int
// }

func TestGameTimeToInt(t *testing.T) {
	val := 3600
	gt := GameTime{sec: 0, min: 0, hour: 1}
	valRes := gt.GameTimeToInt()
	if val != valRes {
		t.Errorf("TimeToInt fail! Sum was incorrect, got: %d, want: %d.", valRes, val)
	}
}

// func (gt GameTime) AddInterval(i GameTime) *GameTime {
// 	return CreateFromInt(gt.GameTimeToInt() + i.GameTimeToInt())
// }

// func CreateFromInt(i int) *GameTime {
// 	c := i
// 	s := c % 60
// 	c = int(c / 60)
// 	min := c % 60
// 	c = int(c / 60)
// 	h := c % 10
// 	c = int(c / 10)
// 	d := c % 5
// 	c = int(c / 5)
// 	w := c % 5
// 	c = int(c / 5)
// 	m := c % 10
// 	c = int(c / 10)
// 	y := c
// 	return &GameTime{year: y, month: m, week: w, day: d, hour: h, min: min, sec: s}

// }

// func (gt GameTime) SubInterval(i GameTime) *GameTime {
// 	return CreateFromInt(gt.GameTimeToInt() - i.GameTimeToInt())
// }

// func (s *Schedule) SetTask(task Task, start GameTime, interval *GameTime) {
// 	if s.CurrentTime.GameTimeToInt() > start.GameTimeToInt() {
// 		return
// 	}
// 	s.tasks = append(s.tasks, task)
// 	s.jobs[start.GameTimeToInt()] = task.job
// 	if interval != nil {
// 		s.jobs[start.GameTimeToInt()+interval.GameTimeToInt()] = task.job
// 	}
// }

// func (s *Schedule) getJobs(to GameTime) []Job {
// 	if s.CurrentTime.GameTimeToInt() > to.GameTimeToInt() {
// 		return []Job{}
// 	}
// 	jobs := s.jobs //???
// }

// // func (j *Job) GrowTree
