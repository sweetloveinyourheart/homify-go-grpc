package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type IElasticSearchClient interface {
	IndexDocument(index string, id string, data interface{}) error
	GetDocument(index string, id string) (*esapi.Response, error)
	UpdateDocument(index string, id string, data interface{}) error
	DeleteDocument(index string, id string) error
}

type ElasticSearchClient struct {
	es *elasticsearch.Client
}

func NewElasticSearchClient() IElasticSearchClient {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	return &ElasticSearchClient{
		es: client,
	}
}

func (c *ElasticSearchClient) IndexDocument(index string, id string, data interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return err
	}

	res, err := c.es.Index(index, &buf, c.es.Index.WithDocumentID(id))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.Status())
	}

	return nil
}

func (c *ElasticSearchClient) GetDocument(index string, id string) (*esapi.Response, error) {
	res, err := c.es.Get(index, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ElasticSearchClient) UpdateDocument(index string, id string, data interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return err
	}

	res, err := c.es.Update(index, id, &buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error updating document: %s", res.Status())
	}

	return nil
}

func (c *ElasticSearchClient) DeleteDocument(index string, id string) error {
	res, err := c.es.Delete(index, id)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error deleting document: %s", res.Status())
	}

	return nil
}
