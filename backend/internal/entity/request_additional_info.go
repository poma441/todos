package entity

type RequestAdditionalInfo struct {
	UserAgent string `json:"user-agent"`
	IP        string `json:"ip"`
}
