package internal

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapError(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		err error
	}{
		"error": {err: fmt.Errorf("[err] obj error")},
	}

	for _, tc := range tests {
		err := WrapError(tc.err)
		assert.True(errors.Is(err, tc.err))
	}
}
