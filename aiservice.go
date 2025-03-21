package auxdataaisdkgo

import (
	"encoding/json"
	"strconv"

	"github.com/AuxData-ai/utilities"
)

// ExecuteAiService executes an AI service for a given agent and service ID with specified parameters.
//
// Parameters:
//   - agentId: The ID of the agent for which the AI service is to be executed.
//   - serviceId: The ID of the AI service to be executed.
//   - parameters: A map of parameters to be passed to the AI service.
//
// Returns:
//   - ExecuteServiceResult: The result of the AI service execution.
//   - error: An error object if there was an issue executing the service, otherwise nil.
func (c *AuxDataClient) ExecuteAiService(agentId int64, serviceId int64, parameters map[string]AiServiceValue) (ExecuteServiceResult, error) {
	routeparameters := make(map[string]string)
	routeparameters["agentid"] = strconv.FormatInt(agentId, 10)
	routeparameters["serviceid"] = strconv.FormatInt(serviceId, 10)
	route := utilities.ReplaceParametersInUrl(AISERVICE_URL_ROUTE, routeparameters)
	return c.execute(route, parameters)
}

func (c *AuxDataClient) execute(route string, parameters map[string]AiServiceValue) (ExecuteServiceResult, error) {
	var result ExecuteServiceResult

	httpClient := generateHttpClient(c, route, utilities.HTTP_METHOD_POST)
	convertedParams := make(map[string]string)

	for key, value := range parameters {

		stringValue, err := value.toString()

		if err != nil {
			return result, err
		}

		convertedParams[key] = stringValue
	}

	body, err := json.Marshal(convertedParams)

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

func (v *AiServiceValueString) toString() (string, error) {
	return v.Value, nil
}

func (v *AiServiceValueFile) toString() (string, error) {
	valueAsString, err := json.Marshal(v.Value)
	return string(valueAsString), err
}
