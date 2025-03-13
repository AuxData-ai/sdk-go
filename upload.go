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
	"time"

	"github.com/AuxData-ai/utilities"
)

func (c *AuxDataClient) UploadFile(agentId int64, containerId int64, file FileData) (UploadedFilesResult, error) {

	parameters := make(map[string]string)
	parameters["agentid"] = strconv.FormatInt(agentId, 10)
	parameters["containerid"] = strconv.FormatInt(containerId, 10)
	route := utilities.ReplaceParametersInUrl(UPLOAD_URL_ROUTE, parameters)

	return c.upload(route, file)
}

func (c *AuxDataClient) UploadFileFromDirectory(agentId int64, containerId int64, file FileDataToLoad) (UploadedFilesResult, error) {

	var result UploadedFilesResult
	var fileData FileData
	fileData.DocumentId = file.DocumentId
	fileData.FileType = filepath.Ext(file.FilePath)
	fileData.Filename = filepath.Base(file.FilePath)
	fileData.Link = file.Link

	data, err := os.ReadFile(file.FilePath)

	if err != nil {
		return result, err
	}

	fileData.FileContent = data

	return c.UploadFile(agentId, containerId, fileData)
}

func (c *AuxDataClient) upload(route string, fileData FileData) (UploadedFilesResult, error) {

	var result UploadedFilesResult
	file := bytes.NewBuffer(fileData.FileContent)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file to the form
	part, err := writer.CreateFormFile("file", fileData.Filename)
	if err != nil {
		return result, err
	}
	io.Copy(part, file)

	// Close the multipart form
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing multipart form:", err)
		return result, err
	}

	httpClient := generateHttpClient(c, route, utilities.HTTP_METHOD_POST)
	httpClient.AddHeader("link", fileData.Link)

	if fileData.DocumentId != "" {
		httpClient.AddHeader("documentid", fileData.DocumentId)
	}

	httpClient.ContentType = writer.FormDataContentType()

	return c.executeOwnHttpRequest(httpClient, body)
}

func (c *AuxDataClient) executeOwnHttpRequest(myHttpClient utilities.SimpleHttpClient, body *bytes.Buffer) (UploadedFilesResult, error) {

	var result UploadedFilesResult
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

	if myHttpClient.ContentType == "" {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", myHttpClient.ContentType)
	}

	httpClient.Timeout = time.Duration(time.Duration.Minutes(2))

	resp, err := httpClient.Do(req)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	resultBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(resultBody, &result)
	return result, err
}
