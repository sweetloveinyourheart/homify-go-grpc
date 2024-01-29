package services

import (
	"encoding/json"
	"homify-go-grpc/internal/shared/search"
)

type IPropertySearchService interface {
	SyncData(message []byte) error
}

type PropertySearchService struct {
	es        search.IElasticSearchClient
	esIndexes search.SearchIndexes
}

func NewPropertySearchService() IPropertySearchService {
	// ES client setup
	esClient := search.NewElasticSearchClient()
	esIndexes := search.GetSearchIndexes()

	return &PropertySearchService{
		es:        esClient,
		esIndexes: esIndexes,
	}
}

func (s *PropertySearchService) SyncData(message []byte) error {
	properties := []map[string]interface{}{}
	json.Unmarshal(message, &properties)

	for _, property := range properties {
		propertyId, ok := property["ID"].(string)
		if !ok {
			break
		}

		err := s.es.IndexDocument(s.esIndexes.PropertyIndex, propertyId, property)
		if err != nil {
			return err
		}
	}

	return nil
}
