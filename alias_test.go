package logging

import (
	"errors"
	"testing"
	"time"
)

func TestFloat32Attr(t *testing.T) {
	key := "testKey"
	val := float32(123.456)

	attr := Float32Attr(key, val)

	// Checking if the attribute correctly converts float32 to float64
	if attr.Key != key || attr.Value.Float64() != float64(val) {
		t.Errorf("Float32Attr() = %v, want %v", attr, Float64Attr(key, float64(val)))
	}
}

func TestStringAttr(t *testing.T) {
	key := "testKey"
	val := "testValue"
	attr := StringAttr(key, val)
	if attr.Key != key || attr.Value.String() != val {
		t.Errorf("key = %v, want %v", attr, StringAttr(key, val))
	}
}

func TestErrAttr(t *testing.T) {
	err := errors.New("test error")
	attr := ErrAttr(err)
	if attr.Key != "error" || attr.Value.String() != err.Error() {
		t.Errorf("ErrAttr() = %v, want %v", attr, StringAttr("error", err.Error()))
	}
}

func TestUInt32Attr(t *testing.T) {
	key := "testKey"
	val := uint32(123456789)

	attr := UInt32Attr(key, val)

	// The expected behavior is that UInt32Attr converts the uint32 to an int.
	// We need to verify that this conversion is correct.
	expectedValue := int(val)
	if attr.Key != key || int(attr.Value.Int64()) != expectedValue {
		t.Errorf("UInt32Attr() = {Key: %s, Value: %v}, want {Key: %s, Value: %v}",
			attr.Key, attr.Value, key, expectedValue)
	}
}

func TestInt32Attr(t *testing.T) {
	key := "testKey"
	val := int32(12345)

	attr := Int32Attr(key, val)

	// Verify that the attribute has the correct key and value
	expectedValue := int(val)
	if attr.Key != key || int(attr.Value.Int64()) != expectedValue {
		t.Errorf("Int32Attr() = {Key: %s, Value: %v}, want {Key: %s, Value: %v}",
			attr.Key, attr.Value, key, expectedValue)
	}
}

func TestTimeAttr(t *testing.T) {
	key := "timestamp"
	val := time.Now()

	attr := TimeAttr(key, val)

	// Verify that the attribute has the correct key and the time in string format
	expectedValue := val.String()
	if attr.Key != key || attr.Value.String() != expectedValue {
		t.Errorf("TimeAttr() = {Key: %s, Value: %v}, want {Key: %s, Value: %v}",
			attr.Key, attr.Value, key, expectedValue)
	}
}
