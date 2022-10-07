package service

import (
	"github.com/ValikoDorodnov/go_sample/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyMockedObject struct {
	repo *repository.GreetingRepository
}

func TestSomething(t *testing.T) {
	testObj := new(MyMockedObject)
	service := NewGreetingService(testObj.repo)
	actual, _ := service.HelloProcess()
	assert.Equal(t, "Hello", actual, "they should be equal")
}
