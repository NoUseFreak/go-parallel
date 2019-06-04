package parallel

type Input []interface{}
type Output []interface{}
type Callback func(interface{}) interface{}
type Processor struct {
	Threads int
}

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
