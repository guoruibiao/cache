# cache
cache in golang.

## how to use

```
go get -u github.com/guoruibiao/cache
```

```go
package main

import (
	"github.com/guoruibiao/cache"
	"fmt"
	"os"
)

func CommonCacheUsage() {
	cacher := cache.NewCommonCacher()
	cacher.Cache("name", "guoruibiao")
	cacher.Cache("age", 25)
	cacher.Cache("address", "北京市朝阳区")
	cacher.Cache("school", "大连理工大学")
	fmt.Println(cacher.Get("name"))
	cacher.Delete("name")
	fmt.Println(cacher.Get("name"))
}

func LRUCacheUsage() {
	/*
	_, err := cache.NewLRUCache(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	*/
	cacher, err := cache.NewLRUCache(3)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	cacher.Cache("name", "guoruibiao")
	cacher.Cache("age", 25)
	cacher.Cache("address", "北京市朝阳区")
	cacher.Cache("school", "大连理工大学")
	fmt.Println(cacher.Get("name"))
	fmt.Println("Peek data is: ")
	fmt.Println(cacher.Peek())
	fmt.Println(cacher.Get("address"))
	fmt.Println("Peek data is: ")
	fmt.Println(cacher.Peek())

}

func main() {
	CommonCacheUsage()
	fmt.Println("-------------------")
	LRUCacheUsage()
}
```

Output:

```
guoruibiao <nil>
<nil> no such data for key: name
-------------------
<nil> <nil>
Peek data is:
大连理工大学 <nil>
北京市朝阳区 <nil>
Peek data is:
北京市朝阳区 <nil>

```
