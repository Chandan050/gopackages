package main

import f "fmt"

const boilingPoint = 100

func sayHello(name string) {
	f.Printf("Hello %s!\n", name)
}

func toFahrenheit(t float64) float64 {
	return t*1.8 + 32
}
