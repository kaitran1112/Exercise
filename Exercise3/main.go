package main

import (
	"fmt"
	"sync"
	"time"
)

type Bakery struct {
	c    chan int
	quit chan int
	free bool
}
type Customer struct {
	name        string
	countOfCake int
}

func (b *Bakery) Grill(wg *sync.WaitGroup, cus *Customer) {
	defer wg.Done()
	id := 1
	for {
		select {
		case b.c <- id:
			id++
		case <-b.quit:
			fmt.Printf("Đã nướng xong bánh cho %v\n", cus.name)
			b.free = true
			return
		}
	}
}
func Action(cus *Customer, b1 *Bakery, b2 *Bakery, wg *sync.WaitGroup) {
	for {
		if b1.free {
			wg.Add(2)
			b1.free = false
			go func() {
				for i := 1; i <= cus.countOfCake; i++ {
					time.Sleep(time.Second)
					fmt.Printf("Bak1 đã nướng xong bánh thứ %d cho %v\n", <-b1.c, cus.name)
				}
				b1.quit <- 0
				wg.Done()
			}()
			go b1.Grill(wg, cus)
			break
		} else if b2.free {
			wg.Add(2)
			b2.free = false
			go func() {
				for i := 1; i <= cus.countOfCake; i++ {
					fmt.Printf(" Bak2 đã nướng xong bánh thứ %d cho %v\n", <-b2.c, cus.name)
				}
				b2.quit <- 0
				wg.Done()
			}()
			go b1.Grill(wg, cus)
			break
		}
	}
	wg.Wait()
}
func main() {
	var wg sync.WaitGroup
	c1 := make(chan int)
	quit1 := make(chan int)
	c2 := make(chan int)
	quit2 := make(chan int)
	Bakery1 := Bakery{c1, quit1, true}
	Bakery2 := Bakery{c2, quit2, true}
	Cus1 := Customer{"Kai", 2}
	Cus2 := Customer{"Palpi", 3}
	Cus3 := Customer{"John", 2}
	Action(&Cus1, &Bakery1, &Bakery2, &wg)
	Action(&Cus2, &Bakery1, &Bakery2, &wg)
	Action(&Cus3, &Bakery1, &Bakery2, &wg)

}
