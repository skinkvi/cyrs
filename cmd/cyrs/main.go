package main

import "github.com/skinkvi/cyrs/internal/config"

func main() {
	cfg := config.MustLoadBuPath("../../config/local.yaml")

	// TODO init application

	// TODO Graceful shotdown
}
