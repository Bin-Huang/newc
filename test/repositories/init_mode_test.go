package repositories

import "testing"

func TestInitMode(t *testing.T) {
	r := NewUserRepository(nil, nil, "")
	if r.TableName != "foo" {
		t.Error("generated init-mode code invalid: the constructor of UserRepository should calls `s.init()` method")
	}
}