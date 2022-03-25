package utils

import (
	"context"
	"errors"
	"log"
)

func PanicRecover(ctx context.Context, p interface{}) error {
	log.Default().Printf("panic: %v %v", ctx, p)
	return errors.New("internal server error")
}
