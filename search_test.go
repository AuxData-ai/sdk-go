package auxdataaisdkgo

import "testing"

const (
	API_KEY_SEARCH      = "f2607347-6d1a-4e2d-b880-e9dd0230b161"
	AGENT_ID_SEARCH     = 119
	CONTAINER_ID_SEARCH = 47
)

func TestSearchOnAgent(t *testing.T) {
	// todo

	c := NewAuxDataClient(API_KEY_SEARCH, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var search Search
	search.QualityGate = 45
	search.ResultLimit = 5
	search.SearchString = "Wie lautet das WLAN Passwort?"

	results, err := c.SearchOnAgent(AGENT_ID_SEARCH, search)

	if err != nil {
		t.Error(err)
	}

	if results.Error != "" {
		t.Error(results.Error)
	}

	found := len(results.Results)

	if found == 0 || int64(found) > search.ResultLimit {
		t.Errorf("Expected result count between 0 and %d but got %d", search.ResultLimit, found)
	}
}

func TestSearchOnContainer(t *testing.T) {
	// todo

	c := NewAuxDataClient(API_KEY_SEARCH, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var search Search
	search.QualityGate = 45
	search.ResultLimit = 5
	search.SearchString = "Wie lautet das WLAN Passwort?"

	results, err := c.SearchOnContainer(AGENT_ID_SEARCH, CONTAINER_ID_SEARCH, search)

	if err != nil {
		t.Error(err)
	}

	if results.Error != "" {
		t.Error(results.Error)
	}

	found := len(results.Results)

	if found == 0 || int64(found) > search.ResultLimit {
		t.Errorf("Expected result count between 0 and %d but got %d", search.ResultLimit, found)
	}
}

func TestSearchOnContainerInvalidApiKey(t *testing.T) {
	// todo

	c := NewAuxDataClient("invalid", DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var search Search
	search.QualityGate = 45
	search.ResultLimit = 5
	search.SearchString = "Wie lautet das WLAN Passwort?"

	results, err := c.SearchOnContainer(AGENT_ID_SEARCH, CONTAINER_ID_SEARCH, search)

	if err == nil {
		t.Error("expected error, but returned nil")
	}

	if results.Error != "" {
		t.Error(results.Error)
	}

	found := len(results.Results)

	if found > 0 {
		t.Errorf("Expected result count  0 but got %d", found)
	}
}

func TestSearchOnContainerWrongContainerId(t *testing.T) {
	// todo

	c := NewAuxDataClient(API_KEY_SEARCH, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var search Search
	search.QualityGate = 45
	search.ResultLimit = 5
	search.SearchString = "Wie lautet das WLAN Passwort?"

	results, err := c.SearchOnContainer(AGENT_ID_SEARCH, 899898, search)

	if err != nil {
		t.Error(err)
	}

	if results.Error != "" {
		t.Error(results.Error)
	}

	found := len(results.Results)

	if found > 0 {
		t.Errorf("Expected result count  0 but got %d", found)
	}
}

func TestSearchOnContainerWrongAgentId(t *testing.T) {
	// todo

	c := NewAuxDataClient(API_KEY_SEARCH, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var search Search
	search.QualityGate = 45
	search.ResultLimit = 5
	search.SearchString = "Wie lautet das WLAN Passwort?"

	results, err := c.SearchOnContainer(899898, CONTAINER_ID_SEARCH, search)

	if err == nil {
		t.Error("expected error, but returned nil")
	}

	if results.Error != "" {
		t.Error(results.Error)
	}

	found := len(results.Results)

	if found > 0 {
		t.Errorf("Expected result count  0 but got %d", found)
	}
}

func TestSearchOnContainerSearchNotPossibleToFInd(t *testing.T) {
	// todo

	c := NewAuxDataClient(API_KEY_SEARCH, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var search Search
	search.QualityGate = 45
	search.ResultLimit = 5
	search.SearchString = "WWs essen Schwaben zum Frühstück?"

	results, err := c.SearchOnContainer(AGENT_ID_SEARCH, CONTAINER_ID_SEARCH, search)

	if err != nil {
		t.Error(err)
	}

	if results.Error != "" {
		t.Error(results.Error)
	}

	found := len(results.Results)

	if found > 0 {
		t.Errorf("Expected result count  0 but got %d", found)
	}
}
