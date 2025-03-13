package auxdataaisdkgo

import "testing"

const (
	API_KEY_UPLOAD      = "a3c336d8-b0eb-4e9a-9653-bbc6921dcf61"
	AGENT_ID_UPLOAD     = 119
	CONTAINER_ID_UPLOAD = 47
	FILEPATH_UPLOAD     = "./testdata/upload.pdf"
	FILE_LINK_UPLOAD    = "https://test.de"

	// Debug
	API_KEY_UPLOAD_DEBUG      = "1019d01b-d562-4319-b30c-c1da18da94e8"
	AGENT_ID_UPLOAD_DEBUG     = 151
	CONTAINER_ID_UPLOAD_DEBUG = 115
)

func TestUpload(t *testing.T) {

	c := NewAuxDataClient(API_KEY_UPLOAD, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var file FileDataToLoad

	file.FilePath = FILEPATH_UPLOAD
	file.Link = FILE_LINK_UPLOAD

	result, err := c.UploadFileFromDirectory(AGENT_ID_UPLOAD, CONTAINER_ID_UPLOAD, file)

	if err != nil {
		t.Error(err)
	}

	if len(result) != 1 {
		t.Fatalf("Expected 1 result but got %d", len(result))
	}

	if result[0].DocumentId == "" {
		t.Error("No DocumentId generated!")
	}
}

func TestUploadAndUpdate(t *testing.T) {

	c := NewAuxDataClient(API_KEY_UPLOAD, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var file FileDataToLoad

	file.FilePath = FILEPATH_UPLOAD
	file.Link = FILE_LINK_UPLOAD

	result, err := c.UploadFileFromDirectory(AGENT_ID_UPLOAD, CONTAINER_ID_UPLOAD, file)

	if err != nil {
		t.Error(err)
	}

	if len(result) != 1 {
		t.Fatalf("Expected 1 result but got %d", len(result))
	}

	if result[0].DocumentId == "" {
		t.Error("No DocumentId generated!")
	}

	firstDocumentId := result[0].DocumentId

	file.DocumentId = firstDocumentId

	result, err = c.UploadFileFromDirectory(AGENT_ID_UPLOAD, CONTAINER_ID_UPLOAD, file)

	if err != nil {
		t.Error(err)
	}

	if len(result) != 1 {
		t.Fatalf("Expected 1 result but got %d", len(result))
	}

	if result[0].DocumentId != firstDocumentId {
		t.Errorf("Expected documentId for update %s, but got %s", firstDocumentId, result[0].DocumentId)
	}
}

func TestUploadInvalidApiKey(t *testing.T) {

	c := NewAuxDataClient("invalid", DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var file FileDataToLoad

	file.FilePath = FILEPATH_UPLOAD
	file.Link = FILE_LINK_UPLOAD

	result, err := c.UploadFileFromDirectory(AGENT_ID_UPLOAD, CONTAINER_ID_UPLOAD, file)

	if err == nil {
		t.Error("Expected error from http client, but got nothing")
	}

	if len(result) != 0 {
		t.Fatalf("Expected 0 result but got %d", len(result))
	}
}

func TestUploadInvalidAgentId(t *testing.T) {

	c := NewAuxDataClient(API_KEY_UPLOAD, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var file FileDataToLoad

	file.FilePath = FILEPATH_UPLOAD
	file.Link = FILE_LINK_UPLOAD

	result, err := c.UploadFileFromDirectory(76868, CONTAINER_ID_UPLOAD, file)

	if err == nil {
		t.Error("Expected error from http client, but got nothing")
	}

	if len(result) != 0 {
		t.Fatalf("Expected 0 result but got %d", len(result))
	}
}

func TestUploadInvalidContainerId(t *testing.T) {

	c := NewAuxDataClient(API_KEY_UPLOAD, DEV_URL, DEFAULT_MAX_RETRIES, DEFAULT_TIMEOUT)

	var file FileDataToLoad

	file.FilePath = FILEPATH_UPLOAD
	file.Link = FILE_LINK_UPLOAD

	result, err := c.UploadFileFromDirectory(AGENT_ID_UPLOAD, 6786768, file)

	if err == nil {
		t.Error("Expected error from http client, but got nothing")
	}

	if len(result) != 0 {
		t.Fatalf("Expected 0 result but got %d", len(result))
	}
}
