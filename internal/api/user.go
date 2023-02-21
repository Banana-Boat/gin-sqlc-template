package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Banana-Boat/gin-sqlc-template/internal/db"
	"github.com/Banana-Boat/gin-sqlc-template/internal/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string         `json:"username" binding:"required"`
	Password string         `json:"password" binding:"required"`
	Gender   db.UsersGender `json:"gender" binding:"required"`
	Age      int32          `json:"age" binding:"required"`
}

type createUserResponse struct {
	ID        int32          `json:"id"`
	Username  string         `json:"username"`
	Gender    db.UsersGender `json:"gender"`
	Age       int32          `json:"age"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	/* 通过gin的binding校验参数合法性 */
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	/* 创建用户 */
	hashedPassword, err := util.HashPassword(req.Password) // 对密码加密
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Gender:   req.Gender,
		Age:      req.Age,
	}
	res, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	/* 查询新增用户 */
	id, _ := res.LastInsertId()
	user, err := server.store.GetUser(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	resp := createUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Age:       user.Age,
		Gender:    user.Gender,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	/* 返回结果 */
	ctx.JSON(http.StatusOK, resp)
}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type listUserRequest struct {
	PageIdx  int32 `form:"pageIdx" binding:"min=0"`
	PageSize int32 `form:"pageSize" binding:"required,min=5,max=10"`
}

func (server *Server) listUsers(ctx *gin.Context) {

	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: req.PageIdx * req.PageSize,
	}

	accounts, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
