package testable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	fake := NewFake()

	assert.Nil(t, fake.Create(&User{"colin", "shenzhen"}))
}
