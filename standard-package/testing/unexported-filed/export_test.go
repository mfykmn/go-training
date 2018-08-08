package unexported // テスト対象と同じパッケージ

func (c *Counter) ExportSetN(n int) {
	c.n = n
}