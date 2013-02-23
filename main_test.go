package dig

import (
	"encoding/json"
	"testing"
)

var jsonTest = `{
	"Bool": false,
	"Int": 0,
	"Int8": 0,
	"Int16": 0,
	"Int32": 0,
	"Int64": 0,
	"Uint": 0,
	"Uint8": 0,
	"Uint16": 0,
	"Uint32": 0,
	"Uint64": 0,
	"Uintptr": 0,
	"Float32": 0,
	"Float64": 0,
	"bar": "",
	"bar2": "",
				"IntStr": "0",
	"PBool": true,
	"PInt": 2,
	"PInt8": 3,
	"PInt16": 4,
	"PInt32": 5,
	"PInt64": 6,
	"PUint": 7,
	"PUint8": 8,
	"PUint16": 9,
	"PUint32": 10,
	"PUint64": 11,
	"PUintptr": 12,
	"PFloat32": 14.1,
	"PFloat64": 15.1,
	"String": "",
	"PString": "16",
	"Map": null,
	"MapP": null,
	"PMap": {
		"17": {
			"Tag": "tag17"
		},
		"18": {
			"Tag": "tag18"
		}
	},
	"PMapP": {
		"19": {
			"Tag": "tag19"
		},
		"20": null
	},
	"EmptyMap": null,
	"NilMap": null,
	"Slice": null,
	"SliceP": null,
	"PSlice": [
		{
			"Tag": "tag20"
		},
		{
			"Tag": "tag21"
		}
	],
	"PSliceP": [
		{
			"Tag": "tag22"
		},
		null,
		{
			"Tag": "tag23"
		}
	],
	"EmptySlice": null,
	"NilSlice": null,
	"StringSlice": null,
	"ByteSlice": null,
	"Small": {
		"Tag": ""
	},
	"PSmall": null,
	"PPSmall": {
		"Tag": "tag31"
	},
	"Interface": null,
	"PInterface": 5.2
}`

func TestList(t *testing.T) {
	var err error

	list := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	var i int

	err = Pick(&list, &i, 5)

	if err != nil {
		t.Errorf("Test failed")
	}

	if i != 5 {
		t.Errorf("Test failed")
	}
}

func TestMatrix(t *testing.T) {
	var err error

	m33 := [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	var i int

	err = Pick(&m33, &i, 1, 1)

	if err != nil {
		t.Errorf("Test failed")
	}

	if i != 4 {
		t.Errorf("Test failed")
	}

	err = Pick(&m33, &i, 2, 0)

	if err != nil {
		t.Errorf("Test failed")
	}

	if i != 6 {
		t.Errorf("Test failed")
	}

	err = Pick(&m33, &i, 9, 9)

	if i != 0 {
		t.Errorf("Test failed")
	}

	if err == nil {
		t.Errorf("Test failed")
	}

	// Non assignable
	err = Pick(&m33, &i, 1)

	if i != 0 {
		t.Errorf("Test failed")
	}

	if err == nil {
		t.Errorf("Test failed")
	}
}

func TestFloatMatrix(t *testing.T) {
	var err error

	mf33 := [][]float64{
		[]float64{0.0, 1.0, 2.0},
		[]float64{3.0, 4.0, 5.0},
		[]float64{6.0, 7.0, 8.0},
	}
	var f float64

	err = Pick(&mf33, &f, 1, 2)

	if err != nil {
		t.Errorf("Test failed")
	}

	if f != 5.0 {
		t.Errorf("Test failed")
	}

}

func TestMap(t *testing.T) {
	var err error

	m := map[string]string{
		"Hello": "World",
	}

	var s string

	// Simple assignment
	err = Pick(&m, &s, "Hello")

	if err != nil {
		t.Errorf("Test failed")
	}

	if s != "World" {
		t.Errorf("Test failed")
	}

	m2 := map[string]map[string]map[string]string{
		"first": map[string]map[string]string{
			"first.1": map[string]string{
				"col.1": "a",
				"col.2": "b",
				"col.3": "c",
			},
		},
		"second": map[string]map[string]string{
			"second.2": map[string]string{
				"col.4": "d",
				"col.5": "e",
				"col.6": "f",
			},
		},
	}

	// Nested keys
	err = Pick(&m2, &s, "second", "second.2", "col.4")

	if err != nil {
		t.Errorf("Test failed")
	}

	if s != "d" {
		t.Errorf("Test failed")
	}

	// Non existent key
	err = Pick(&m2, &s, "third", "doest", "not", "exists")

	if err == nil {
		t.Errorf("Test failed")
	}

	if s != "" {
		t.Errorf("Test failed")
	}

	// Non assignable key
	err = Pick(&m2, &s, "second")

	if err == nil {
		t.Errorf("Test failed")
	}

	if s != "" {
		t.Errorf("Test failed")
	}

}

func TestJSON(t *testing.T) {

	var m map[string]interface{}

	json.Unmarshal([]byte(jsonTest), &m)

	var s string
	Pick(&m, &s, "PMap", "17", "Tag")

	if s != "tag17" {
		t.Errorf("Test failed.")
	}

	var f32 float32
	Pick(&m, &f32, "PFloat32")

	if f32 != float32(14.1) {
		t.Errorf("Test failed.")
	}

	var f64 float64
	Pick(&m, &f64, "PFloat32")

	if f64 != float64(14.1) {
		t.Errorf("Test failed.")
	}

	var b bool
	Pick(&m, &b, "PBool")

	if b != true {
		t.Errorf("Test failed.")
	}

	var ui64 uint64
	Pick(&m, &ui64, "PUint64")

	if ui64 != uint64(11) {
		t.Errorf("Test failed.")
	}

	var i int
	Pick(&m, &ui64, "String")

	if i != 0 {
		t.Errorf("Test failed.")
	}

	Pick(&m, &s, "PSlice", 1, "Tag")

	if s != "tag21" {
		t.Errorf("Test failed.")
	}

}

func TestJSON2(t *testing.T) {

	var m map[string]interface{}

	json.Unmarshal([]byte(jsonTest), &m)

	foo := map[string]string{
		"test": String(&m, "PMap", "17", "Tag"),
	}

	if foo["test"] != "tag17" {
		t.Errorf("Test failed.")
	}

}
