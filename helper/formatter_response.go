package formatter_response

var version = "1.1.0"

const Application = "Golang"

func GetVersion() string {
	return version
}

func SetVersion(input string) {
	version = input
}

func Success(code string, message string) string {
	return "Code: " + code + " Message: " + message
}