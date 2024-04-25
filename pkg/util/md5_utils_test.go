package util

import "testing"

func TestMD5SaltHash(t *testing.T) {
	hash := MD5SaltHash("text", "salt")
	expected := "5ff85b3c02a14a1cc24aaa14d43654cd"
	if hash != expected {
		t.Errorf("MD5SaltHash was incorrect, got: %s, want: %s.", hash, expected)
	}
}

func TestMD5Hash(t *testing.T) {
	hash := MD5Hash("123456")
	expected := "e10adc3949ba59abbe56e057f20f883e"
	if hash != expected {
		t.Errorf("MD5Hash was incorrect, got: %s, want: %s.", hash, expected)
	}
}
