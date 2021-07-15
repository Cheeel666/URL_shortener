package main

import "URL_shortener/server"


func main() {
	server.SetupServer().Run("0.0.0.0:8070")
}
