package main

import "URL_shortener/server"

// запускаем мейн и радуемся жизни
func main() {
	server.SetupServer().Run()
}
