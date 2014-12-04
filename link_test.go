package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateLink(t *testing.T) {
  destination := "http://app.com"

  link := CreateLink(destination)
  assert.Nil(t, link)


  assert.Empty(t, nil)
}
