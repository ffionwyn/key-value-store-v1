package main

import (
    //"reflect"
    "errors"
    "testing"
)

// func TestAdd(t *testing.T) {
//  parseRequest("add|name lastname|location")
//  expectedPersonStorage := map[string]string{"name lastname": "location"}
//  if !reflect.DeepEqual(expectedPersonStorage, personStorage) {
//      t.Errorf("expected %v, got %v", expectedPersonStorage, personStorage)
//  }
// }

// func TestUpdate(t *testing.T) {
//  parseRequest("add|name lastname|location")
//  parseRequest("update|name lastname|newlocation")
//  expectedPersonStorage := map[string]string{"name lastname": "newlocation"}
//  if !reflect.DeepEqual(expectedPersonStorage, personStorage) {
//      t.Errorf("expected %v, got %v", expectedPersonStorage, personStorage)
//  }
// }

// func TestDelete(t *testing.T) {
//  parseRequest("add|name lastname|location")
//  parseRequest("delete|name lastname")
//  expectedPersonStorage := map[string]string{}
//  if !reflect.DeepEqual(expectedPersonStorage, personStorage) {
//      t.Errorf("expected %v, got %v", expectedPersonStorage, personStorage)
//  }
// }

// func TestGet(t *testing.T) {
//  parseRequest("add|name lastname|location")
//  parseRequest("get|name lastname")
//  expectedPersonStorage := map[string]string{}
//  if !reflect.DeepEqual(expectedPersonStorage, personStorage) {
//      t.Errorf("expected %v, got %v", expectedPersonStorage, personStorage)
//  }
// }

func TestValidateParts(t *testing.T) {

    input := []string{
        "add", "wrexham",
    }
    err := validateParts(input)
    if err == nil {
        t.Errorf("expected error")
    }
}

func TestValidatePartsTable(t *testing.T) {
    cases := []struct {
        name  string
        input []string
        err   error
    }{
        {"valid input", []string{"add", "ffion griffiths", "wrexham"}, nil},
        {"invalid input", []string{"add", "wrexham"}, errors.New("expected full name and location got 2")},
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            err := validateParts(tc.input)
            if err != nil && err.Error() != tc.err.Error() {
                t.Errorf("validate parts failed at case %s expected %q got %q", tc.name, tc.err.Error(), err.Error())
            }
        })
    }
}

