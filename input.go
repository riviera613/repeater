package repeater

type InputParam struct {
	Concurrence int64
	TotalCount  int64
}

type InputFunc struct {
	Name string
	Func func() error
}
