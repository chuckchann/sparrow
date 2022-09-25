package util

import "github.com/pborman/uuid"

func GenUUID() string {
	return uuid.New()
}
