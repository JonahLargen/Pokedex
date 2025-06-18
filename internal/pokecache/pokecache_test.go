package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(100 * time.Millisecond)
	key := "test-key"
	val := []byte("test-value")

	cache.Add(key, val)
	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected key to be present")
	}
	if !bytes.Equal(got, val) {
		t.Errorf("expected value %q, got %q", val, got)
	}
}

func TestCacheGetMissingKey(t *testing.T) {
	cache := NewCache(100 * time.Millisecond)
	_, ok := cache.Get("missing-key")
	if ok {
		t.Errorf("expected missing key to not be present")
	}
}

func TestCacheExpiration(t *testing.T) {
	cache := NewCache(50 * time.Millisecond)
	key := "expire-key"
	val := []byte("expire-value")

	cache.Add(key, val)
	time.Sleep(20 * time.Millisecond)
	if _, ok := cache.Get(key); !ok {
		t.Fatalf("expected key to be present before expiration")
	}

	time.Sleep(60 * time.Millisecond)
	_, ok := cache.Get(key)
	if ok {
		t.Errorf("expected key to be expired and removed from cache")
	}
}

func TestCacheMultipleEntries(t *testing.T) {
	cache := NewCache(100 * time.Millisecond)
	entries := map[string][]byte{
		"a": []byte("value-a"),
		"b": []byte("value-b"),
		"c": []byte("value-c"),
	}
	for k, v := range entries {
		cache.Add(k, v)
	}
	for k, expected := range entries {
		got, ok := cache.Get(k)
		if !ok {
			t.Errorf("expected key %q to be present", k)
		}
		if !bytes.Equal(got, expected) {
			t.Errorf("expected value %q, got %q for key %q", expected, got, k)
		}
	}
}

func TestCacheConcurrentAccess(t *testing.T) {
	cache := NewCache(100 * time.Millisecond)
	key := "concurrent"
	val := []byte("safe")

	done := make(chan bool)
	go func() {
		for i := 0; i < 100; i++ {
			cache.Add(key, val)
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 100; i++ {
			cache.Get(key)
		}
		done <- true
	}()
	<-done
	<-done
}
