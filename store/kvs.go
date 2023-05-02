package store

import (
    "errors"
    "fmt"
)

// Struct - can accept all types, "person" has both location and age.
type person struct {
    location string
    age      int
}

func newPerson(loc string, age int) person {
    p := person{
        location: loc,
        age:      age,
    }
    return p
}

var personStorage = make(map[string]person)

//Checks to see if person is in storage before deleting person via fullname.
func DeletePerson(name string) (string, error) {
    if _, ok := personStorage[name]; !ok {
        return "", errors.New("person does not exist")
    }
    delete(personStorage, name)
    if _, ok := personStorage[name]; ok {
        return "", fmt.Errorf("failed to delete %s from store", name)
    }
    return fmt.Sprintf("%s is now deleted from store", name), nil
}

//Adds user input into storage.
func AddToStorage(name string, location string, age int) {
    fmt.Println("Hello " + name + " ")
    p := newPerson(location, age)
    personStorage[name] = p
    PrintPersonStorage()
}

//Prints names and locations of people already in the storage once somebody has been added.
func PrintPersonStorage() {
    fmt.Println("These are the names already inputted & being stored")
    for key, person := range personStorage {
        // Printing name and location inputted by user.
        fmt.Printf("name:%s \n", key)
        fmt.Printf("location:%s \n", person.location)
        fmt.Printf("age:%d \n", person.age)
    }
}

//Gets location of person you search via fullname.
func GetPersonLocation(name string) (string, error) {
    Person, found := personStorage[name]
    if !found {
        return "", fmt.Errorf("person does not exist")
    }
    return Person.location, nil
}

//Checks to see if person already exists.
func CheckPerson(name string) bool {
    _, exists := personStorage[name]
    return exists
}

