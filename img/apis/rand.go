package main

import(
	"math/rand"
)


func randStr(animals []string) string {
	return animals[rand.Intn(len(animals))]
}