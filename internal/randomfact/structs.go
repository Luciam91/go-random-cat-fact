package randomfact

// FactAPIResponse API Response from REST API
type FactAPIResponse struct {
	Facts []fact `json:"all"`
}

type fact struct {
	ID   string `json:"_id"`
	Text string `json:"text"`
}
