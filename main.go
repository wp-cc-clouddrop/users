package main

import (
	"users/internal/api"
)

func main() {
	// http.Handle("/metrics", promhttp.Handler())
	// go panic(http.ListenAndServe(":8081", nil))

	api.Init(8080)
}
