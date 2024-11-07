package rand

import (
	"fmt"
	"testing"
)

func TestClientSecret(t *testing.T) {
	secret, err := ClientSecret()
	if err != nil {
		t.Errorf("Error generating secret: %v", err)
	}
	if len(secret) != 72 {
		t.Errorf("Expected length of 128, got %v", len(secret))
	}
	fmt.Println("ClientSecret:", secret)
}

func TestUidV4(t *testing.T) {
	Uid := UidV4()

	if len(Uid) != 36 {
		t.Errorf("Expected length of 36, got %v", len(Uid))
	}
	fmt.Println("UidV4:", Uid)
}

func TestStringFixed(t *testing.T) {
	str := StringFixed(32)
	if len(str) != 32 {
		t.Errorf("Expected length of 10, got %v", len(str))
	}
	fmt.Println("StringFixed:", str)
}
