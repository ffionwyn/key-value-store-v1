package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "reader/store"
    "strconv"
    "strings"
    "time"
)

func main() {
    log.Println("Starting application")
    //Creates a continuous loop.
    for {
        //Reads user input.
        fmt.Println("Please input action")
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\r')
        if err != nil {
            fmt.Errorf("input not as expected %v", err)
            continue
        }
        input = strings.Trim(input, "\r")
        err = parseRequest(input)
        if err != nil {
            fmt.Println(err)
            continue
        }
    }
}

func parseRequest(input string) error {
    //Var to make every input lowercase to ensure it works.
    valueLower := strings.ToLower(input)
    //Strings.split using pipe as the separator.
    parts := strings.Split(valueLower, "|")
    action := parts[0]
    //Switch/cases for different actions that users can do.
    switch action {
    case "add":
        err := validateParts(parts)
        if err != nil {
            return err
        }
        name := parts[1]
        location := parts[2]
        age, err := strconv.Atoi(parts[3])
        if err != nil {
            return err
        }
        if !store.CheckPerson(name) {
            store.AddToStorage(name, location, age)
        } else {
            return fmt.Errorf("person already in the storage")
        }
    case "get":
        name := parts[1]
        loc, _ := store.GetPersonLocation(name)
        printTerminalMessages(fmt.Sprintf("Location found: %s", loc))
    case "delete":
        name := parts[1]
        message, err := store.DeletePerson(name)
        if err != nil {
            return err
        }
        log.Println(message)
    case "update":
        err := validateParts(parts)
        if err != nil {
            return err
        }
        name := parts[1]
        location := parts[2]
        age, err := strconv.Atoi(parts[3])
        if err != nil {
            return err
        }
        if store.CheckPerson(name) {
            store.AddToStorage(name, location, age)
        } else {
            return fmt.Errorf("person does not exist")
        }
    default:
        fmt.Println("Action not supported")
    }
    return nil
}

//Checks to see if full name, location and age has been entered, if not error.
func validateParts(parts []string) error {
    if len(parts) != 4 {
        return fmt.Errorf("expected full name, location and age got %d", len(parts))
    }
    return nil
}

//Prints terminal messages with a timestamp.
func printTerminalMessages(message string) {
    fmt.Printf("%v: %s\n", time.Now(), message)
}

