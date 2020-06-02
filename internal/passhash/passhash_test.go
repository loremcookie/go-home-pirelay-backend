package passhash

import "testing"

//TestHashString is the test_plugin function of HashString
func TestHashString(t *testing.T) {
	var err error

	pass := "password"

	hash, err := HashString(pass)
	if err != nil {
		t.Error(err)
	}

	if !MatchString(pass, hash) {
		t.Error("String and hash does not match")
	}
}
