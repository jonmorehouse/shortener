package main

import (
  "os"
  "log"
  "strconv"
)

var linkChars string
var linkCharsCount int
var linkLength int

type Link struct {
  link string 
  destination string
  live bool
}

func linkInit() {
  // grab link characters
  if linkChars == "" || linkCharsCount == 0 {
    linkChars = os.Getenv("LINK_CHARS")
    linkCharsCount = len(linkChars)

    if linkCharsCount == 0 {
      log.Fatalln("$LINK_CHARS is not set.")
    }
  }

  // grab link length
  if linkLength == 0 {
    length := os.Getenv("LINK_LENGTH")
    if len(length) == 0 {
      log.Fatalln("$LINK_LENGTH is not set.")
    } else {
      intLength, err := strconv.Atoi(length)
      if err != nil {
        log.Fatalln("$LINK_LENGTH is not an integer.")
      } else {
        linkLength = intLength
      }
    }
  }//link length set
}

// create a short link
func CreateLink(destination string) *Link {


  return nil
}

func UseLink(destination string) *Link {

  return nil
}



