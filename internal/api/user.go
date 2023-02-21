package api

import (
	"net/http"

	"github.com/Banana-Boat/gin-sqlc-template/internal/db"
	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Username string         `json:"username" binding:"required"`
	Password string         `json:"password" binding:"required"`
	Gender   db.UsersGender `json:"gender" binding:"required"`
	Age      int32          `json:"age" binding:"required"`
}

func (server *Server) register(ctx *gin.Context) {
	var req registerRequest

	// 通过gin的binding校验参数合法性
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// 创建用户
	arg := db.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
		Gender:   req.Gender,
		Age:      req.Age,
	}
	res, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	// 查询新增用户
	id, _ := res.LastInsertId()
	user, err := server.store.GetUser(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	// 返回结果
	ctx.JSON(http.StatusOK, user)

}
