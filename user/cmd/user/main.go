package main

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/user/config"
)

func main() {
	cfg := config.InitConfig()

	fmt.Printf("%+v\n", cfg)
}
