package child

var handlers = map[string]func(*Request) []byte{}

//AddHandler associates a function with your request object
func AddHandler(route string, f func(*Request) []byte) {
	handlers[route] = f
}
