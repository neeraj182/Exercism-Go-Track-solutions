package techpalace

import "strings"
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	message := "Welcome to the Tech Palace," + " " + strings.ToUpper(customer)
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	message := strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
	return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	message := strings.ReplaceAll(oldMsg, "*", "")
	message1 := strings.ReplaceAll(message, "\n", "")
	message2 := strings.TrimSpace(message1)

    return message2
}
