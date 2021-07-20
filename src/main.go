package main

import (
  "github.com/mizu-ryo/gin-sam2/src/seed"
  "github.com/mizu-ryo/gin-sam2/src/router"
)

func main() {
  seed.Seed()
  router.Router()
}
