package main 

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	} 
}
func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//initalizate 2 workers
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	// send the jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	//close the channel indicates that no more data will be send
	close(jobs)
	
	//get the results of work
	for i := 1;  i <= numJobs ; i++ {
		res := <- results
		fmt.Println("Result job",i,res)
	}

	close(results)
}