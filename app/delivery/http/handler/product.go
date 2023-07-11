package handler

import (
	"github.com/gin-gonic/gin"

	"pearshop_backend/app/delivery/http/payload"
	"pearshop_backend/app/delivery/http/presenter"
	"pearshop_backend/app/domain/entity"
	appErrors "pearshop_backend/app/errors"
	"pearshop_backend/app/usecase"
	"pearshop_backend/app/usecase/dto"
	"pearshop_backend/pkg/hashid"
)

type ProductHandler struct {
	productFindUsecase   usecase.ProductFind
	productUpdateUsecase usecase.ProductUpdate
	productCreateUsecase usecase.ProductCreate
	idHasher             hashid.IDHasher
}

func NewProductHandler(
	productFindUsecase usecase.ProductFind,
	productUpdateUsecase usecase.ProductUpdate,
	productCreateUsecase usecase.ProductCreate,
	idHasher hashid.IDHasher,
) *ProductHandler {
	return &ProductHandler{
		productFindUsecase:   productFindUsecase,
		productUpdateUsecase: productUpdateUsecase,
		productCreateUsecase: productCreateUsecase,
		idHasher:             idHasher,
	}
}

// Find return a list of product
// @Summary Find return a list of product
// @Description Find return a list of product
// @Tags products
// @Accept  json
// @Produce json
// @Param payloadQuery query payload.ProductFindRequest false "Conditions for filtering products"
// @Success 200 {array} presenter.Product
// @Router /products [get]
func (hdl *ProductHandler) Find(ctx *gin.Context) {
	req := payload.ProductFindRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		presenter.RenderErrors(ctx, appErrors.NewInvalidArgumentErr(
			appErrors.CodeInvalidPayload,
			"invalid query param",
			"",
		))

		return
	}

	if err := req.Validate(); err != nil {
		presenter.RenderErrors(ctx, err)

		return
	}

	data, err := hdl.productFindUsecase.Execute(ctx, dto.ProductFindRequest{}, entity.NoopPagingRequest{})
	if err != nil {
		presenter.RenderErrors(ctx, err)

		return
	}

	presenter.RenderData(ctx, presenter.FormProducts(hdl.idHasher, data), nil)
}

// Update modfiy product infomation by id
// @Summary Update modify product information by id
// @Description Update modify product information by id
// @Tags products
// @Accept  json
// @Produce json
// @Param	 payloadBody   body   payload.ProductSaveRequest true    "Body of request"
// @Success 200 {object} presenter.Product
// @Router /products/:id [put]
func (hdl *ProductHandler) Update(ctx *gin.Context) {
	rawID := ctx.Param("id")

	id, err := hdl.idHasher.Decode(rawID)
	if err != nil {
		presenter.RenderErrors(ctx, appErrors.NewInvalidArgumentErr(
			appErrors.CodeProductIDInvalid,
			"product id invalid",
			rawID,
		))
		return
	}

	req := payload.ProductSaveRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		presenter.RenderErrors(ctx, appErrors.NewInvalidArgumentErr(
			appErrors.CodeInvalidPayload,
			"invalid request body",
			"",
		))

		return
	}

	if err := req.Validate(); err != nil {
		presenter.RenderErrors(ctx, err)

		return
	}

	data, err := hdl.productUpdateUsecase.Execute(ctx, 1, id, req.ToDTO())
	if err != nil {
		presenter.RenderErrors(ctx, err)

		return
	}

	presenter.RenderData(ctx, presenter.FormProduct(hdl.idHasher, data), nil)
}

// Create add new product
// @Summary Create add new product
// @Description Create add new product
// @Tags products
// @Accept  json
// @Produce json
// @Param	 payloadBody   body   payload.ProductSaveRequest true    "Body of request"
// @Success 200 {object} presenter.Product
// @Router /products [post]
func (hdl *ProductHandler) Create(ctx *gin.Context) {
	req := payload.ProductSaveRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		presenter.RenderErrors(ctx, appErrors.NewInvalidArgumentErr(
			appErrors.CodeInvalidPayload,
			"invalid request body",
			"",
		))

		return
	}

	if err := req.Validate(); err != nil {
		presenter.RenderErrors(ctx, err)

		return
	}

	data, err := hdl.productCreateUsecase.Execute(ctx, 1, req.ToDTO())
	if err != nil {
		presenter.RenderErrors(ctx, err)

		return
	}

	presenter.RenderData(ctx, presenter.FormProduct(hdl.idHasher, data), nil)
}
