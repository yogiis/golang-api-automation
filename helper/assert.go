package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	testing testing.T
}

func (t *Case) assertNew() *assert.Assertions {
	return assert.New(&t.testing)
}

func (t *Case) AssertEqual(expected, actual, err interface{}) bool {
	equal := t.assertNew().Equal(fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual), err)
	if !equal {
		LogPanicln(ErrorHandleEqual(expected, actual))
	}
	return true
}

func (t *Case) AssertTrue(expected bool, err interface{}) bool {
	isTrue := t.assertNew().True(expected, err)

	if !isTrue {
		LogPanicln(ErrorHandleBool(isTrue))
	}

	return true
}
