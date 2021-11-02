pprof

https://github.com/gin-contrib/pprof

Run Tests codecov Go Report Card GoDoc Join the chat at https://gitter.im/gin-gonic/gin

gin pprof middleware

Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.

Usage
Start using it
Download and install it:

go get github.com/gin-contrib/pprof
Import it in your code:

import "github.com/gin-contrib/pprof"
Example
package main

import (
"github.com/gin-contrib/pprof"
"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()
pprof.Register(router)
router.Run(":8080")
}
change default path prefix
func main() {
router := gin.Default()
// default is "debug/pprof"
pprof.Register(router, "dev/pprof")
router.Run(":8080")
}
custom router group
package main

import (
"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
router := gin.Default()
pprof.Register(router)
adminGroup := router.Group("/admin", func(c *gin.Context) {
if c.Request.Header.Get("Authorization") != "foobar" {
c.AbortWithStatus(http.StatusForbidden)
return
}
c.Next()
})
pprof.RouteRegister(adminGroup, "pprof")
router.Run(":8080")
}
Use the pprof tool
Then use the pprof tool to look at the heap profile:

go tool pprof http://localhost:8080/debug/pprof/heap
Or to look at a 30-second CPU profile:

go tool pprof http://localhost:8080/debug/pprof/profile
Or to look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:

go tool pprof http://localhost:8080/debug/pprof/block
Or to collect a 5-second execution trace:

wget http://localhost:8080/debug/pprof/trace?seconds=5