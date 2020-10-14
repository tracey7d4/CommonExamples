package main

import (
	"fmt"
	"sync"
	"time"
)
// These constants can be changed, so the problem is scalable.
const (
	noP             = 5 //number of Philosophers
	noCoursesPerP   = 3 // Number of courses per Philosophers
	noPEatAtOneTime = 2 // number of Philosophers eating concurrently
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id                int
	leftCS            *Chopstick
	rightCS           *Chopstick
	permission        <-chan bool
	notifyHostIAmDone chan<- int
	courses           int
}

func (p *Philosopher) eat(wg *sync.WaitGroup) {
	for p.courses > 0 {
		// get permission from Host
		<-p.permission
		// lock the resources
		p.leftCS.Lock()
		p.rightCS.Lock()
		// process: eating --- finishing
		fmt.Printf("P %v started eating\n", p.id)
		p.courses--
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("P %v finished eating. Remaining courses %v\n", p.id, p.courses)
		//unlock the resources
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		//notify Host
		p.notifyHostIAmDone <- p.id
	}
	// finished all course
	wg.Done()
}

//status of a Philosopher
type Status struct {
	eating           bool
	courses          int
	eatingPermission chan<- bool
}
type Host struct {
	totalCourses    int //so host knows when to stop while all P have finished all courses
	PStatus         map[int]Status
	noEating        int
	pFinishedEating <-chan int
}

func (h *Host) coordinate(wg *sync.WaitGroup) {
	for h.totalCourses > 0 {
		select {
		case id := <-h.pFinishedEating:
			m := h.PStatus[id]
			m.eating = false
			h.PStatus[id] = m
			h.totalCourses--
			h.noEating--
		default:
			h.letThemEat()
		}
	}
	wg.Done()
}

func (h *Host) letThemEat() {
	for id, pStatus := range h.PStatus {
		// if P is not eating, and he didn't reach his max courses, and number of people eating at the same time is less than limit
		if !pStatus.eating && pStatus.courses > 0 && h.noEating <= noPEatAtOneTime {
			// check left and right are not eating
			lId := (id - 2 + noP) % noP
			rId := id % noP
			if !h.PStatus[lId].eating && !h.PStatus[rId].eating {
				pStatus.eating = true
				pStatus.courses--
				pStatus.eatingPermission <- true
				h.noEating++
				h.PStatus[id] = pStatus
			}
		}
	}
}

func main() {
	cs := make([]*Chopstick, noP)
	for i := 0; i < noP; i++ {
		cs[i] = new(Chopstick)
	}


	done := make(chan int)
	p := make([]*Philosopher, noP)
	status := make(map[int]Status)
	for i := 0; i < noP; i++ {
		permissionFromHost := make(chan bool)
		p[i] = &Philosopher{
			id:                i+1,
			leftCS:            cs[i],
			rightCS:           cs[(i+1)%noP],
			permission:        permissionFromHost,
			notifyHostIAmDone: done,
			courses:           noCoursesPerP,
		}
		status[i+1] = Status{
			eating:           false,
			courses:          noCoursesPerP,
			eatingPermission: permissionFromHost,
		}
	}
	host := Host{
		totalCourses:    noCoursesPerP * noP,
		PStatus:         status,
		noEating:        0,
		pFinishedEating: done,
	}
	var wg sync.WaitGroup

	for i := 0; i < noP; i++ {
		wg.Add(1)
		go p[i].eat(&wg)
	}

	wg.Add(1)
	go host.coordinate(&wg)

	wg.Wait()
}
