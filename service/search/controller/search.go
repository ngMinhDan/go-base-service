package controller

import (
	"base/pkg/constant"
	"base/pkg/elastic"
	"base/pkg/log"
	"base/pkg/router"
	"base/pkg/utils"
	"base/service/search/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/go-chi/chi"
)

const _indexSample = "sample"
const _fieldName = "content" // about message in this case

// Signin Function: Check Input, Create JWT Token
func SearchWithKeyWord(w http.ResponseWriter, r *http.Request) {

	keyword := chi.URLParam(r, "key_word")
	if keyword == "" {
		router.ResponseBadGateway(w, constant.MissingFieldInputed, constant.MissingFieldInputed)
		return
	}
	listMessages, err := searchMessages(elastic.ES, _indexSample, _fieldName, keyword)
	if err != nil {
		router.ResponseInternalError(w, constant.ElasticSearchError, err.Error())
		return
	} else {
		if listMessages == nil && err == nil {
			router.ResponseSuccess(w, constant.SearchWithNotResults, constant.SearchWithNotResults)
			return
		}
	}
	router.ResponseSuccessWithData(w, constant.SearchWithResults, constant.SearchWithResults, listMessages)
}

// Search For Items in Elasticsearch
// And Map Data For List Iteams
func searchMessages(es *elastic.ElasticSearch, indexName, fieldName, keyWord string) ([]model.Message, error) {
	var listMessages []model.Message

	// Define A Query
	// You Can Define More Query Below, Multi Match, Match ...
	var buf strings.Builder
	buf.WriteString(`{
		"query": {
			"match": {
				"` + fieldName + `": "` + keyWord + `"
			}
		}
	}`)

	// Setup Request
	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(buf.String()),
	}
	// Search
	res, err := req.Do(context.Background(), es.Client)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("Elasticsearch search error: %s", res.Status())
	}

	result := make(map[string]interface{})

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println(log.LogLevelError, "decode-body", err.Error())
		return nil, err
	}
	// Loop Response
	// Map And Append Document
	if result["hits"] == nil {
		return nil, nil
	}

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var mess = model.Message{}

		utils.Mapping(hit.(map[string]interface{})["_source"], &mess)
		listMessages = append(listMessages, mess)
	}

	return listMessages, nil

}
