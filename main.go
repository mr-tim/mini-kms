package main

import (
	"./api"
	"./kms"
)

func main() {
	k := kms.New()
	a := api.New(k)
	a.Run()
}
