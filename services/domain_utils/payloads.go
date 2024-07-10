package services_domain_utils

type DNSVerificationPayload struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}

type SiteValidationPayload struct {
	Link string `json:"link"`
}

type InfoPayload struct {
	ID string `json:"id"`
}

type URLValidationPayload struct {
	Link string `json:"link"`
}

type TaskResultPayload struct {
	IsValid bool    `json:"isValid"`
	Message *string `json:"message,omitempty"`
	Error   *string `json:"error,omitempty"`
}
