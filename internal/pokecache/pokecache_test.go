package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
    cache := NewCache(time.Millisecond)
    if cache.cache == nil{
        t.Error("cache is nil")
    }
}

func TestAddGetCache(t *testing.T) {
    cache := NewCache(time.Millisecond)

    cases := []struct {
        inputKey string
        inputVal []byte
    } {
        {
            inputKey: "key1",
            inputVal: []byte("val1"),
        },
        {
            inputKey: "key2",
            inputVal: []byte("val2"),
        },
        {
            inputKey: "",
            inputVal: []byte("val3"),
        },
    }

    for _, cas := range cases {
        cache.Add(cas.inputKey, cas.inputVal)
        actual, ok := cache.Get(cas.inputKey)

        if !ok {
            t.Errorf("%s not found\n", cas.inputKey)
            continue
        }

        val := string(cas.inputVal)
        act := string(actual) 
        if act != val {
            t.Errorf("%s doesn't match %s", act, val)
            
        }
    }
}

func TestReapCache(t *testing.T) {
    interval := 10*time.Millisecond
    cache := NewCache(interval)
    cache.Add("key1", []byte("val1"))

    time.Sleep(interval + time.Millisecond)

    _, ok := cache.Get("key1")
    if ok {
        t.Error("key1 should have been reaped")
    }
}
