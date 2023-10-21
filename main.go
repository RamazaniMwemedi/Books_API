package main

import "main/GO_RESTAPI/controllers/login"

func main() {
	user := login.Login()
	print("Name ", user)
}
