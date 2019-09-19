package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	NumRoutines = 3
	NumRequests = 1000
)

// global semaphore monitoring the number of routines
var semRout = make(chan int, NumRoutines)

// global semaphore monitoring console
var semDisp = make(chan int, 1)

// Waitgroups to ensure that main does not exit until all done
var wgRout sync.WaitGroup
var wgDisp sync.WaitGroup

type Task struct {
	a, b float32
	disp chan float32
}

var t1 Task


func solve(t *Task) { //A function that sleeps for a random time between 1 and 15 seconds, adds
	//the numbers a and b and sends the result on the display channel.
	
	rand.Seed(time.Now().UTC().UnixNano())
	ranNum := rand.Intn(15) + 1
	time.Sleep(time.Duration(ranNum) * time.Second)
	t.disp <- t.a + t.b

}

func handleReq(t *Task) { // A function that acts as intermediary between ComputeServer and solve.

	solve(t)
	<-semRout
	
}
func ComputeServer() chan *Task { // A function that uses the channel factory pattern
	//(lambda) and listens for requests on the created channel for tasks. It calls the handleReq function.
	var t2 = make(chan *Task)

	
	wgRout.Add(1)
	go func() {

		for {
			sampleTask, ok := <-t2
			if !ok {
				break
			}
			semRout <- 1

			wgRout.Add(1)
			go handleReq(sampleTask)
			
		}

		wgRout.Done()
		
	}()
	
	return t2
}
func DisplayServer() chan float32 { // A function that uses the channel factory pattern
	//(lambda) and listens for requests on the created channel for results to print to the console.

	var displayChan = make(chan float32)
	wgDisp.Add(1)
	go func() {
		

		for {
			output, ok := <-displayChan 
			if !ok {
				break
			}

			fmt.Println("-------------")
			fmt.Println("Result:", output)
			wgRout.Done() // waitgroup being used in compute server is being done after the result has been displayed to the console
		}
		
		wgDisp.Done()

	}()

	return displayChan
}

func main() {
	dispChan := DisplayServer()
	reqChan := ComputeServer()
	for {
		var a, b float32
		//var t Task
		// make sure to use semDisp
		semDisp <- 1
		// …
		fmt.Print("Enter two numbers: ")
		fmt.Scanf("%f %f \n", &a, &b)

		fmt.Printf("%f %f \n", a, b)
		if a == 0 && b == 0 {
			//
			break
		}
		// Create task and send to ComputeServer
		// …
		var t Task
		t.a = a
		t.b = b
		t.disp = dispChan
		reqChan <- &t
		time.Sleep(1e9)
		<-semDisp
	}
	// Don’t exit until all is done
	<-semDisp
	close(reqChan)
	wgRout.Wait()
	close(dispChan)
	wgDisp.Wait()
	close(semRout)
	close(semDisp)

}
