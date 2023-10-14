package elastic

type ResponseError struct {
	Error struct {
		RootCause []struct {
			Type   string `json:"type"`
			Reason string `json:"reason"`
		} `json:"root_cause"`
		Type     string `json:"type"`
		Reason   string `json:"reason"`
		CausedBy struct {
			Type   string `json:"type"`
			Reason string `json:"reason"`
		} `json:"caused_by"`
	} `json:"error"`
	Status int `json:"status"`
}

type ResponseRequest struct {
	Index         string  `json:"_index"`
	Type          string  `json:"_type"`
	ID            string  `json:"_id"`
	Version       float64 `json:"_version"`
	Result        string  `json:"result"`
	ForcedRefresh bool    `json:"forced_refresh"`
	Shards        struct {
		Total      float64 `json:"total"`
		Successful float64 `json:"successful"`
		Failed     float64 `json:"failed"`
	} `json:"_shards"`
	SeqNo       float64 `json:"_seq_no"`
	PrimaryTerm float64 `json:"_primary_term"`
}
