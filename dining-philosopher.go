package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	id              int
}

func (p Philo) eat(sem chan int) {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ", p.id)
		fmt.Println("finished eating ", p.id)

		// time.Sleep(2 * time.Second)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	<-sem // remove a 1 from sem, allowing another to proceed
}

func main() {

	sem := make(chan int, 2) // capacity 2
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}

	k := 0
	for {
		sem <- 1 // will block if there are more than two 1s in sem
		fmt.Println("***")
		go philos[k].eat(sem)
		k = (k + rand.Intn(5)) % 5
	}
}
