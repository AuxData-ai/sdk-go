package auxdataaisdkgo

import "testing"

const (
	API_KEY_AISERVICE  = "9f0dc80b-9d8a-42c4-a545-d26769998710"
	AGENT_ID_AISERVICE = 119
	SERVICE_ID         = 390

	// Debug
	API_KEY_AISERVICE_DEBUG  = "90e3d22d-0969-4a9c-8862-741f152cea3f"
	AGENT_ID_AISERVICE_DEBUG = 249
	SERVICE_ID_DEBUG         = 1583
)

func TestAiService(t *testing.T) {

	c := NewAuxDataClient(API_KEY_AISERVICE, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	parameters := make(map[string]AiServiceValue)

	var language AiServiceValueString
	language.Value = "Englisch"
	var content AiServiceValueString
	content.Value = "Das ist ein Test."

	parameters["language"] = &language
	parameters["content"] = &content

	result, err := c.ExecuteAiService(AGENT_ID_AISERVICE, SERVICE_ID, parameters)

	if err != nil {
		t.Error(err)
	}

	if result.Error != "" {
		t.Error(result.Error)
	}

	if len(result.MulitResults.Results) != 1 {
		t.Errorf("Expected 1 result but got %d", len(result.MulitResults.Results))
	}
}

func TestAiServiceInvalidToken(t *testing.T) {

	c := NewAuxDataClient("invalid", DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	parameters := make(map[string]AiServiceValue)

	var language AiServiceValueString
	language.Value = "Englisch"
	var content AiServiceValueString
	content.Value = "Das ist ein Test."

	parameters["language"] = &language
	parameters["content"] = &content

	result, err := c.ExecuteAiService(AGENT_ID_AISERVICE, SERVICE_ID, parameters)

	if err != nil {
		t.Error(err)
	}

	if result.Error != "" {
		t.Error(result.Error)
	}

	if len(result.MulitResults.Results) != 0 {
		t.Errorf("Expected 0 result but got %d", len(result.MulitResults.Results))
	}
}

func TestAiServiceInvalidAgentId(t *testing.T) {

	c := NewAuxDataClient(API_KEY_AISERVICE, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	parameters := make(map[string]AiServiceValue)

	var language AiServiceValueString
	language.Value = "Englisch"
	var content AiServiceValueString
	content.Value = "Das ist ein Test."

	parameters["language"] = &language
	parameters["content"] = &content

	result, err := c.ExecuteAiService(6767, SERVICE_ID, parameters)

	if err != nil {
		t.Error(err)
	}

	if result.Error != "" {
		t.Error(result.Error)
	}

	if len(result.MulitResults.Results) != 0 {
		t.Errorf("Expected 0 result but got %d", len(result.MulitResults.Results))
	}
}

func TestAiServiceInvalidServiceId(t *testing.T) {

	c := NewAuxDataClient(API_KEY_AISERVICE, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	parameters := make(map[string]AiServiceValue)

	var language AiServiceValueString
	language.Value = "Englisch"
	var content AiServiceValueString
	content.Value = "Das ist ein Test."

	parameters["language"] = &language
	parameters["content"] = &content

	result, err := c.ExecuteAiService(AGENT_ID_AISERVICE, 677868, parameters)

	if err == nil {
		t.Error("expected error, but returned nil")
	}

	if len(result.MulitResults.Results) != 0 {
		t.Errorf("Expected 0 result but got %d", len(result.MulitResults.Results))
	}
}

func TestAiServiceIWrongParameters(t *testing.T) {

	c := NewAuxDataClient(API_KEY_AISERVICE, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	parameters := make(map[string]AiServiceValue)

	var language AiServiceValueString
	language.Value = "Englisch"
	var content AiServiceValueString
	content.Value = "Das ist ein Test."

	parameters["blubb"] = &language
	parameters["blabb"] = &content

	result, err := c.ExecuteAiService(AGENT_ID_AISERVICE, SERVICE_ID, parameters)

	if err != nil {
		t.Error(err)
	}

	if result.Error != "" {
		t.Error(result.Error)
	}

	if len(result.MulitResults.Results) != 1 {
		t.Errorf("Expected 1 result but got %d", len(result.MulitResults.Results))
	}
}
