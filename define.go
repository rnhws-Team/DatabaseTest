package main

import (
	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

type DataBase struct {
	DB           *bbolt.DB
	DataReceived map[uuid.UUID]uuid.UUID
}
