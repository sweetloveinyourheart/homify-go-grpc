package services

import (
	"homify-go-grpc/internal/property-service/models"
	"homify-go-grpc/internal/property-service/producers"
	"homify-go-grpc/internal/property-service/repositories"
	kafka_configs "homify-go-grpc/internal/shared/kafka-configs"

	"gorm.io/gorm"
)

type IPropertyService interface {
	AddNewProperty(newProperty models.Property) (bool, error)
}

type PropertyService struct {
	repo     repositories.IPropertyRepository
	producer producers.IPropertyProducer
}

func NewPropertyService(db *gorm.DB, p producers.IPropertyProducer) IPropertyService {
	return &PropertyService{
		repo:     repositories.NewPropertyRepository(db),
		producer: p,
	}
}

func (s *PropertyService) AddNewProperty(newProperty models.Property) (bool, error) {
	context := kafka_configs.GetContext()
	s.producer.ProduceMessages(context.SearchTopic, "Hi there !")

	return true, nil
}
