package unicode

// Kode for Oppgave 4a
func Translate(expression string, language string) string {
	switch language {
	case "Norwegian":
		return expression
	case "Japanese":
		return "Japanese"
	case "Icelandic":
		return "Icelandic"
	}
	return "Error";
}
