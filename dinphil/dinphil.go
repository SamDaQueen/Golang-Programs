package main

import (
	"fmt"
	"slices"
	"sync"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	number int
	leftCS, rightCS *ChopStick
}

func (p Philosopher) Eat(host *Host, wg *sync.WaitGroup)	{
	// get permission from host if possible
	if host.GetPermission(p.number) {
		for i := 0; i < 3; i++ {
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Printf("starting to eat %d\n", p.number)
			fmt.Printf("finished eating %d\n", p.number)
			p.leftCS.Unlock()
			p.rightCS.Unlock()
			
		}
	}
	// release permission after eating
	go host.ReleasePermission(p.number)
	wg.Done()
}

type Host struct {
	eating []int
}

func (h Host) GetPermission(idx int) bool {
	var permission bool
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		if len(h.eating) < 2 && !slices.Contains(h.eating, idx) {
		h.eating = append(h.eating, idx)
		permission = true
		wg.Done()
		return
		}
		fmt.Printf("Please wait! Philosophers %d and %d are eating",
							h.eating[0], h.eating[1])
		permission = false
		wg.Done()
	}()
	wg.Wait()
	return permission
}

func (h Host) ReleasePermission(idx int)	{
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i, num := range h.eating {
			if num == i {
				h.eating = append(h.eating[:i], h.eating[i+1:]...)
			}
		}
		wg.Done()
	}()
	wg.Wait()
	
}

func main() {

	// initialize chopsticks and philosophers
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5 ; i++ {
		chopSticks[i] = new(ChopStick)
	}

	// assign left and right chopsticks to philosopher
	philosophers := make([]*Philosopher, 5)
	for i := 0; i <5; i++ {
		philosophers[i] = &Philosopher{
			i+1, chopSticks[i], chopSticks[(i+1)%5],
		}
	}

	// initialize the host for managing permissions
	host := new(Host)

	fmt.Print("initialized host and going to dine")

	var wg sync.WaitGroup
	wg.Add(5)
	// start the dining
	for i := 0; i <5; i++ {
		go philosophers[i].Eat(host, &wg)
	}
	wg.Wait()
}
