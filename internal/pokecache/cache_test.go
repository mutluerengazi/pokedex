package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{"https://example.com", []byte("testdata")},
		{"https://example.com/path", []byte("moretestdata")},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)

			got, ok := cache.Get(c.key)
			if !ok {
				t.Fatalf("expected key %q to exist", c.key)
			}
			if string(got) != string(c.val) {
				t.Fatalf("expected %q, got %q", c.val, got)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const base = 5 * time.Millisecond
	cache := NewCache(base)

	cache.Add("k", []byte("v"))
	if _, ok := cache.Get("k"); !ok {
		t.Fatalf("key missing right after Add")
	}

	time.Sleep(base + 5*time.Millisecond)

	if _, ok := cache.Get("k"); ok {
		t.Fatalf("key should have been reaped")
	}
}
