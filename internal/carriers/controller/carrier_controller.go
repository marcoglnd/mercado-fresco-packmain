package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/domain"
)

type CarrierController struct {
	service domain.CarrierService
}

func NewCarrierController(cs domain.CarrierService) *CarrierController {
	return &CarrierController{service: cs}
}

// @Summary Create carrier
// @Tags Carriers
// @Description Add a new carrier checking for duplicate carriers cid before
// @Accept json
// @Produce json
// @Param carrier body carriers.CreateCarrierInput true "Carrier to create"
// @Success 201 {object} carriers.Carrier
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /carriers [post]
func (cc *CarrierController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var carrierInput domain.CreateCarrierInput
		if err := ctx.ShouldBindJSON(&carrierInput); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		if err := cc.service.IsCidAvailable(ctx, carrierInput.Cid); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusConflict,
				gin.H{"error": err.Error()},
			)
			return
		}

		carrier := domain.Carrier{
			Cid:         carrierInput.Cid,
			CompanyName: carrierInput.CompanyName,
			Address:     carrierInput.Address,
			Telephone:   carrierInput.Telephone,
			LocalityId:  carrierInput.LocalityId,
		}

		createdCarrier, err := cc.service.Create(ctx, &carrier)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{"error": err.Error()},
			)
			return
		}

		ctx.JSON(
			http.StatusCreated, createdCarrier,
		)
	}
}
