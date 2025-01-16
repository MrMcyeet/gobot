package interactionhandler

var (
	CommandRegistry = make(map[string]func())
)
