package main

import (
	"fmt"
	"os"
	"sync"
	"math"
	"time"
	"errors"
)

func main() {

	names := []string{"adit", "kumar", "hello"};
	var wg sync.WaitGroup; // A new Wait Group
	// Concurrency => Independence Parallelism => Simulataneous

	wg.Add(len(names)) // Add number of elements to Wait on to WaitGroup

	for _, name := range names {
		go heavyComputation(name, &wg);
		// Send Reference to Wait Group so calling func can notify when done
		// go -> Keyword Fires a Goroutine .. need to implement Wait
		// so main method waits for all HeavyComputation to finish before exit
	}


	// Args[1] Argument supplied by user
	// Args[0] is the name of the program
	if len(os.Args) >= 3 {
		message, err := fmt.Println(GetGreeting(&os.Args[1], &os.Args[2]))
		if err == nil {
			fmt.Println(message)
		}
	} else {
		fmt.Println("Not enough args")
	}

	// Slice Interation
	for i,lang := range getLangs() {
		fmt.Print(i);
		fmt.Println(" - > " + lang);
	}

	// Struct
	for _, g := range getGophers() {
		g.jump();
	}
}

// Wait Group Reference Not Value so original WaitGroup is Notified when done
func heavyComputation(n string, wg  *sync.WaitGroup) {
	result := 0.0
	for i:=0; i<1000000; i++ {
		// Heavy Computation to Simulate parellelism with goroutines
		// Parallelism  - > Simulataneous
		// Concurrennt - > Independent
		result += math.Pi *  math.Sin(float64(len(n)));
	}
	wg.Done() // Notify the Wait Group that Computation is done


	// for 1 Core -> CPU bound multi threaded Concurrent Doesnt perform Much better
	// For Multiple core -> each thread can run parallel in the cores
	// For IO Bound -> Multiple Threads on perform Better as most Threads are Waiting on IO
	// So Even on Single Core, It performs better

}


func GetGreeting(morning, night *string) (message string, err error) {
	hourofday := time.Now().Hour()
	if morning == nil || night == nil {
		// no need to return values if already given name in Func Signature
		message = "error"
		err = errors.New("No greetings given")
		return
	}
	if hourofday < 7 {
		return *morning, err
	} else {
		return *night, err
	}
}

// var langs []string is a Slice

func getLangs() (langs []string) {
	// Append to Slice -> Adds dynamically if size greater than capacity, Allocate new slice
	langs = append(langs, "go")
	langs = append(langs, "c++")
	return
}

func getGophers() []*gopher {
	gophers := []*gopher{&gopher{"a", 2}, &gopher{"b", 1}}
	return gophers
}

type gopher struct {
	name string
	age  int
} // To Create -> can use Struct Literal ex -> gopher1 := gopher{name, age}

type horse struct {
	weight int
	age    int
}

// Interfaces lists methods, any type implementing these methods also implements this interface
type jumper interface {
	jump() string
}

func getJumpers() []jumper {
	list := []jumper{
		&gopher{"A", 10},
		&horse{10, 20}, // has a Trailing " , "
	} // All Types in the jumper slice must implement the jumper interface


	//Types implements interfaces implicitly, Simply by implementing methods in the interface
	return list
}

func (g *gopher) jump() string {
	// if Method on * gopher it can be called by both pointer to gopher and gopher
	if g.age < 30 {
		return "can jump"
	} else {
		return "cant jump"
	}
}
func (h *horse) jump() string {
	return "can jump"
}




//	go build  to build app with dependencies
//	go build -i to compile all dependencies and store them at $GOPATH/pkg/*/* as .a files.
//	next time go run would be faster if all dependencies are already compiled
//
//	goimports -w main.go -> Write all Necessary imports into file
//	gofmt -w main.go -> Formats correctly and write back the code
//
//	go install builds and saves executable in Gopath/bin -> then run
//	All Packages in Gopath/src -> ex Gopath/src/github.com/aditandadit


