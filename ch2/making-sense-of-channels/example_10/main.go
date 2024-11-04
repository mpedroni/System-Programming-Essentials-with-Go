package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	clownChannel := make(chan int, 3)
	clowns := 5

	var wg sync.WaitGroup
	driverCh := make(chan struct{})

	// sender/driver logic
	go func() {
		for clownID := range clownChannel {
			balloon := fmt.Sprintf("Balloon %d", clownID)
			fmt.Printf("Driver: Drove the car with %s inside\n", balloon)
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Driver: Clown finished with %s, the car is ready for more!\n", balloon)
		}
		fmt.Println("Driver: I'm done for the day!")
		close(driverCh)
	}()

	// receiver/clowns logic
	for clown := 1; clown <= clowns; clown++ {
		wg.Add(1)
		go func(clownID int) {
			defer wg.Done()
			balloon := fmt.Sprintf("Balloon %d", clownID)
			fmt.Printf("Clown %d: Hopped into the car with %s\n", clownID, balloon)
			select {
			case clownChannel <- clownID:
				fmt.Printf("Clown %d: Finished with %s\n", clownID, balloon)
			default:
				fmt.Printf("Clown %d: Oops, the car is full, can't fit %s!\n", clownID, balloon)
			}
		}(clown)
	}

	wg.Wait()
	close(clownChannel)
	<-driverCh
	fmt.Println("Circus car ride is over!")
}
