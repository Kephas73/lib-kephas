package elastic_search

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Kephas73/lib-kephas/document"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"io/ioutil"
)

func (client *ElasticClient) SearchDocument(ctx context.Context, hasPrefix string, bodyQuery io.Reader) (*document.ResponseElastic, error) {

	listQueries := make([]func(request *esapi.SearchRequest), 0)
	listQueries = append(listQueries, client.client.Search.WithIndex(GetIndexES(hasPrefix)))
	listQueries = append(listQueries, client.client.Search.WithBody(bodyQuery))
	listQueries = append(listQueries, client.client.Search.WithContext(ctx))

	resp, err := client.client.Search(listQueries...)
	if err != nil {
		fmt.Println(fmt.Sprintf("ElasticClient::SearchDocument - Error: %v", err))

		return nil, err
	}

	defer resp.Body.Close()

	dataLog, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("ElasticClient::SearchDocument - Can not parse the body of response error: %v", err))

		return nil, err
	}

	if resp.IsError() {
		var e map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return nil, err
		}

		fmt.Println(fmt.Sprintf("ElasticClient::SearchDocument - Count request is error: %s", dataLog))

		errRes := fmt.Errorf("[%s] %s: %s", resp.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])

		return nil, errRes
	}

	result := &document.ResponseElastic{}
	if err = json.Unmarshal(dataLog, result); err != nil {
		fmt.Println(fmt.Sprintf("ElasticClient::SearchDocument - Can not parse the response error: %v", err))

		return nil, err
	}

	return result, nil
}

func (client *ElasticClient) CountDocument(ctx context.Context, hasPrefix string, bodyQuery io.Reader) (*document.ResponseElastic, error) {

	listQueries := make([]func(*esapi.CountRequest), 0)
	listQueries = append(listQueries, client.client.Count.WithIndex(GetIndexES(hasPrefix)))
	listQueries = append(listQueries, client.client.Count.WithBody(bodyQuery))
	listQueries = append(listQueries, client.client.Count.WithContext(ctx))

	resp, err := client.client.Count(listQueries...)
	if err != nil {
		fmt.Println(fmt.Sprintf("ElasticClient::CountDocument - Error: %v", err))

		return nil, err
	}

	defer resp.Body.Close()

	dataLog, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("ElasticClient::CountDocument - Can not parse the body of response error: %v", err))

		return nil, err
	}

	if resp.IsError() {
		var e map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return nil, err
		}

		fmt.Println(fmt.Sprintf("ElasticClient::CountDocument - Count request is error: %s", dataLog))

		errRes := fmt.Errorf("[%s] %s: %s", resp.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])

		return nil, errRes
	}

	result := &document.ResponseElastic{}
	if err = json.Unmarshal(dataLog, result); err != nil {
		fmt.Println(fmt.Sprintf("ElasticClient::CountDocument - Can not parse the response error: %v", err))

		return nil, err
	}

	return result, nil
}
