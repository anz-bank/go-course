package puppy

import "fmt"

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code %d: %s",
		e.Code, e.Message)
}

const (
	ErrorValueFormat = 1000
	NegativeValue    = 1001
	NonExistentPuppy = 1002

	PuppyValueLessThanZero = "Puppy value can't be less than 0."
	PuppyValueNotFound     = "Puppy ID can not be found, read operation failed."
	InvalidPuppyValue      = "Unrecongised puppy value."
	InvalidIDForRead       = "Puppy ID can not be found, read operation failed."
	InvliadIDForUpdate     = "Puppy ID can not be found, update operation failed."
	InvalidIDForDelete     = "Puppy ID can not be found, delete operation failed."

	InvalidPort = "Invalid port number entered, default port 8888 will be used"
)
