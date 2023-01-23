package partyrobot

import (
    "fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
    result := fmt.Sprintf("Welcome to my party, %s!", name)
    return result
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    result := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
    return result
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    result := Welcome(name) + "\n"
    result += fmt.Sprintf("You have been assigned to table %.3d. Your table is %s, exactly %.1f meters from here.\n", table, direction, distance )
    result += fmt.Sprintf("You will be sitting next to %s.", neighbor)
    return result
}
