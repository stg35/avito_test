package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/stg35/avito_test/internal/handler/dto"
)

type deleteSegmentRequest struct {
	Id uint64 `uri:"id" binding:"required,min=1"`
}

// @Summary CreateSegment
// @Tags segment
// @Description create segment
// @ID create-segment
// @Accept json
// @Produce json
// @Param input body dto.SegmentDto true "segment info"
// @Success 201 {object} model.Segment
// @Failure 400
// @Failure 403
// @Failure 500
// @Router /api/segment [post]
func (h *Handler) createSegment(ctx *gin.Context) {
	var req dto.SegmentDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	segment, err := h.service.CreateSegment(req)
	if err != nil {
		if pqError, ok := err.(pg.Error); ok {
			switch pqError.Field('C') {
			case "23505":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, segment)

}

// @Summary DeleteSegment
// @Tags segment
// @Description delete segment
// @ID delete-segment
// @Produce json
// @Param segment_id path integer true "segment id"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/segment/{segment_id} [delete]
func (h *Handler) deleteSegment(ctx *gin.Context) {
	var req deleteSegmentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := h.service.DeleteSegment(req.Id)
	if err != nil {
		if err == pg.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, statusResponse("Successfully deleted"))

}
