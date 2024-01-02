package models

type BCFact struct {
	ID   int    `json:"fact_id,omitempty"`
	Fact string `json:"fact_text"`
}
