package common

// HookRequest is a request from a plugin provider to be registered to a specific hook
type HookRequest struct {
	Name     string
	Priority int
}

// HookResponse is a response from the server to the plugin provider
type HookResponse struct {
	Name        string
	Priority    int
	RequestName string
}
