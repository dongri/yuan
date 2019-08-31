package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := []string{"🍎", "🍇", "🍈", "🍉"}
	i := 0
	for i < 100000000 {
		s := a[rand.Intn(4)]
		fmt.Print(s)
		i = i + 1
	}
}
