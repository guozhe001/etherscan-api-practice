package model

type SourceCodeModel struct {
	Language string                            `json:"language"`
	Sources  map[string]map[string]interface{} `json:"sources"`
	Settings interface{}                       `json:"settings"` //如果是一个struct类型，则使用interface{}类型，很像java里面的Object
}

const KeyContent = "content"
