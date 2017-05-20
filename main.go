package main

import "golang.org/x/exp/io/i2c"

func main() {
	_, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-0"}, 0x48)
	if err != nil {
    		panic(err)
	}
}
