package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"your.module/name/domain/dto"
	"your.module/name/internal/response"
	"your.module/name/usecase/social"
)

type PostController struct {
	Create *social.CreatePostUsecase
	Update *social.UpdatePostUsecase
	Delete *social.DeletePostUsecase
	List   *social.ListFeedUsecase
}

func NewPostController(c *social.CreatePostUsecase, u *social.UpdatePostUsecase, d *social.DeletePostUsecase, l *social.ListFeedUsecase) *PostController {
	return &PostController{Create: c, Update: u, Delete: d, List: l}
}

func (pc *PostController) CreateHandler(c echo.Context) error {
	uid := c.Get("x-user-id").(int64)
	var req dto.CreatePostRequest
	if err := c.Bind(&req); err != nil || req.Body == "" {
		return c.JSON(http.StatusBadRequest, response.Error{Code: 400, Message: "invalid body"})
	}
	res, err := pc.Create.Execute(c.Request().Context(), uid, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, response.Success[dto.PostResponse]{Code: 201, Message: "ok", Data: *res})
}

func (pc *PostController) UpdateHandler(c echo.Context) error {
	uid := c.Get("x-user-id").(int64)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req dto.UpdatePostRequest
	if err := c.Bind(&req); err != nil || req.Body == "" {
		return c.JSON(http.StatusBadRequest, response.Error{Code: 400, Message: "invalid body"})
	}
	if err := pc.Update.Execute(c.Request().Context(), id, uid, req.Body); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Success[any]{Code: 200, Message: "ok", Data: map[string]any{"id": id}})
}

func (pc *PostController) DeleteHandler(c echo.Context) error {
	uid := c.Get("x-user-id").(int64)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := pc.Delete.Execute(c.Request().Context(), id, uid); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Success[any]{Code: 200, Message: "ok", Data: map[string]any{"id": id}})
}

func (pc *PostController) FeedHandler(c echo.Context) error {
	uid := c.Get("x-user-id").(int64)
	posts, err := pc.List.Execute(c.Request().Context(), uid, 20, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Success[[]*dto.PostResponse]{Code: 200, Message: "ok", Data: posts})
}
