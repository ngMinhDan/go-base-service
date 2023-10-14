/*
Package elastic: Provides methods for working with Elasticsearcg.
Package Functionality: Connects to Elasticsearch, Insert, Remove item to Elasticsearch

Author: MinhDan <nguyenmd.works@gmail.com>
*/
package elastic

import (
	"base/pkg/config"
	"base/pkg/log"
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

type ElasticSearch struct {
	Client *elasticsearch.Client
	// You can add more info below
}

var ES *ElasticSearch

func init() {
	if strings.ToLower(config.Config.GetString("ENABLE_ELASTIC_SEARCH")) == "true" {
		// Do Elasticsearch Connection
		ES = elasticseachConnect()
	}
}

// ElasticSearch Connect Function
func elasticseachConnect() *ElasticSearch {

	host := strings.ToLower(config.Config.GetString("ES_HOST"))
	port := strings.ToLower(config.Config.GetString("ES_PORT"))
	address := "http://" + host + ":" + port

	// username := config.Config.GetString("ES_USER")
	// password := config.Config.GetString("ES_PASSWORD")
	cfg := elasticsearch.Config{
		Addresses: []string{address},
		// Username: username,
		// Password: password,
		// Add Transport Config
		// Add CACert: cert,
	}

	// Create New Elasticsearch Client
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println(log.LogLevelFatal, "create-new-client-elasticsearch", err.Error())
		return nil
	}
	// Check Client Info
	_, err = esClient.Info()
	if err != nil {
		log.Println(log.LogLevelFatal, "es-client es.Info()", err.Error())
		return nil
	} else {
		return &ElasticSearch{
			Client: esClient,
		}
	}
}

// Insert Item Into Index With DocumentID
func (es *ElasticSearch) Insert(index string, data any, documentID string) error {
	out, err := json.Marshal(data)
	if err != nil {
		log.Println(log.LogLevelError, "marshal-json-insert-elastic", err.Error())
		return err
	}
	// Set Up The Request Object
	req := esapi.IndexRequest{
		DocumentID: documentID,
		Index:      index,
		Body:       strings.NewReader(string(out)),
		Refresh:    "true",
	}

	// Insert Into Elasticsearch
	res, err := req.Do(context.Background(), es.Client)

	if err != nil {
		log.Println(log.LogLevelError, "insert-item-elasticsearch", err.Error())
		return err
	}
	defer res.Body.Close()

	responseRequest := ResponseRequest{}

	if res.IsError() {
		var e ResponseError
		err = json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return err
		} else {
			if e.Error.Reason != "" {
				log.Println(log.LogLevelError, "insert-item-elasticsearch", errors.New(e.Error.Reason+"->"+e.Error.CausedBy.Reason).Error())
				return errors.New(e.Error.Reason + "->" + e.Error.CausedBy.Reason)
			} else {
				return errors.New("response error elasticsearch invalid, can't find reason")
			}
		}
	} else {
		err := json.NewDecoder(res.Body).Decode(&responseRequest)
		if err != nil {
			return err
		} else {
			if strings.ToLower(responseRequest.Result) == "created" {
				return nil
			}
			return errors.New("Not Inserted")
		}
	}
}

// Remove Item With DocumentId In Index of Elasticsearch
func (es *ElasticSearch) Remove(documentID string, index string) error {
	req := esapi.DeleteRequest{
		DocumentID: documentID,
		Index:      index,
		Refresh:    "true",
	}

	// Do Request Remove Item
	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		log.Println(log.LogLevelError, "remove-item-elasticsearch", err.Error())
		return err
	}
	defer res.Body.Close()
	return nil
}

// Ping to check health to connect Elasticsearch
func (es *ElasticSearch) Ping() error {
	// Perform a ping to check the Elasticsearch cluster's availability
	res, err := es.Client.Ping()
	if err != nil {
		log.Println(log.LogLevelError, "ping-elasticsearch", err.Error())
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Println(log.LogLevelError, "ping-elasticsearch", err.Error())
		return errors.New("Error to connect with Elasticsearch")
	} else {
		return nil
	}
}
