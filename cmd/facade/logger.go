package main

import (
  "log"
  "os"
)

var logger *log.Logger = log.New(os.Stderr, "", log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
