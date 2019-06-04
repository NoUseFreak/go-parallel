/*
Package go-parallel aims to simplify processing data in parallel by providing a small helper.

  package main

  import (
    "fmt"
    "time"

    "github.com/NoUseFreak/go-parallel"
  )

  type Payload struct {
    Number int
  }

  func main() {
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


*/
package parallel
