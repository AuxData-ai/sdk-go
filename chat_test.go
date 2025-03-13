package auxdataaisdkgo

import "testing"

const (
	API_KEY_CHAT      = "4550e9cd-c18a-48f9-8448-0caffba086dd"
	AGENT_ID_CHAT     = 119
	CONTAINER_ID_CHAT = 47

	// Debug localhost
	API_KEY_CHAT_DEBUG      = "84e66e6c-ac47-44cd-9544-9bde7642eb6d"
	AGENT_ID_CHAT_DEBUG     = 151
	CONTAINER_ID_CHAT_DEBUG = 47
)

func TestChatOnAgent(t *testing.T) {

	c := NewAuxDataClient(API_KEY_CHAT, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var chat Chat
	chat.ComUuid = ""
	chat.Prompt = "Wie lautet das WLAN Passwort?"

	result, err := c.ChatWithAllContainers(AGENT_ID_CHAT, chat)

	if err != nil {
		t.Error(err)
	}

	if result.ComUuid == "" {
		t.Error("No ComUuiD generated in Chat!")
	}

	if result.Result == "" {
		t.Error("No answer generated in Chat!")
	}
}

func TestChatOnSpecificContainer(t *testing.T) {
	// todo

	c := NewAuxDataClient(API_KEY_CHAT, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var chat Chat
	chat.ComUuid = ""
	chat.Prompt = "Wie lautet das WLAN Passwort?"

	result, err := c.ChatWithOneContainers(AGENT_ID_CHAT, CONTAINER_ID_CHAT, chat)

	if err != nil {
		t.Error(err)
	}

	if result.ComUuid == "" {
		t.Error("No ComUuiD generated in Chat!")
	}

	if result.Result == "" {
		t.Error("No answer generated in Chat!")
	}
}

func TestChatOnAgentWithInvalidAPIKey(t *testing.T) {

	c := NewAuxDataClient("invalid", DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var chat Chat
	chat.ComUuid = ""
	chat.Prompt = "Wie lautet das WLAN Passwort?"

	result, err := c.ChatWithAllContainers(AGENT_ID_CHAT, chat)

	if err == nil {
		t.Error("expected error, but returned nil")
	}

	if result.ComUuid != "" {
		t.Errorf("Expected no ComUuid but got %s", result.ComUuid)
	}

	if result.Result != "" {
		t.Errorf("Expected no result but got %s", result.Result)
	}
}

func TestChatOnAgentWithInvalidAgentId(t *testing.T) {

	c := NewAuxDataClient("invalid", DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var chat Chat
	chat.ComUuid = ""
	chat.Prompt = "Wie lautet das WLAN Passwort?"

	result, err := c.ChatWithAllContainers(68768, chat)

	if err == nil {
		t.Error("expected error, but returned nil")
	}

	if result.ComUuid != "" {
		t.Errorf("Expected no ComUuid but got %s", result.ComUuid)
	}

	if result.Result != "" {
		t.Errorf("Expected no result but got %s", result.Result)
	}
}

func TestChatOnAgentWithInvalidContainerId(t *testing.T) {

	c := NewAuxDataClient("invalid", DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var chat Chat
	chat.ComUuid = ""
	chat.Prompt = "Wie lautet das WLAN Passwort?"

	result, err := c.ChatWithOneContainers(AGENT_ID_CHAT, 7876, chat)

	if err == nil {
		t.Error("expected error, but returned nil")
	}

	if result.ComUuid != "" {
		t.Errorf("Expected no ComUuid but got %s", result.ComUuid)
	}

	if result.Result != "" {
		t.Errorf("Expected no result but got %s", result.Result)
	}
}
