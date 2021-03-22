package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapPerson(t *testing.T) {
	s, _ := DummyHTTP()
	person, err := MapPerson(s.URL)
	assert.Equal(t, "Cagatay Cali", person[0].Name)
	assert.Equal(t, nil, err)
}
func TestMapPersonViaInvalidURI(t *testing.T) {
	_, err := MapPerson("http://invalid.com")
	assert.NotEqual(t, nil, err)
}
