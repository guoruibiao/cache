package cache

import "testing"

func TestLRUCache_Get(t *testing.T) {
	lrucacher, err := NewLRUCache(3)
	if err != nil {
		t.Error(err)
	}
	lrucacher.Cache("name", "guoruibiao")
	lrucacher.Cache("age", 25)
	lrucacher.Cache("address", "北京市朝阳区")
	//t.Log(lrucacher.GetCachedData())
	t.Log(lrucacher.Get("name"))
	lrucacher.Cache("school", "大连理工大学")
	t.Log(lrucacher.Get("name"))
	t.Log(lrucacher.Get("age"))
	t.Log(lrucacher.Get("school"))
    //t.Log(lrucacher.GetCachedData())

}


func TestLRUCache_Delete(t *testing.T) {
	lrucacher, err := NewLRUCache(3)
	if err != nil {
		t.Error(err)
	}
	lrucacher.Cache("name", "guoruibiao")
	lrucacher.Cache("age", 25)
	lrucacher.Cache("address", "北京市朝阳区")
	t.Log(lrucacher.Get("name"))
	t.Log(lrucacher.Peek())
	// test for deleting element
	lrucacher.Delete("name")
	t.Log(lrucacher.Size())
	t.Log(lrucacher.Peek())
}