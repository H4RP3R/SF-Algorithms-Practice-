package main

import (
	"testing"
	"time"
)

func TestInMemoryCache(t *testing.T) {
	cache := NewInMemoryCache(5 * time.Second)
	cache.Set("key1", "value1")

	value := cache.Get("key1")
	if value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	time.Sleep(15 * time.Second)

	value = cache.Get("key1")
	if value != nil {
		t.Errorf("Expected nil, got %v", value)
	}

	quit <- true
}
