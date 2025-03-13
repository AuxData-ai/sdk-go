package auxdataaisdkgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/AuxData-ai/utilities"
)

// UploadFile uploads a file to a specified container for a given agent.
//
// Parameters:
//   - agentId: The ID of the agent to which the file belongs.
//   - containerId: The ID of the container where the file will be uploaded.
//   - file: The file data to be uploaded.
//
// Returns:
//   - UploadedFilesResult: The result of the file upload operation.
//   - error: An error object if the upload fails, otherwise nil.
func (c *AuxDataClient) UploadFile(agentId int64, containerId int64, file FileData) ([]UploadedFilesResult, error) {

	parameters := make(map[string]string)
	parameters["agentid"] = strconv.FormatInt(agentId, 10)
	parameters["containerid"] = strconv.FormatInt(containerId, 10)
	route := utilities.ReplaceParametersInUrl(UPLOAD_URL_ROUTE, parameters)

	return c.upload(route, file)
}

// UploadFileFromDirectory uploads a file from a specified directory to a container.
// It takes the agent ID, container ID, and file data to load as parameters.
// The function reads the file content from the provided file path and uploads it using the UploadFile method.
// It returns the result of the upload and an error if any occurs during the process.
//
// Parameters:
//   - agentId: The ID of the agent performing the upload.
//   - containerId: The ID of the container to which the file is being uploaded.
//   - file: The file data to load, including the document ID, file path, and link.
//
// Returns:
//   - UploadedFilesResult: The result of the file upload.
//   - error: An error if any occurs during the file reading or uploading process.
func (c *AuxDataClient) UploadFileFromDirectory(agentId int64, containerId int64, file FileDataToLoad) ([]UploadedFilesResult, error) {

	var fileData FileData
	fileData.DocumentId = file.DocumentId
	fileData.FileType = filepath.Ext(file.FilePath)

	if index := strings.Index(fileData.FileType, "."); index != -1 {
		fileData.FileType = fileData.FileType[0:index] + fileData.FileType[index+1:]
	}

	fileData.Filename = filepath.Base(file.FilePath)
	fileData.Link = file.Link

	data, err := os.ReadFile(file.FilePath)

	if err != nil {
		return nil, err
	}

	fileData.FileContent = data

	return c.UploadFile(agentId, containerId, fileData)
}

func (c *AuxDataClient) upload(route string, fileData FileData) ([]UploadedFilesResult, error) {

	file := bytes.NewBuffer(fileData.FileContent)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file to the form
	part, err := writer.CreateFormFile("files", fileData.Filename)
	if err != nil {
		return nil, err
	}
	io.Copy(part, file)

	// Close the multipart form
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing multipart form:", err)
		return nil, err
	}

	httpClient := generateHttpClient(c, route, utilities.HTTP_METHOD_PUT)
	httpClient.AddHeader("link", fileData.Link)

	if fileData.DocumentId != "" {
		httpClient.AddHeader("documentid", fileData.DocumentId)
	}

	httpClient.ContentType = writer.FormDataContentType()

	return c.executeOwnHttpRequest(httpClient, body)
}

func (c *AuxDataClient) executeOwnHttpRequest(myHttpClient utilities.SimpleHttpClient, body *bytes.Buffer) ([]UploadedFilesResult, error) {

	var result []UploadedFilesResult
	// Create a new HTTP Client
	httpClient := &http.Client{}

	// Create a new HTTP Request
	req, err := http.NewRequest(string(myHttpClient.Method), myHttpClient.Url, body)

	if err != nil {
		return result, err
	}

	// Set the headers for the request
	for key, value := range myHttpClient.Headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("Content-Type", myHttpClient.ContentType)
	httpClient.Timeout = time.Duration(time.Duration.Minutes(2))

	resp, err := httpClient.Do(req)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return result, fmt.Errorf("error: %s", resp.Status)
	}

	resultBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(resultBody, &result)
	return result, err
}
