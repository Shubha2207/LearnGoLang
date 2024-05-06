package utils

/**
* check if the input is valid
 */
// first letter is capitalized so that this function can be exported
// and another packages can import it
func IsValidInput(name string) bool {
	return len(name) > 1
}
