package hashmap

import (
	"testing"
)

func Test_hashmap_Set(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value string
	}{
		{"empty", "", ""},
		{"empty key", "", "value"},
		{"empty value", "key", ""},
		{"test", "key", "value"},
		{"collission", "key1", "value1"},
		{"collission", "key2", "value2"},
	}

	hm := NewHashmap(10)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm.Set(tt.key, tt.value)
			contains := false
			for _, bucket := range hm.storage {
				for _, item := range bucket {
					if item.value == tt.value {
						contains = true
					}
				}
			}
			if !contains {
				t.Errorf("Can't find %s in the hashmap", tt.value)
			}
		})
	}
}

func Test_hashmap_Get(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		val     string
		wantVal string
		wantOk  bool
	}{
		{"empty", "", "", "", true},
		{"empty key", "", "value", "value", true},
		{"empty value", "key", "", "", true},
		{"test 0", "key", "value", "value", true},
		{"test 1", "key1", "value1", "value1", true},
		{"test 2", "key2", "value2", "value2", true},
		{"test 3", "key3", "value3", "value3", true},
		{"test 4", "key4", "value4", "value4", true},
		{"test 5", "key5", "value5", "value5", true},
		{"test 6", "key6", "value6", "value6", true},
		{"test 7", "key7", "value7", "value7", true},
		{"test 8", "key8", "value8", "value8", true},
		{"test 9", "key9", "value9", "value9", true},
		{"test 10", "key10", "value10", "value10", true},
		{"test 11", "hello", "world", "world", true},
		{"test 12", "foo", "bar", "bar", true},
		{"test 13", "abc123", "def456", "def456", true},
		{"test 14", "testkey", "testvalue", "testvalue", true},
		{"test 15", "longkeyvalue", "thisisalongvalue", "thisisalongvalue", true},
	}

	hm := NewHashmap(10)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm.Set(tt.key, tt.val)
			got, ok := hm.Get(tt.key)
			if got != tt.wantVal || ok != tt.wantOk {
				t.Errorf("want: val=%s ok=%t, got: val=%s ok=%t", tt.wantVal, tt.wantOk, got, ok)
			}
		})
	}
}

func Test_hashmap_Get_NoMatch(t *testing.T) {
	hm := NewHashmap(4)
	hm.Set("qwerty", "aaa")
	hm.Set("", "123")
	wantVal := ""
	wantOk := false
	got, ok := hm.Get("abs")
	if got != wantVal || ok != wantOk {
		t.Errorf("want: val=%s ok=%t, got: val=%s ok=%t", wantVal, wantOk, got, ok)
	}
}

func Test_hashmap_Delete(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		wantVal string
		wantOk  bool
	}{
		{"existing key", "qwerty", "", false},
		{"non-existing key", "qqqq", "", false},
		{"empty key", "", "", false},
	}

	hm := NewHashmap(8)
	hm.Set("qwerty", "aaa")
	hm.Set("", "123")
	hm.Set("666", "bbbb")
	hm.Set("ururu", "a-a-a-a")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm.Delete(tt.key)
			got, ok := hm.Get(tt.key)
			if got != tt.wantVal || ok != tt.wantOk {
				t.Errorf("want: val=%s ok=%t, got: val=%s ok=%t", tt.wantVal, tt.wantOk, got, ok)
			}
		})
	}
}
