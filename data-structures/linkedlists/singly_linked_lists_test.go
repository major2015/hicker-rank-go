package singlylinkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglylinkedList_GetFirstValue(t *testing.T) {
	list := New()

	_, err := list.GetFirstValue()
	assert.NotNil(t, err)

	list.Add("abc")
	value, err := list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")

	list.Add("def")
	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")

	list.InsertAt(0, "hij")

	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "hij")
}