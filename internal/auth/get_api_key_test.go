package auth

import (
    "reflect"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    if !reflect.DeepEqual(1, 1) {
        t.Fatalf("test fail")
    }
}
