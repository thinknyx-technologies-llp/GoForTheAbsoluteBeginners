package main
import ("fmt"
	"sync"
	"time"
)

func square_number(number int, wg *sync.WaitGroup){
	fmt.Println(number * number)
	wg.Done()
}

func demo4(){
	var wg sync.WaitGroup
	initial_time := time.Now()
	wg.Add(10)
	for i:=1;i<=10;i++{
		go square_number(i, &wg)
	}
	end_time := time.Since(initial_time)
	wg.Wait()
	fmt.Println(end_time)
}