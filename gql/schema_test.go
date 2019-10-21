package gql

import "testing"

func TestGetSchema(t *testing.T) {
	_, err := GetSchema()
	if err != nil {
		t.Error(err)
	}
}
