package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"your.module/name/domain/dto"
	"your.module/name/internal/response"
	"your.module/name/usecase/social"
)

type CommentController struct {
	Create *social.CreateCommentUsecase
	Delete *social.DeleteCommentUsecase
}

func NewCommentController(cu *social.CreateCommentUsecase, du *social.DeleteCommentUsecase) *CommentController {
	return &CommentController{Create: cu, Delete: du}
}

func (cc *CommentController) CreateHandler(c echo.Context) error {
	userID, ok := c.Get("userID").(int64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, response.Error{Code: 401, Message: "Unauthorized"})
	}
	uid := int64(userID)
	pid, _ := strconv.ParseInt(c.Param("postID"), 10, 64)
	var req dto.CreateCommentRequest
	if err := c.Bind(&req); err != nil || req.Body == "" {
		return c.JSON(http.StatusBadRequest, response.Error{Code: 400, Message: "invalid body"})
	}
	res, err := cc.Create.Execute(c.Request().Context(), uid, pid, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, response.Success[dto.CommentResponse]{Code: 201, Message: "ok", Data: *res})
}

func (cc *CommentController) DeleteHandler(c echo.Context) error {
	uid, ok := c.Get("userID").(int64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, response.Error{Code: 401, Message: "Unauthorized"})
	}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := cc.Delete.Execute(c.Request().Context(), id, uid); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Success[any]{Code: 200, Message: "ok", Data: map[string]any{"id": id}})
}
