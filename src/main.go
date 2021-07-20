package main

import (
  "github.com/NUTMEG-hackathon/gin-sam-1/src/seed"
  "github.com/NUTMEG-hackathon/gin-sam-1/src/router"
)

func main() {
  seed.Seed()
  router.Router()
}
