package main

import "github.com/plankiton/joaoantonio/pkg/rest"

func main() {
	handler := rest.New()
	handler.Logger.Fatal(handler.Start(":1323"))
}
