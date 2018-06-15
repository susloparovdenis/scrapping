package scrapping

type Result struct {
	title   string
	url     string
	sources []string
}

func resultToSlice(c <-chan Result) []Result {
	s := make([]Result, 0)
	for i := range c {
		s = append(s, i)
	}
	return s
}
