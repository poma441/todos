package entity

type RequestAdditionalInfo struct {
	UserAgent string `json:"user-agent"`
	SrcIP     string `json:"ip"`
}
