package unexported

type Counter struct {
	n int
}

func (c *Counter) Count() {
	c.n++
}

func (c *Counter) reset() {
	c.n = 0
}

func (c *Counter) Get() int {
	return c.n
}
