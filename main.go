package main

import (
	"github.com/ivo100/multimod/cache"
	"log"
	"time"
)
//

// type Svc struct {
// 	userservice.UserService
// }

func main() {
	cycle := 100 * time.Millisecond
	c := cache.New(cycle)
	defer c.Close()

	c.Set("sticky", "forever", 0)
	c.Set("hello", "Hello", cycle/2)
	hello, found := c.Get("hello")
	log.Printf("found %v, hello %v", found, hello)
	// var svc Svc
	// _ = svc

}
