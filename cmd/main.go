package main

import (
	"fmt"
	"github.com/fianprabowo/TRAFFIC-LIGHT/internal/machine"
)

func main() {
	fmt.Println("Welcome to mini traffic light simulation🚦")
	tlMachine := machine.NewMachine()
	for {
		fmt.Println("Current Light Color:", tlMachine.GetColor())
		if err := tlMachine.Change(); err != nil {
			fmt.Println("Error changing state:", err)
			break
		}
	}

}
