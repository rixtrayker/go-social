package main

import (
	"github.com/joho/godotenv"
	S "github.com/rixtrayker/go-social/server"
)

func init() {
	godotenv.Load()
}
func main() {
	S.Server()
}
