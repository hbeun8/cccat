package cmd

import (
	"testing"
)

// TestInsertValid calls map.Insert with a key value, checking
// for a valid return value.
func TestCatNakedValid(t *testing.T) {
	o := Object{datastore: []string{}}
    want := 1
    value, err := Insert(o, 1, "value")
    if value != want || err != nil {
        t.Errorf(`TestInsertValid:, %d, %x, %d`, value, err, want)
    }
}

// TestInsertNull calls map.Insert with a key value, checking
// checking for an error.
func TestInsertNull(t *testing.T) {
	o := Object{datastore: []string{}}
    want := 0
	value, err := Insert(o, 0, "")
    if value != want || err != nil {
        t.Errorf(`TestInsertNull:, Actual:, %d, Err, %x, Expected: %d`, value, err, want)
    }
}


// TestDelete calls greetings.Hello with an empty string,
// checking for an error.
func TestDelete(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

func (m *HashTable)validate_Insert() error{
	if m == nil {
		return errors.New("hashTable is nil")
	}
	if len(m.Key) == 0 {
		return errors.New("key is nil")
	}
	if len(m.Value) == 0 {
		return errors.New("value is nil")
	}
	return nil
}

func (m *HashTable)validate_Retrieve() error{
	if w := string(m.Key); len(w)==0 {
		return errors.New("key is nil")
	}
	return nil

}