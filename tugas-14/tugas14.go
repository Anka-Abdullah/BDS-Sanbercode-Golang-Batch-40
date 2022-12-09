package main

import "BDS-Sanbercode-Golang-Batch-40/tugas-14/routers"

var PORT = ":8080"

func main() {
	routers.StartServer().Run(PORT)
}
