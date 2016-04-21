package request

type request struct {
	Method string
	RawUrl string
	Header map[string][]string
}
