package main

import (
	"fmt"
	"machine"
	"math/rand"
	"time"

	"github.com/stevegt/com16566"
)

func main() {

	i2c := machine.I2C0
	err := i2c.Configure(machine.I2CConfig{})
	if err != nil {
		fmt.Println("could not configure I2C:", err)
		return
	}

	relay := &com16566.COM16566{Addr: com16566.I2C_ADDRESS_DEFAULT, I2c: i2c}

	var bit bool

	for r := 1; r <= 4; r++ {
		err = relay.Off(r)
		fmt.Println("off", r, err)
		bit, err = relay.Status(r)
		fmt.Println("status", r, bit, err)

		bit, err = relay.Toggle(r)
		fmt.Println("toggle", r, bit, err)
		bit, err = relay.Toggle(r)
		fmt.Println("toggle", r, bit, err)

		err = relay.On(r)
		fmt.Println("on", r, err)
		bit, err = relay.Status(r)
		fmt.Println("status", r, bit, err)

		err = relay.Off(r)
		fmt.Println("off", r, err)
		bit, err = relay.Status(r)
		fmt.Println("status", r, bit, err)

		fmt.Println()
		time.Sleep(1 * time.Second)
	}

	for r := 1; r <= 4; r++ {
		go play(relay, r)
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func play(relay *com16566.COM16566, r int) {
	for {
		bit, err := relay.Toggle(r)
		fmt.Println("toggle", r, bit, err)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
