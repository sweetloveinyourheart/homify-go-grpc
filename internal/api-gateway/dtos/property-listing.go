package dtos

type NewDestinationDTO struct {
	Country   string  `validate:"required,min=3,max=250" json:"country"`
	City      string  `validate:"required,min=3,max=250" json:"city"`
	Latitude  float32 `validate:"required" json:"lat"`
	Longitude float32 `validate:"required" json:"long"`
}

type NewPropertyDTO struct {
	Title       string            `validate:"required,min=3,max=250" json:"title"`
	Description string            `validate:"required,min=3" json:"description"`
	Price       float32           `validate:"required" json:"price"`
	CategoryId  uint              `validate:"required" json:"category_id"`
	AmenityId   uint              `validate:"required" json:"amenity_id"`
	Destination NewDestinationDTO `validate:"required" json:"destination"`
}

type EditDestinationDTO struct {
	Country   string  `validate:"omitempty,min=3,max=250" json:"country"`
	City      string  `validate:"omitempty,min=3,max=250" json:"city"`
	Latitude  float32 `validate:"omitempty" json:"lat"`
	Longitude float32 `validate:"omitempty" json:"long"`
}

type EditPropertyDTO struct {
	Title       string            `validate:"omitempty,min=3,max=250" json:"title"`
	Description string            `validate:"omitempty,min=3" json:"description"`
	Price       float32           `validate:"omitempty" json:"price"`
	CategoryId  uint              `validate:"omitempty" json:"category_id"`
	AmenityId   uint              `validate:"omitempty" json:"amenity_id"`
	Destination NewDestinationDTO `validate:"omitempty" json:"destination"`
}
