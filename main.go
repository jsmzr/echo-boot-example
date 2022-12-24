package main

import (
	"fmt"

	boot "github.com/jsmzr/boot-echo"
	_ "github.com/jsmzr/boot-plugin-apollo"
	_ "github.com/jsmzr/boot-plugin-logrus"
	_ "github.com/jsmzr/boot-plugin-prometheus"
	_ "github.com/jsmzr/boot-plugin-skywalking"
	_ "github.com/jsmzr/echo-boot-example/router"
)

func main() {
	if err := boot.Run(); err != nil {
		fmt.Println(err)
	}
}
