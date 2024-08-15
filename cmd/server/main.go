package main

import config "github.com/ragazoni/goexpert/apis/configs"

func main() {
	config, _ := config.LoadConfig(".")
	println(config.DBDriiver)
}
