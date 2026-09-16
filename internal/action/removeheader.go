package action

import "github.com/google/uuid"

type RemoveHeader struct {
	Id uuid.UUID
}

func NewRemoveHeader(id uuid.UUID) RemoveHeader {
	return RemoveHeader{
		Id: id,
	}
}
