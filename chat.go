package auxdataaisdkgo

import (
	"encoding/json"
	"strconv"

	"github.com/AuxData-ai/utilities"
)

func (c *AuxDataClient) ChatWithAllContainers(agentId int64, chat Chat) (ChatResult, error) {

	parameters := make(map[string]string)
	parameters["agentid"] = strconv.FormatInt(agentId, 10)
	route := utilities.ReplaceParametersInUrl(CHAT_URL_ROUTE, parameters)
	return c.chat(route, chat)
}

func (c *AuxDataClient) ChatWithOneContainers(agentId int64, containerId int64, chat Chat) (ChatResult, error) {
	parameters := make(map[string]string)
	parameters["agentid"] = strconv.FormatInt(agentId, 10)
	parameters["containerid"] = strconv.FormatInt(containerId, 10)
	route := utilities.ReplaceParametersInUrl(CHAT_CONTAINER_URL_ROUTE, parameters)
	return c.chat(route, chat)
}

func (c *AuxDataClient) chat(route string, chat Chat) (ChatResult, error) {
	var result ChatResult

	httpClient := generateHttpClient(c, route, utilities.HTTP_METHOD_POST)
	body, err := json.Marshal(chat)

	if err != nil {
		return result, err
	}

	httpClient.Body = string(body)
	httpResult, err := httpClient.Execute()

	if err != nil {
		return result, err
	}

	err = json.Unmarshal([]byte(httpResult), &result)
	return result, err
}
