package document

import "encoding/json"

type ResponseElastic struct {
	Count int `json:"count,omitempty"`
	Hits  struct {
		Total struct {
			Value int `json:"value,omitempty"`
		} `json:"total,omitempty"`
		Hits []struct {
			ID         string          `json:"_id,omitempty"`
			Source     json.RawMessage `json:"_source,omitempty"`
			Highlights json.RawMessage `json:"highlight,omitempty"`
			Sort       []interface{}   `json:"sort,omitempty"`
		} `json:"hits,omitempty"`
	} `json:"hits,omitempty"`
	Aggregations struct {
		CountBy struct {
			Buckets []struct {
				KeyAsString string `json:"key_as_string,omitempty"`
				Key         int    `json:"key,omitempty"`
				DocsCount   int    `json:"doc_count,omitempty"`
				Total       struct {
					Value float64 `json:"value,omitempty"`
				} `json:"Total,omitempty"`
			} `json:"buckets,omitempty"`
		} `json:"count_by,omitempty"`
	} `json:"aggregations,omitempty"`
}

type TermStringBuilder struct {
	Term map[string]interface{} `json:"term,omitempty"`
}

type TermsStringBuilder struct {
	Terms map[string]interface{} `json:"terms,omitempty"`
}

type RangeStringBuilder struct {
	Range map[string]interface{} `json:"range,omitempty"`
}

type MultiMatchStringBuilder struct {
	MultiMatch struct {
		Query    string   `json:"query,omitempty"`
		Fields   []string `json:"fields,omitempty"`
		Operator string   `json:"operator,omitempty"`
	} `json:"multi_match,omitempty"`
}

type MatchStringBuilder struct {
	Match map[string]interface{} `json:"match,omitempty"`
}

type AggsCondition struct {
	CountBy struct {
		DateHistogram struct {
			Field            string `json:"field,omitempty"`
			CalendarInterval string `json:"calendar_interval,omitempty"`
		} `json:"date_histogram"`
	} `json:"count_by,omitempty"`
}

type AggsSumCondition struct {
	CountBy struct {
		TermsStringBuilder
		Aggs struct {
			Total struct {
				Sum struct {
					Field string `json:"field,omitempty"`
				} `json:"sum,omitempty"`
			} `json:"total,omitempty"`
		} `json:"aggs,omitempty"`
	} `json:"count_by,omitempty"`
}

type SortStringBuilder map[string]interface{}

type QueryBuilderES struct {
	Source []string `json:"_source,omitempty"`
	Query  struct {
		Bool struct {
			Must []interface{} `json:"must,omitempty"`
		} `json:"bool,omitempty"`
	} `json:"query,omitempty"`
	Aggs interface{}         `json:"aggs,omitempty"`
	Size int                 `json:"size,omitempty"`
	From int                 `json:"from,omitempty"`
	Sort []SortStringBuilder `json:"sort,omitempty"`
}

//GET auction_es-develop.product/_search
//{
//"_source": ["document.data.keywords.product_id"],
//"query": {
//"bool": {
//"must": [
//{
//"multi_match": {
//"query" : "tranh",
//"fields" : ["document.data.keywords.title"],
//"operator": "and"
//}
//},
//{
//"match": {
//"document.data.filters.category_id": 4
//}
//}
//]
//}
//},
//"from": 0,
//"size" : 1,
//"sort" : [ { "document.data.sorts.created_time": "desc" }]
//}

//GET auction_es-develop.product/_count
//{
//"query": {
//"bool": {
//"must": [
//{
//"multi_match": {
//"query" : "1",
//"fields" : ["document.data.keywords.title"],
//"operator": "and"
//}
//},
//{
//"match": {
//"document.data.filters.category_id": 4
//}
//}
//]
//}
//}
//}
