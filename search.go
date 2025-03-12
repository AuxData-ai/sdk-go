package auxdataaisdkgo

import (
	"encoding/json"
	"strconv"

	"github.com/AuxData-ai/utilities"
)

func (c *AuxDataClient) SearchOnAgent(agentId int64, search Search) ([]SearchChunkResult, error) {

	parameters := make(map[string]string)
	parameters["agentid"] = strconv.FormatInt(agentId, 10)
	route := utilities.ReplaceParametersInUrl(SEARCH_URL_ROUTE_AGENT, parameters)
	return c.search(route, search)
}

func (c *AuxDataClient) SearchOnContainer(agentId int64, containerId int64, search Search) ([]SearchChunkResult, error) {
	parameters := make(map[string]string)
	parameters["agentid"] = strconv.FormatInt(agentId, 10)
	parameters["containerid"] = strconv.FormatInt(containerId, 10)
	route := utilities.ReplaceParametersInUrl(SEARCH_URL_ROUTE_CONTAINER, parameters)
	return c.search(route, search)
}

func (c *AuxDataClient) search(route string, search Search) ([]SearchChunkResult, error) {
	var httpClient utilities.SimpleHttpClient
	httpClient.AddBearerAuthentificationToken(c.apiKey)
	httpClient.Method = "POST"

	httpClient.Url = c.url + route
	body, err := json.Marshal(search)

	if err != nil {
		return nil, err
	}

	httpClient.Body = string(body)
	result, err := httpClient.Execute()

	if err != nil {
		return nil, err
	}

	var results []SearchChunkResult
	err = json.Unmarshal([]byte(result), &results)

	return results, err

}
