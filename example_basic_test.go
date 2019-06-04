package parallel_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/NoUseFreak/go-parallel"
)

type Payload struct {
	Number int
}

func Example_basic() {
	input := parallel.Input{}

	for i := 0; i < 20; i++ {
		input = append(input, Payload{Number: i})
	}

	p := parallel.Processor{Threads: 5}
	result := p.Process(input, func(i interface{}) interface{} {
		item := i.(Payload)
		time.Sleep(1 * time.Second)

		item.Number = item.Number * 2

		return item
	})

	fmt.Printf("%v\n", result)
}

func TestExampleBasic(t *testing.T) {
	Example_basic()
}
