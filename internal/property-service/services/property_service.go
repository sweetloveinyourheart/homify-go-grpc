package services

import (
	"fmt"
	"homify-go-grpc/internal/property-service/models"
	"homify-go-grpc/internal/property-service/producers"
	"homify-go-grpc/internal/property-service/repositories"
	"homify-go-grpc/internal/property-service/types"
	kafka_configs "homify-go-grpc/internal/shared/kafka-configs"

	"gorm.io/gorm"
)

type IPropertyService interface {
	AddNewProperty(
		assetIds types.PropertyAssetIds,
		newProperty models.Property,
		newDestination models.Destination,
	) (bool, error)
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

	property := models.Property{
		HostId:      uint(newProperty.HostId),
		Title:       newProperty.Title,
		Description: newProperty.Description,
		Price:       newProperty.Price,
	}

	s.repo.Association(&property, "Category").Append(category)
	s.repo.Association(&property, "Amenity").Append(amenity)

	s.repo.CreateProperty(&property)

	destination := models.Destination{
		Country:   newDestination.Country,
		City:      newDestination.City,
		Latitude:  newDestination.Latitude,
		Longitude: newDestination.Longitude,
		Property:  property,
	}

	s.destinationRepo.CreateDestination(&destination)

	// Publish to kafka
	context := kafka_configs.GetContext()
	s.producer.ProduceMessages(context.SearchTopic, "Hi there !")

	return true, nil
}
