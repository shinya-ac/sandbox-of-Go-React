package structure

type Question struct {
	Id       int64  `json:"Id"`
	Content  string `json:"Content"`
	FolderId int64  `json:"FolderId,omitempty"`
}

type Answer struct {
	Id         int64  `json:"Id"`
	Content    string `json:"Content"`
	FolderId   int64  `json:"FolderId,omitempty"`
	QuestionId int64  `json:"QuestionId,"`
}

type SelectedQA struct {
	Question Question `json:"question"`
	Answer   Answer   `json:"answer"`
}

type RequestPayload struct {
	SelectedQAs []SelectedQA `json:"selectedQAs"`
}
