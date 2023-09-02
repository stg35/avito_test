package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/hibiken/asynq"
	"github.com/stg35/avito_test/internal/handler/dto"
)

type showSegmentsRequest struct {
	Id uint64 `uri:"id" binding:"required,min=1"`
}

// @Summary CreateUser
// @Tags user
// @Description create user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body dto.UserDto true "user info"
// @Success 201 {object} model.User
// @Failure 403
// @Failure 500
// @Router /api/user [post]
func (h *Handler) createUser(ctx *gin.Context) {
	var req dto.UserDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := h.service.CreateUser(req)
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

	ctx.JSON(http.StatusCreated, user)
}

// @Summary AddSegments
// @Tags user
// @Description add segments to user with TTL
// @ID add-segments
// @Accept json
// @Produce json
// @Param input body dto.ChangeSegmentDto true "user id, list of segment's name and TTL. Если хотите добавить сегмент пользователю без TTL, то ставьте TTL равным 0"
// @Success 204
// @Failure 400
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /api/user/addSegments [patch]
func (h *Handler) addSegments(ctx *gin.Context) {
	var req dto.ChangeSegmentDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := h.service.AddSegments(req)
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
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

	if req.TTL > 0 {
		opts := []asynq.Option{
			asynq.MaxRetry(10),
			asynq.ProcessIn(time.Duration(req.TTL) * time.Second),
		}
		err = h.taskDistributor.DistributeTaskSegmentExpiration(ctx, &req, opts...)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("ttl problem: %s", err)))
		}
	}

	ctx.JSON(http.StatusNoContent, statusResponse("Segments successfully added"))
}

// @Summary DeleteSegments
// @Tags user
// @Description Delete user's segments
// @ID delete-segments
// @Accept json
// @Produce json
// @Param input body dto.ChangeSegmentDto true "user id and list of segment's name"
// @Success 204
// @Failure 403
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/user/deleteSegments [patch]
func (h *Handler) deleteSegments(ctx *gin.Context) {
	var req dto.ChangeSegmentDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := h.service.DeleteSegments(req)
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
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

	ctx.JSON(http.StatusNoContent, statusResponse("Segments successfully deleted"))
}

// @Summary ShowSegments
// @Tags user
// @Description Show segments of user
// @ID show-segment
// @Produce json
// @Param user_id path integer true "user's id"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/user/showSegments/{user_id} [get]
func (h *Handler) showSegments(ctx *gin.Context) {
	var req showSegmentsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	segments, err := h.service.GetSegments(req.Id)
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, segments)
}
