package configs

import (
  "os"
)

var LYRICS_FETCH_LIMITS = os.Getenv("LYRICS_FETCH_LIMITS")
var API_SERVER_PORT = os.Getenv("API_SERVER_PORT")
