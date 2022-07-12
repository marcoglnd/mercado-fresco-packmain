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

// @Summary Report Carriers
// @Tags Carriers
// @Description Get quantity of carriers by locality id
// @Accept json
// @Produce json
// @Param id query int true "locality ID"
// @Router /reportCarriers [get]
func (cc *CarrierController) ReportCarriers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reportCarriersInput struct {
			Id int64 `form:"id"`
		}
		if err := ctx.ShouldBindQuery(&reportCarriersInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var reports *[]domain.CarrierReport
		if reportCarriersInput.Id == 0 {
			allReports, err := cc.service.GetAllCarriersReport(ctx)
			if err != nil {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
				return
			}
			reports = allReports
		} else {
			customReport, err := cc.service.GetCarriersReportById(
				ctx,
				reportCarriersInput.Id,
			)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			reports = &[]domain.CarrierReport{
				*customReport,
			}
		}
		ctx.JSON(http.StatusOK, gin.H{"data": reports})
	}
}
