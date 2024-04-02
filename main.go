package main

import (
	"ranking/router"
)

func main() {
	r := router.Router()
	r.Run(":9999")
}
