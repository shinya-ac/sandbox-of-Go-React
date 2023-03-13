package answer

type Answer struct {
	Id         int64  `json:"Id"`
	Content    string `json:"Content"`
	FolderId   int64  `json:"FolderId,omitempty"`
	QuestionId int64  `json:"QuestionId,"`
}
