package main

import (
	"trade-union-service/internal/app"
)


func main() {
	a := app.NewApp()
	if err := a.Run(); err != nil {
		panic(err)
	}
}