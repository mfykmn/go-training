package unexported // テスト対象と同じパッケージ

func SetBaseURL(s string) (resetFunc func()) {
	var tmp string
	tmp, baseURL = baseURL, s
	return func() {
		baseURL = tmp
	}
}

type ExportGetResponse = getResponse
