package handlers

import (
	"context"
	proto "homify-go-grpc/api/property"
	"homify-go-grpc/internal/api-gateway/dtos"
	"homify-go-grpc/internal/api-gateway/helpers"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AssetsHandler struct {
	grpcClient proto.PropertyClient
	validator  *validator.Validate
}

func NewAssetsHandler(c proto.PropertyClient, validate *validator.Validate) *AssetsHandler {
	return &AssetsHandler{
		grpcClient: c,
		validator:  validate,
	}
}

// @Tags Assets
// @Summary Get assets by asset type
// @Description Get assets based on the provided asset type
// @Accept json
// @Produce json
// @Param asset_type query string true "Asset type to filter by"
// @Success 200 {object} interface{} "The assets list"
// @Failure 400 {object} interface{} "Error"
// @Router /assets [get]
func (h *AssetsHandler) GetAssets(ctx *gin.Context) {
	assetType := ctx.Query("asset_type")
	if assetType == "" {
		ctx.JSON(400, gin.H{"error": "must provide asset_type on query params"})
		return
	}

	assetsCtx := context.Background()

	grpcReq := &proto.GetAssetsRequest{
		AssetType: assetType,
	}

	grpcRes, err := h.grpcClient.GetAssets(assetsCtx, grpcReq)
	if err != nil {
		log.Printf("Get assets thrown error: %v", err)
		ctx.JSON(400, gin.H{"error": "Get assets failed"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    grpcRes,
	})
}

// AddNewAsset handles the creation of a new asset.
// @Tags Assets
// @Summary Create a new asset
// @Description Creates a new asset based on the provided JSON payload.
// @ID addNewAsset
// @Accept json
// @Security Authorization
// @Produce json
// @Param input body dtos.NewAssets true "JSON payload containing data for the new asset"
// @Success 201 {object} interface{} "The asset was created successfully"
// @Failure 400 {object} interface{} "Failed to create the new asset due to validation errors or other issues"
// @Router /assets [post]
func (h *AssetsHandler) AddNewAsset(ctx *gin.Context) {
	newAssetData := dtos.NewAssets{}
	if bindError := ctx.ShouldBindJSON(&newAssetData); bindError != nil {
		ctx.JSON(400, gin.H{"error": bindError.Error()})
		return
	}

	validatorError := h.validator.Struct(newAssetData)
	if validatorError != nil {
		helpers.HandleValidationErrors(ctx, validatorError)
		return
	}

	assetsCtx := context.Background()

	grpcReq := &proto.AddAssetRequest{
		AssetType: newAssetData.AssetType,
		IconURL:   newAssetData.IconURL,
		Name:      newAssetData.Name,
	}

	grpcRes, err := h.grpcClient.AddAsset(assetsCtx, grpcReq)
	if err != nil {
		log.Printf("New asset thrown error: %v", err)
		ctx.JSON(400, gin.H{"error": "Create new asset failed"})
		return
	}

	if !grpcRes.Success {
		ctx.JSON(400, gin.H{
			"message": "Create new asset failed",
			"data":    grpcRes,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Created",
		"data":    grpcRes,
	})
}

// ModifyExistingAsset modifies an existing asset.
// @Tags Assets
// @Summary Modify an existing asset
// @Description Modifies an existing asset based on the provided JSON payload.
// @ID modifyExistingAsset
// @Param asset_id query string true "Asset id to modify"
// @Accept json
// @Security Authorization
// @Produce json
// @Param input body dtos.ModifyAssets true "JSON payload containing data for modifying the asset"
// @Success 201 {object} interface{} "The asset was updated successfully"
// @Failure 400 {object} interface{} "Failed to update the asset due to validation errors or other issues"
// @Router /assets/modify [put]
func (h *AssetsHandler) ModifyExistingAsset(ctx *gin.Context) {
	assetId := ctx.Query("asset_id")
	if assetId == "" {
		ctx.JSON(400, gin.H{"error": "must provide asset_id"})
		return
	}

	assetIdUint, parseErr := strconv.ParseUint(assetId, 10, 32)
	if parseErr != nil {
		ctx.JSON(400, gin.H{"error": "asset_id has invalid format"})
		return
	}

	modifyAssetData := dtos.ModifyAssets{}
	if bindError := ctx.ShouldBindJSON(&modifyAssetData); bindError != nil {
		ctx.JSON(400, gin.H{"error": bindError.Error()})
		return
	}

	validatorError := h.validator.Struct(modifyAssetData)
	if validatorError != nil {
		helpers.HandleValidationErrors(ctx, validatorError)
		return
	}

	assetsCtx := context.Background()

	grpcReq := &proto.ModifyAssetRequest{
		Id:        uint32(assetIdUint),
		AssetType: modifyAssetData.AssetType,
		IconURL:   modifyAssetData.IconURL,
		Name:      modifyAssetData.Name,
	}

	grpcRes, err := h.grpcClient.ModifyAsset(assetsCtx, grpcReq)
	if err != nil {
		log.Printf("Update asset thrown error: %v", err)
		ctx.JSON(400, gin.H{"error": "Update new asset failed"})
		return
	}

	if !grpcRes.Success {
		ctx.JSON(400, gin.H{
			"message": "Update new asset failed",
			"data":    grpcRes,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Updated",
		"data":    grpcRes,
	})
}

// DisableAnAsset disables an existing asset.
// @Tags Assets
// @Summary Disable an existing asset
// @Description Disables an existing asset based on the provided asset_id and asset_type.
// @ID disableAnAsset
// @Security Authorization
// @Param asset_id query string true "Asset ID to be disabled"
// @Param asset_type query string true "Type of the asset to be disabled"
// @Success 201 {object} interface{} "The asset was disabled successfully"
// @Failure 400 {object} interface{} "Failed to disable the asset due to validation errors or other issues"
// @Router /assets/disable [put]
func (h *AssetsHandler) DisableAnAsset(ctx *gin.Context) {
	assetId := ctx.Query("asset_id")
	assetType := ctx.Query("asset_type")

	if assetId == "" || assetType == "" {
		ctx.JSON(400, gin.H{"message": "Must provide necessary information about asset_id and asset_type"})
		return
	}

	assetsCtx := context.Background()

	assetIdUint, parseErr := strconv.ParseUint(assetId, 10, 32)
	if parseErr != nil {
		ctx.JSON(400, gin.H{"message": "Invalid asset_id format"})
		return
	}

	grpcReq := &proto.DisableAssetRequest{
		Id:        uint32(assetIdUint),
		AssetType: assetType,
	}

	grpcRes, err := h.grpcClient.DisableAsset(assetsCtx, grpcReq)
	if err != nil {
		log.Printf("Update asset thrown error: %v", err)
		ctx.JSON(400, gin.H{"error": "Update new asset failed"})
		return
	}

	if !grpcRes.Success {
		ctx.JSON(400, gin.H{
			"message": "Update new asset failed",
			"data":    grpcRes,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Updated",
		"data":    grpcRes,
	})
}
