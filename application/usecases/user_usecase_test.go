package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSuccess(t *testing.T) {
	assert.EqualValues(t, 10+10, 20)
}
