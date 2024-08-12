package redis

import "github.com/google/uuid"

func Lock() (uuid.UUID, string) {
	id := uuid.New()
	cmd := "SETNX " + id.String() + " EXPIRE 100000"
	return id, cmd
}
