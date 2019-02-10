package main

import "github.com/stocktables/stocktables"

func main() {
	r := stocktables.Router()
	stocktables.Serve("127.0.0.1:0", r)
}
