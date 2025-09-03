// package main

// import (
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func main() {
// 	dsn := "host=localhost user=postgres password=secret dbname=app port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.AutoMigrate(&model.Product{})
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

// 5つのゴルーチンが3回counterを+1する
func demoMutex() {
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				mu.Lock()
				counter++
				fmt.Printf("goroutine %d incremented counter to %d\n", id, counter)
				mu.Unlock()
				time.Sleep(100 * time.Microsecond)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Final counter =", counter)
}

func demoWaigGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("Worker", id, "start")
			time.Sleep(time.Second)
			fmt.Println("Worker", id, "done")
		}(i)
	}
	wg.Wait()
	fmt.Println("All workers finished")
}

func main() {
	demoMutex()
	demoWaigGroup()
}
