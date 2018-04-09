package btreestorage

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testdata struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func TestTreeGreenPath(t *testing.T) {
	keys := []string{"1", "2", "3", "4", "5", "6"}
	data := []testdata{
		testdata{"Boston", "USA"},
		testdata{"NYC", "USA"},
		testdata{"Seatle", "USA"},
		testdata{"Berlin", "Germany"},
		testdata{"Munich", "Germany"},
		testdata{"Paris", "France"},
	}

	tt := BTree{}

	for i := 0; i < 6; i++ {
		tt.insert(keys[i], data[i])
		ett, err := json.Marshal(tt)
		assert.NoError(t, err)
		fmt.Println(string(ett))
	}

	assert.True(t, tt.isBalanced())

	d4, found := tt.find("4")
	assert.True(t, found)

	d4td, ok := d4.(testdata)
	assert.True(t, ok)
	assert.Equal(t, "Berlin", d4td.City)
	assert.Equal(t, "Germany", d4td.Country)
}
