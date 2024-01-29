package services

import (
	"encoding/json"
	"fmt"
	"homify-go-grpc/internal/property-service/models"
	"homify-go-grpc/internal/property-service/producers"
	"homify-go-grpc/internal/property-service/repositories"
	"homify-go-grpc/internal/property-service/types"
	broker "homify-go-grpc/internal/shared/broker"

	"gorm.io/gorm"
)

type IPropertyService interface {
	AddNewProperty(assetIds types.PropertyAssetIds, newProperty models.Property, newDestination models.Destination) (bool, error)
	SyncProperties()
}

type PropertyService struct {
	repo            repositories.IPropertyRepository
	destinationRepo repositories.IDestinationRepository
	categoryService ICategoryService
	amenityService  IAmenityService
	producer        producers.IPropertyProducer
}

func NewPropertyService(db *gorm.DB, p producers.IPropertyProducer) IPropertyService {
	return &PropertyService{
		repo:            repositories.NewPropertyRepository(db),
		destinationRepo: repositories.NewDestinationRepository(db),
		producer:        p,
		categoryService: NewCategoryService(db),
		amenityService:  NewAmenityService(db),
	}
}

func (s *PropertyService) AddNewProperty(
	assetIds types.PropertyAssetIds,
	newProperty models.Property,
	newDestination models.Destination,
) (bool, error) {

	category, categoryErr := s.categoryService.GetCategoryByID(assetIds.CategoryId)
	if categoryErr != nil {
		return false, fmt.Errorf("no category found")
	}

	amenity, amenityErr := s.amenityService.GetAmenityByID(assetIds.AmenityId)
	if amenityErr != nil {
		return false, fmt.Errorf("no amenity found")
	}

	destination := models.Destination{
		Country:   newDestination.Country,
		City:      newDestination.City,
		Latitude:  newDestination.Latitude,
		Longitude: newDestination.Longitude,
	}

	s.destinationRepo.CreateDestination(&destination)

	property := models.Property{
		HostId:      uint(newProperty.HostId),
		Title:       newProperty.Title,
		Description: newProperty.Description,
		Price:       newProperty.Price,
		Destination: destination,
	}

	s.repo.CreateProperty(&property)

	s.repo.Association(&property, "Category").Append(category)
	s.repo.Association(&property, "Amenity").Append(amenity)

	// Serialize the property object into JSON
	propertyJSON, err := json.Marshal(property)
	if err != nil {
		return false, fmt.Errorf("failed to serialize property: %v", err)
	}

	// Publish to kafka
	topics := broker.GetTopics()
	s.producer.ProduceMessages(topics.SyncProperties, propertyJSON)

	return true, nil
}

func (s *PropertyService) SyncProperties() {
	properties, pErr := s.repo.GetAllProperties(true)
	if pErr != nil {
		fmt.Println(pErr.Error())
		return
	}

	propertyJSON, err := json.Marshal(properties)
	if err != nil {
		fmt.Printf("failed to serialize property: %v", err)
		return
	}

	// Publish to kafka
	topics := broker.GetTopics()
	s.producer.ProduceMessages(topics.SyncProperties, propertyJSON)
}
