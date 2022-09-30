package networkingerrors

import (
	"log"

	"github.com/pkg/errors"
)

func Error(msg string) error {
	log.Print(msg)
	return errors.New(msg)
}
