package rescue

import (
	"log"

	"github.com/pkg/errors"
)

func RunWithRecover(f func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%+v", errors.Errorf("Unexpected panic in util.RunWithRecover(): %+v", err))
		}
	}()

	f()
}
