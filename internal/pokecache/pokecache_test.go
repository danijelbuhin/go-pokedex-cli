package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	if cache.cache == nil {
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
			inputKey: "key5",
			inputVal: []byte("val2"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, cs.inputVal)

		actual, ok := cache.Get(cs.inputKey)

		if !ok {
			t.Errorf("key %s not found", cs.inputKey)
			continue
		}

		if string(actual) != string(cs.inputVal) {
			t.Errorf("%s doesn't match %s", string(actual), string(cs.inputVal))
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10

	cache := NewCache(interval)
	keyOne := "key1"
	cache.Add(keyOne, []byte("value"))

	time.Sleep(interval + time.Millisecond)
	
	_, ok := cache.Get(keyOne)

	if ok {
		t.Errorf("%s should have been deleted from the cache", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10

	cache := NewCache(interval)
	keyOne := "key1"
	cache.Add(keyOne, []byte("value"))

	time.Sleep(interval / 2)
	
	_, ok := cache.Get(keyOne)

	if !ok {
		t.Errorf("%s should not have been deleted from the cache", keyOne)
	}
}
