package main

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"
)

type KeyVal struct {
	key   string
	value []byte
}

func slicesEqual(slice1, slice2 []byte) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}

	return true
}

func testAdd(t *testing.T) {
	str := [3]string{"Fly me to the moon", "My way", "Frank Sinatra"}
	byteSlice1 := []byte(str[0])
	byteSlice2 := []byte(str[1])
	byteSlice3 := []byte(str[2])
	cases := []struct {
		input  KeyVal
		output Cache
	}{
		{
			input: KeyVal{
				key:   "myKey",
				value: byteSlice1,
			},
			output: Cache{
				cache: map[string]cacheEntry{
					"myKey": {
						createdAt: time.Now(),
						val:       byteSlice1,
					},
				},
				mu: &sync.RWMutex{},
			},
		},
		{
			input: KeyVal{
				key:   "yourKey",
				value: byteSlice3,
			},
			output: Cache{
				cache: map[string]cacheEntry{
					"yourKey": {
						createdAt: time.Now(),
						val:       byteSlice1,
					},
				},
				mu: &sync.RWMutex{},
			},
		},
		{
			input: KeyVal{
				key:   "franksKey",
				value: byteSlice2,
			},
			output: Cache{
				cache: map[string]cacheEntry{
					"franksKey": {
						createdAt: time.Now(),
						val:       byteSlice1,
					},
				},
				mu: &sync.RWMutex{},
			},
		},
	}
	for _, c := range cases {
		result := Cache{
			cache: make(map[string]cacheEntry),
			mu:    &sync.RWMutex{},
		}
		result.Add(c.input.key, c.input.value)
		llave := c.input.key
		slice1 := c.output.cache[llave].val
		slice2 := result.cache[llave].val

		if !slicesEqual(slice1, slice2) {
			t.Errorf("no se agrega correctamente")
		}

	}
}

func testGet(t *testing.T) {
	strArray := [3]string{"Fly me to the moon", "My way", "Frank Sinatra"}
	byteSlice1 := []byte(strArray[0])
	byteSlice3 := []byte(strArray[2])
	cases := []struct {
		input  KeyVal
		output Cache
	}{
		{
			input: KeyVal{
				key:   "myKey",
				value: byteSlice1,
			},
			output: Cache{
				cache: map[string]cacheEntry{
					"myKey": {
						createdAt: time.Now(),
						val:       byteSlice1,
					},
				},
				mu: &sync.RWMutex{},
			},
		},
		{
			input: KeyVal{
				key:   "tumami en bici",
				value: byteSlice3,
			},
			output: Cache{
				cache: map[string]cacheEntry{
					"franksKey": {
						createdAt: time.Now(),
						val:       byteSlice1,
					},
				},
				mu: &sync.RWMutex{},
			},
		},
	}

	result, exists := cases[0].output.Get(cases[0].input.key)
	if !exists {
		t.Errorf("Expected key %s to exist", cases[0].input.key)
	}
	if !bytes.Equal(result, cases[0].input.value) {
		t.Errorf("get no esta funcionando")
	}

	result1, exists1 := cases[1].output.Get(cases[1].input.key)
	if !exists1 {
		t.Errorf("Expected key %s not to exist", cases[1].input.key)
	}
	if result1 != nil {
		t.Errorf("Expected nil result for non-existing key")
	}

}
func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 10*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	a, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to not find key: %v", a)
		return
	}
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}
