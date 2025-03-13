package auxdataaisdkgo

import (
	"encoding/json"
	"strconv"

	"github.com/AuxData-ai/utilities"
)

// SearchOnAgent performs a search operation for a specific agent.
//
// Parameters:
//   - agentId: The ID of the agent to search on.
//   - search: The search criteria to use.
//
// Returns:
//   - DetailSearchResult: The result of the search operation.
//   - error: An error object if the search operation fails.
func (c *AuxDataClient) SearchOnAgent(agentId int64, search Search) (DetailSearchResult, error) {

	parameters := make(map[string]string)
	parameters["agentid"] = strconv.FormatInt(agentId, 10)
	route := utilities.ReplaceParametersInUrl(SEARCH_URL_ROUTE_AGENT, parameters)
	return c.search(route, search)
}

// SearchOnContainer performs a search operation within a specified container.
// It takes the following parameters:
// - agentId: The ID of the agent performing the search.
// - containerId: The ID of the container to search within.
// - search: The search criteria to be used.
//
// It returns a DetailSearchResult containing the search results, or an error if the search fails.
func (c *AuxDataClient) SearchOnContainer(agentId int64, containerId int64, search Search) (DetailSearchResult, error) {
	parameters := make(map[string]string)
	parameters["agentid"] = strconv.FormatInt(agentId, 10)
	parameters["containerid"] = strconv.FormatInt(containerId, 10)
	route := utilities.ReplaceParametersInUrl(SEARCH_URL_ROUTE_CONTAINER, parameters)
	return c.search(route, search)
}

func (c *AuxDataClient) search(route string, search Search) (DetailSearchResult, error) {
	var results DetailSearchResult
	httpClient := generateHttpClient(c, route, utilities.HTTP_METHOD_POST)
	body, err := json.Marshal(search)

	if err != nil {
		return results, err
	}

	httpClient.Body = string(body)
	result, err := httpClient.Execute()

	if err != nil {
		return results, err
	}

	err = json.Unmarshal([]byte(result), &results)
	return results, err
}
