package techpalace

import (
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// panic("Please implement the WelcomeMessage() function")
    result := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
    return result
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// panic("Please implement the AddBorder() function")
    border := strings.Repeat("*", numStarsPerLine)
    result := border + "\n" + welcomeMsg + "\n" + border
    return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// panic("Please implement the CleanupMessage() function")
    result := strings.Trim(oldMsg, "*")
    result = strings.TrimLeft(result, "\n*")
    result = strings.TrimLeft(result, " ")
    result = strings.TrimRight(result, "\n*")
    result = strings.TrimRight(result, " ")
    return result
}
