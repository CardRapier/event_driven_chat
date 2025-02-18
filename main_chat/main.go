package main

import (
	"chat/cmd/api"
	"chat/cmd/persistence"
)

func main() {
	go api.StartApi()

	persistence.StartPersistenceService()
}