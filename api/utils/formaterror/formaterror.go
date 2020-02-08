package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	if strings.Contains(err, "titulo") {
		return errors.New("Titulo já cadastrado")
	}
	if strings.Contains(err, "isbn") {
		return errors.New("Isbn já cadastrado")
	}
	return errors.New("Incorrect Details")
}
