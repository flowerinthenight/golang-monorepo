package util

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// Err2 returns an error with stacktrace when glog level is 2, otherwise,
// return the same error with the optional wrap string, if provided.
func Err2(err error, s ...string) error {
	if glog.V(2) {
		if len(s) > 0 {
			// add wrap text with stacktrace
			return errors.Wrap(err, s[0])
		}

		// add stacktrace
		return errors.WithStack(err)
	}

	if len(s) > 0 {
		return fmt.Errorf("%s: %v", s[0], err)
	}

	return err
}
