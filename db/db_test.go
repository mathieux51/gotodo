package db

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	rh := "redis-host"
	os.Setenv("REDIS_HOST", rh)
	got := getEnv("REDIS_HOST", "127.0.0.1")
	want := rh
	if got != want {
		t.Errorf("getEnv(\"REDIS_HOST\", \"127.0.0.1\") = %v, got %v", want, got)
	}
}

func TestGetEnv_fallback(t *testing.T) {
	// Init
	os.Unsetenv("REDIS_HOST")

	got := getEnv("REDIS_HOST", "127.0.0.1")
	want := "127.0.0.1"
	if got != want {
		t.Errorf("getEnv(\"REDIS_HOST\", \"127.0.0.1\") = %v, got %v", want, got)
	}
}

// func TestGetStorage(t *testing.T) {
// 	// Init
// 	os.Unsetenv("REDIS_HOST")

// 	got := NewStorage()
// 	want := Storage{}
// 	if got != want {
// 		t.Errorf("getStorage() = %v, got %v", want, got)
// 	}
// }
