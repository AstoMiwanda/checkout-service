package main

import (
	"checkout-service/app/cmd"

	// Load tzdata
	_ "time/tzdata"
)

func main() {
	cmd.Execute()
}
