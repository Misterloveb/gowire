package main

import "gindemo/cmd/server/wire"

func main() {
	wire.App().Start()
}
