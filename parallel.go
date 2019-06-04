package parallel

// Input is used to pass multiple payloads to the processor.
type Input []interface{}

// Output contains all results produced by the callback.
type Output []interface{}

// Callback defines the sigiture the callback function needs to follow.
type Callback func(interface{}) interface{}

// Processor does the work.
type Processor struct {
	Threads int
}

// Process takes the input and callback and returns the output. It will
// start a defined number of threads that will start processing the input
// untill all are done.
func (ap *Processor) Process(input Input, callback Callback) Output {
	var out Output

	jobs := make(chan interface{}, len(input))
	results := make(chan interface{}, len(input))

	for t := 1; t <= ap.Threads; t++ {
		go func(id int, jobs chan interface{}, out interface{}) {
			for job := range jobs {
				results <- callback(job)
			}
		}(t, jobs, results)
	}

	for _, i := range input {
		jobs <- i
	}
	close(jobs)

	for a := 1; a <= len(input); a++ {
		out = append(out, <-results)
	}

	return out
}
