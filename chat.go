package auxdataaisdkgo

import (
	"encoding/json"
	"strconv"

	"github.com/AuxData-ai/utilities"
)

// ChatWithAllContainers sends a chat message to all containers associated with the given agent ID.
// It constructs the appropriate URL route with the agent ID and sends the chat message.
//
// Parameters:
//   - agentId: The ID of the agent to send the chat message to.
//   - chat: The chat message to be sent.
//
// Returns:
//   - ChatResult: The result of the chat operation.
//   - error: An error if the chat operation fails.
func (c *AuxDataClient) ChatWithAllContainers(agentId int64, chat Chat) (ChatResult, error) {

	parameters := make(map[string]string)
	parameters["agentid"] = strconv.FormatInt(agentId, 10)
	route := utilities.ReplaceParametersInUrl(CHAT_URL_ROUTE, parameters)
	return c.chat(route, chat)
}

// ChatWithOneContainers sends a chat message to a specific container identified by the given agent and container IDs.
// It constructs the appropriate URL route with the provided parameters and sends the chat message.
//
// Parameters:
//   - agentId: The ID of the agent initiating the chat.
//   - containerId: The ID of the container to which the chat message is sent.
//   - chat: The chat message to be sent.
//
// Returns:
//   - ChatResult: The result of the chat operation.
//   - error: An error object if an error occurred during the chat operation.
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
	httpClient.AddHeader("usermail", chat.UserMail)
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
