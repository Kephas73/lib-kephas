package document

type ResponseOpenSearch struct {
	Count int `json:"count,omitempty"`
	Hits  struct {
		Total struct {
			Value int `json:"value,omitempty"`
		} `json:"Total,omitempty"`
	} `json:"hits,omitempty"`
	Aggregations struct {
		ResponseCodes struct {
			Buckets []struct {
				//KeyAsString string      `json:"key_as_string,omitempty"`
				Key       interface{} `json:"key,omitempty"`
				DocsCount int         `json:"doc_count,omitempty"`
			} `json:"buckets,omitempty"`
		} `json:"response_codes,omitempty"`
	} `json:"aggregations,omitempty"`
}

type TermStringBuilderOpenSearch struct {
	Term map[string]interface{} `json:"term,omitempty"`
}

type TermsStringBuilderOpenSearch struct {
	Terms map[string]interface{} `json:"terms,omitempty"`
}

type RangeStringBuilderOpenSearch struct {
	Range map[string]interface{} `json:"range,omitempty"`
}

type AggsConditionOpenSearch struct {
	ResponseCodes interface{} `json:"response_codes,omitempty"`
}

type QueryBuilderOpenSearch struct {
	Query struct {
		Bool struct {
			Must []interface{} `json:"must,omitempty"`
		} `json:"bool,omitempty"`
	} `json:"query,omitempty"`
	Aggs *AggsConditionOpenSearch `json:"aggs,omitempty"`
}
