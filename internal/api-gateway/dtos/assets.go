package dtos

// NewAssets represents the data required for create new asset
type NewAssets struct {
	AssetType string `validate:"required,oneof=category amenity" json:"asset_type"`
	Name      string `validate:"required"`
	IconURL   string `validate:"required,url" json:"icon_url"`
}

// ModifyAssets represents the data required for update the assets information
type ModifyAssets struct {
	AssetType string `validate:"required,oneof=category amenity" json:"asset_type"`
	Name      string `validate:"omitempty"`
	IconURL   string `validate:"omitempty,url" json:"icon_url"`
}
