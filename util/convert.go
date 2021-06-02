package util

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/joaovicdsantos/discord-clone-api/exception"
)

func VerifyAndConvertID(id string) (uint, exception.HttpError) {
	idConvertido, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, exception.HttpError{
			Err:        errors.New(fmt.Sprintf("ID %s isn't valid", id)),
			StatusCode: 400,
		}
	}
	return uint(idConvertido), exception.HttpError{}
}
