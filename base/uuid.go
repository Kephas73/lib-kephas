package base

import (
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v3"
)

func NewUuid() string {
	var uuid_ uuid.UUID
	uuid_, _ = uuid.NewRandom()

	return uuid_.String()
}

func ShortUUIDFromString(uuidStr string) (string, error) {
	uuid_, err := uuid.Parse(uuidStr)
	if err != nil {
		return "", err
	}
	shortUUID := shortuuid.DefaultEncoder.Encode(uuid_)
	return shortUUID, nil
}

func UUIDFromShortUUID(shortUUID string) (string, error) {
	uuid_, err := shortuuid.DefaultEncoder.Decode(shortUUID)
	if err != nil {
		return "", err
	}
	uuidStr := uuid_.String()
	return uuidStr, err
}
