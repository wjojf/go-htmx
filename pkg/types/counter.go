package types

var (
	DefaultCounter = NewCounter()
)

type Counter struct {
	Count int
}

func NewCounter() *Counter {
	return &Counter{
		Count: 0,
	}
}
