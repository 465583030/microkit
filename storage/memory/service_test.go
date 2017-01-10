package memory

import (
	"testing"
)

func Test_List(t *testing.T) {
	config := DefaultConfig()
	newStorage, err := New(config)
	if err != nil {
		panic(err)
	}

	val := "my-val"

	err = newStorage.Create("key/one", val)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.Create("key/two", val)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	values, err := newStorage.List("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(values) != 2 {
		t.Fatal("expected", 2, "got", len(values))
	}
}

func Test_List_Invalid(t *testing.T) {
	config := DefaultConfig()
	newStorage, err := New(config)
	if err != nil {
		panic(err)
	}

	val := "my-val"

	err = newStorage.Create("key/one", val)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.Create("key/two", val)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	_, err = newStorage.List("ke")
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Service(t *testing.T) {
	newStorage, err := New(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	key := "test-key"
	value := "test-value"

	ok, err := newStorage.Exists(key)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if ok {
		t.Fatal("expected", false, "got", true)
	}

	err = newStorage.Create(key, value)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	ok, err = newStorage.Exists(key)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if !ok {
		t.Fatal("expected", true, "got", false)
	}

	v, err := newStorage.Search(key)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if v != value {
		t.Fatal("expected", value, "got", v)
	}

	err = newStorage.Delete(key)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	ok, err = newStorage.Exists(key)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if ok {
		t.Fatal("expected", false, "got", true)
	}

	v, err = newStorage.Search(key)
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
	if v != "" {
		t.Fatal("expected", "empty string", "got", v)
	}
}