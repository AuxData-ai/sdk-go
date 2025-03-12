package auxdataaisdkgo

import "time"

const (
	DEFAULT_URL                = "https://auxdata.ai"
	DEV_URL                    = "https://dev.auxdata.ai"
	DEFAULT_MAX_RETRIES        = 5
	DEFAULT_TIMEOUT            = 120 * time.Second
	BASE_ROUTE                 = "/api/v1"
	SEARCH_URL_ROUTE_AGENT     = BASE_ROUTE + "/agent/${agentid}/document"
	SEARCH_URL_ROUTE_CONTAINER = BASE_ROUTE + "/agent/${agentid}/container/${containerid}/document"
	UPLOAD_URL_ROUTE           = BASE_ROUTE + "/agent/${agentid}/container/${containerid}/document"
	CHAT_URL_ROUTE             = BASE_ROUTE + "/agent/${agentid}/chat"
	CHAT_CONTAINER_URL_ROUTE   = BASE_ROUTE + "/agent/${agentid}/container/${containerid}/chat"
	AISERVICE_URL_ROUTE        = BASE_ROUTE + "/agent/${agentid}/executeservice/${serviceid}"
)

type LlmParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AiServiceResult struct {
	Result             string                    `json:"answer"`
	NextQuestionResult AiServiceMultiAnswer      `json:"nextQuestion"`
	QuestionObject     ExecuteServiceStepCommand `json:"questionObject"`
}

type AiServiceMultiAnswer struct {
	Prompt  string            `json:"questions"`
	Results []AiServiceResult `json:"answers"`
}

type ExecuteServiceResult struct {
	Command        ExecuteServiceCommand `json:"command"`
	MulitResults   AiServiceMultiAnswer  `json:"answer"`
	Error          string                `json:"error"`
	BackgroundMode bool                  `json:"background"`
}

type ExecuteServiceCommand struct {
	AgentId        int64          `json:"botid"`
	ServiceId      int64          `json:"templateid"`
	Parameters     []LlmParameter `json:"parameters,omitempty"`
	BackgroundMode bool           `json:"backgroundmodepossible"`
}

type UserConfig struct {
	Role       string `json:"role"`
	Behavior   string `json:"behavior"`
	TextSample string `json:"textsample"`
	Language   string `json:"language"`
}

type User struct {
	Id             int64      `json:"id"`
	KeycloakId     string     `json:"keycloakuserId"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"name"`
	OrganisationId int64      `json:"organisationId"`
	Role           int64      `json:"role"`
	Email          string     `json:"email"`
	Config         UserConfig `json:"config"`
}

type ChatCommunication struct {
	Timestamp time.Time `json:"timestamp"`
	Prompt    string    `json:"user"`
	Response  string    `json:"bot"`
}

type ResultExample struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type AiExecutionConfig struct {
	Provider                         string  `json:"provider"`
	ChainOfThougtPrompting           bool    `json:"chainofthought"`
	PalPrompting                     bool    `json:"palprompting"`
	ReflectionPrompting              bool    `json:"reflectionprompting"`
	LogicRules                       bool    `json:"logicrules"`
	PromptDebugging                  bool    `json:"promptdebugging"`
	RollingChunks                    bool    `json:"rollingchunks"`
	MaxParameterSize                 int64   `json:"maxparametersize"` // 1000 - 100000
	TopP                             int64   `json:"topp"`             // 0 - 99
	Temperature                      int64   `json:"temperature"`      // 0 - 150
	Language                         string  `json:"language"`
	Role                             string  `json:"role"`
	QualityGate                      float32 `json:"qualitygate"`
	GptFallback                      bool    `json:"gptfallback"`
	ChunkLimit                       int64   `json:"chunklimit"`
	DisableRAG                       bool    `json:"disablerag"`
	Containers                       []int64 `json:"containers"`
	Customizing                      string  `json:"customizing"`
	HistoryDepth                     int64   `json:"historydepth"`
	PresencePenalty                  float32 `json:"presencepenalty"`  // openai specific
	FrequencyPenalty                 float32 `json:"frequencypenalty"` // openai specific
	Greeting                         string  `json:"greeting"`
	ImageSize                        string  `json:"imagesize"`
	Voice                            string  `json:"voice"`
	SaveInKnowledgeDb                bool    `json:"saveinknowledgedb"`
	ContainerForSaveId               int64   `json:"containerforsaveid"`
	HierarchicalSearch               bool    `json:"hierarchicalsearch"`
	HierarchicalSearchChunkLimit     int64   `json:"hierarchicalsearchchunklimit"`
	ChatbotMemory                    bool    `json:"chatbotmemoryactive"`
	ChatbotShowSourceInformation     bool    `json:"chatbotshowsourceinformation"`
	ChatbotMaxServiceRecommandations int64   `json:"chatbotmaxservicerecommandations"`
	ChatbotQualityGateRecommendation int64   `json:"chatbotqualitygaterecommendation"`
	ChatbotQualityGateMemory         int64   `json:"chatbotqualitygatememory"`
}

type ProcessCallDefinition struct {
	Id     int64    `json:"id"`
	Name   string   `json:"name"`
	Key    string   `json:"key"`
	Params []string `json:"params"`
}

type Processor struct {
	Http     []ProcessCallDefinition `json:"http"`
	Function []ProcessCallDefinition `json:"function"`
}

type PostProcessor struct {
	Type    string                `json:"type"`
	Command string                `json:"command"`
	Call    ProcessCallDefinition `json:"call"`
}

type ExecuteServiceStepCommand struct {
	Prompt                  string              `json:"question"`
	User                    User                `json:"user"`
	AgentId                 int64               `json:"botid"`
	ResultType              string              `json:"answertype"`
	Title                   string              `json:"title"`
	Displaytype             string              `json:"displaytype"`
	ComUuid                 string              `json:"comuuid"`
	History                 []ChatCommunication `json:"history"`
	Examples                []ResultExample     `json:"examples"`
	OutputFormatDescription string              `json:"outputformatdescription"`
	CommandConfig           AiExecutionConfig   `json:"queryconfig"`
	Preprocessing           Processor           `json:"prepocessing"`
	Postprocessing          PostProcessor       `json:"postpocessing"`
}

type ChatResult struct {
	Command            ExecuteServiceStepCommand `json:"command"`
	Result             string                    `json:"answer"`
	Context            string                    `json:"context"`
	Error              string                    `json:"error"`
	ComUuid            string                    `json:"comuuid"`
	InformationSources []SearchChunkResult       `json:"sources"`
}

type Chat struct {
	Prompt  string `json:"question"`
	ComUuid string `json:"comuuid"`
}

type SearchChunkResult struct {
	Chunk            string  `json:"chunk"`
	DocumentId       string  `json:"documentId"`
	Score            float32 `json:"score"`
	Name             string  `json:"name"`
	Link             string  `json:"link"`
	ChunkId          int64   `json:"chunkid"`
	CreationDate     string  `json:"creationdate"`
	LastModifiedDate string  `json:"lastmodifieddate"`
	AgentId          int64   `json:"agentid"`
}

type Search struct {
	SearchString string  `json:"question"`
	QualityGate  float32 `json:"qualitygate"`
	ResultLimit  int64   `json:"resultLimit"`
}

type FileData struct {
	FileType    string
	Filename    string
	Link        string
	FileContent []byte
	DocumentId  string
}

type File struct {
	FileType string `json:"name"`
	Filename string `json:"type"`
	Content  string `json:"content"` // base64String
}

type AiServiceValue interface {
	toString() (string, error)
}

type AiServiceValueString struct {
	Value string
}

type AiServiceValueFile struct {
	Value File
}
