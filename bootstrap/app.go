package bootstrap

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"your.module/name/api/controller"
	"your.module/name/api/router"
	gormrepo "your.module/name/repository/gorm"
	"your.module/name/usecase/auth"
	"your.module/name/usecase/social"
	useuser "your.module/name/usecase/user"
)

func NewAppWithDeps(db *gorm.DB, cfg Config) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover(), middleware.Logger(), middleware.RequestID())
	e.HideBanner = true
	e.HidePort = true

	// Repositories
	userRepo := gormrepo.NewUserRepository(db)
	sessRepo := gormrepo.NewSessionRepository(db) // <-- cần cho refresh/logout
	postRepo := gormrepo.NewPostRepository(db)
	cmtRepo := gormrepo.NewCommentRepository(db)
	likeRepo := gormrepo.NewLikeRepository(db)
	folRepo := gormrepo.NewFollowRepository(db)

	// Services
	tokenSvc := auth.NewHS256TokenService(cfg.JWTSecret)

	// Usecases
	// NewLoginUsecase(users, sessions, tokenSvc, accessTTL, refreshTTL)
	loginUC := auth.NewLoginUsecase(
		userRepo,
		sessRepo,
		tokenSvc,
		cfg.AccessTokenTTL,
		cfg.RefreshTokenTTL,
	)
	signupUC := auth.NewSignupUsecase(userRepo)

	// NewRefreshTokenUsecase(sessions, tokenSvc, accessTTLSeconds)
	// Nếu ctor của bạn nhận "giây", dùng dòng dưới:
	refreshUC := auth.NewRefreshTokenUsecase(
		sessRepo,
		tokenSvc,
		int64(cfg.AccessTokenTTL.Seconds()),
	)

	logoutUC := auth.NewLogoutUsecase(sessRepo)

	profileUC := useuser.NewProfileUsecase(userRepo)

	// Social usecases
	createPostUC := social.NewCreatePostUsecase(postRepo)
	updatePostUC := social.NewUpdatePostUsecase(postRepo)
	deletePostUC := social.NewDeletePostUsecase(postRepo)
	listFeedUC := social.NewListFeedUsecase(postRepo, folRepo)

	createCmtUC := social.NewCreateCommentUsecase(cmtRepo)
	deleteCmtUC := social.NewDeleteCommentUsecase(cmtRepo)

	likeUC := social.NewLikePostUsecase(likeRepo)
	followUC := social.NewFollowUserUsecase(folRepo)

	// Controllers
	authCtl := controller.NewAuthController(signupUC, loginUC, refreshUC, logoutUC)
	userCtl := controller.NewUserController(profileUC)
	postCtl := controller.NewPostController(createPostUC, updatePostUC, deletePostUC, listFeedUC)
	cmtCtl := controller.NewCommentController(createCmtUC, deleteCmtUC)
	lfCtl := controller.NewLikeFollowController(likeUC, followUC)

	// Routes
	api := router.RegisterRoutes(e)
	e.GET("/debug/routes", func(c echo.Context) error {
		type R struct{ Method, Path string }
		rs := make([]R, 0, len(e.Routes()))
		for _, r := range e.Routes() {
			rs = append(rs, R{Method: r.Method, Path: r.Path})
		}
		return c.JSON(200, rs)
	})
	router.MountAuth(api, authCtl)
	router.MountUser(api, userCtl, cfg.JWTSecret)
	router.MountPost(api, postCtl, cfg.JWTSecret)
	router.MountComment(api, cmtCtl, cfg.JWTSecret)
	router.MountLikeFollow(api, lfCtl, cfg.JWTSecret)

	return e
}

func Run(e *echo.Echo, cfg Config) error {
	return e.Start(fmt.Sprintf(":%s", cfg.AppPort))
}
