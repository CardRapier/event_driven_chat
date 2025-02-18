package main

import (
	"chat/cmd/api"
	"chat/cmd/persistence"
)

func main() {
	go persistence.StartPersistenceService()
	api.StartApi()

}