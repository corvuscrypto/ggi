package child

func handleRequest(r *Request) *responsePack {

	if v, ok := handlers[r.Path]; ok {
		return &responsePack{
			pid,
			200,
			v(r),
		}
	}
	return &responsePack{
		pid,
		404,
		nil,
	}
}
