package dtos

type NewAssets struct {
	AssetType string `validate:"required,oneof=category amenity" json:"asset_type"`
	Name      string `validate:"required"`
	IconURL   string `validate:"required,url" json:"icon_url"`
}

type ModifyAssets struct {
	AssetType string `validate:"required,oneof=category amenity" json:"asset_type"`
	Name      string `validate:"omitempty"`
	IconURL   string `validate:"omitempty,url" json:"icon_url"`
}
