package episode0

import "testing"

func TestBusinessLogic(t *testing.T) {
	ht := NewMemHashTable()
	BusinessLogic(ht)

	val, err := ht.Get("hello")
	if err != nil {
		t.Fatalf("error on GET %s", err)
	}
	if string(val) != "world" {
		t.Fatalf("expected 'world', got '%s'", string(val))
	}
}
