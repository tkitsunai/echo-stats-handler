[golang-stats-api-handler](https://github.com/fukata/golang-stats-api-handler) available for framework "echo"

## Install

```
go get github.com/tkitsunai/echo-stats-handler
```

or,

If you use "dep" for package manager.

```
dep ensure -add github.com/tkitsunai/echo-stats-handler
```

If you use "glide" for package manager.

```
glide get github.com/tkitsunai/echo-stats-handler
```


## Usage Example
```
import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tkitsunai/echo-stats-handler"
)

func main() {
	e := echo.New()
	e.GET("/stats", echo-stats-handler.Stats)
	e.Start(":8000")
}
```

## Response

```
{
  time: 1522261688402036500,
  go_version: "go1.10",
  go_os: "darwin",
  go_arch: "amd64",
  cpu_num: 4,
  goroutine_num: 3,
  gomaxprocs: 4,
  cgo_call_num: 1,
  memory_alloc: 409680,
  memory_total_alloc: 409680,
  memory_sys: 4262136,
  memory_lookups: 5,
  memory_mallocs: 4957,
  memory_frees: 101,
  memory_stack: 360448,
  heap_alloc: 409680,
  heap_sys: 1736704,
  heap_idle: 761856,
  heap_inuse: 974848,
  heap_released: 0,
  heap_objects: 4856,
  gc_next: 4473924,
  gc_last: 0,
  gc_num: 0,
  gc_per_second: 0,
  gc_pause_per_second: 0,
  gc_pause: [ ]
}
```

## using in framework "echo"

- [echo](https://github.com/labstack/echo)
