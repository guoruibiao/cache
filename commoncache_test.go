package cache

import (
	"testing"
	"fmt"
)

func TestCommonCache_Get(t *testing.T) {
	cacher := NewCommonCacher()
	cacher.Cache("name", "guoruibiao")
	cacher.Cache("age", 100)
	cacher.Cache("address", "北京市朝阳区")

	fmt.Println(cacher.GetCachedData())
}

func TestCommonCache_Delete(t *testing.T) {
	cacher := NewCommonCacher()
	cacher.Cache("name", "guoruibiao")
	cacher.Cache("age", 25)
	cacher.Cache("address", "北京市朝阳区")
	t.Log(cacher.GetCachedData())
	cacher.Delete("name")
	t.Log(cacher.GetCachedData())
}