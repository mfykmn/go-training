package unexported

type Counter struct {
	n int
}

func (c *Counter) ExportN() int {
	return c.n
}
