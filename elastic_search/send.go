package elastic_search

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/document"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"strings"
	"time"
)

func GetIndexES(hasPrefix string) string {
	// auction.backend-cms.product
	// auction.backend-cms.session
	return strings.ToLower(fmt.Sprintf("%s-%s.%s", esConf.Name, esConf.Environment, hasPrefix))
}

func (client *ElasticClient) InsertDocument(hasPrefix string, document document.Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := esapi.IndexRequest{
		Index:      GetIndexES(hasPrefix),
		Body:       strings.NewReader(base.JSONDebugDataString(document)),
		DocumentID: document.IDDoc,
		Refresh:    "true",
	}

	resp, err := req.Do(ctx, client.client)
	if err != nil {
		fmt.Println(fmt.Sprintf("ElasticClient::InsertDocument - Error: %+v", err))
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		var e map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return err
		}

		errRes := fmt.Errorf("[%s] %s: %s", resp.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		return errRes
	}

	return nil
}
