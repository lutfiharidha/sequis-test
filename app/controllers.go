package app

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lutfiharidha/sequis-test/helper"
)

type FriendsController struct {
	model FriendModel
}

func NewFriendsController() *FriendsController {
	return &FriendsController{
		model: NewFriendModels(),
	}
}

func ValidationEmail(email string) bool {
	re, _ := regexp.Compile(`[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`)
	if re.MatchString(strings.ToLower(email)) {
		return true
	}
	return false
}

func (c *FriendsController) Init(m FriendModel) {
	c.model = m //includes a model for creating transaction data in the database
}

// Get implement net http handler
// @Tags Request Friend
// @Summary Request a friend.
// @Accept json
// @Produce json
// @Param data body FriendRequest true "request data"
// @Success 200 {object} FriendResponse
// @Success 400 {object} helper.ResponseError
// @Router /api/v1/friend/request [post]
func (c *FriendsController) RequestFriend(context *gin.Context) {
	var req FriendRequest
	err := context.ShouldBindJSON(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if !ValidationEmail(req.Requestor) || !ValidationEmail(req.To) {
		res := helper.BuildErrorResponse("Failed to process request", "invalid email")
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if req.Requestor == req.To {
		res := helper.BuildErrorResponse("Failed to process request", "Email requestor cannot be same with email to")
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res, err := c.model.FriendRequest(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	context.JSON(http.StatusOK, res)
}

// Get implement net http handler
// @Tags List Friend Request
// @Summary List friend request.
// @Accept json
// @Produce json
// @Param data body ListRequest true "request data"
// @Success 200 {object} ListFriendsRequestResponse
// @Success 400 {object} helper.ResponseError
// @Router /api/v1/friend/list/request [post]
func (c *FriendsController) ListRequest(context *gin.Context) {
	var req ListRequest
	err := context.ShouldBindJSON(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if !ValidationEmail(req.Email) {
		res := helper.BuildErrorResponse("Failed to process request", "invalid email")
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res, err := c.model.ListRequests(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	context.JSON(http.StatusOK, res)
}

// Get implement net http handler
// @Tags Approve
// @Summary Approve friend request.
// @Accept json
// @Produce json
// @Param data body FriendRequest true "request data"
// @Success 200 {object} FriendResponse
// @Success 400 {object} helper.ResponseError
// @Router /api/v1/friend/approve [post]
func (c *FriendsController) Approve(context *gin.Context) {
	var req FriendRequest

	err := context.ShouldBindJSON(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if !ValidationEmail(req.Requestor) || !ValidationEmail(req.To) {
		res := helper.BuildErrorResponse("Failed to process request", "invalid email")
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if req.Requestor == req.To {
		res := helper.BuildErrorResponse("Failed to process request", "Email requestor cannot be same with email to")
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res, err := c.model.Approve(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	context.JSON(http.StatusOK, res)
}

// Get implement net http handler
// @Tags Reject
// @Summary Reject friend request.
// @Accept json
// @Produce json
// @Param data body FriendRequest true "request data"
// @Success 200 {object} FriendResponse
// @Success 400 {object} helper.ResponseError
// @Router /api/v1/friend/reject [post]
func (c *FriendsController) Reject(context *gin.Context) {
	var req FriendRequest

	err := context.ShouldBindJSON(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if !ValidationEmail(req.Requestor) || !ValidationEmail(req.To) {
		res := helper.BuildErrorResponse("Failed to process request", "invalid email")
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if req.Requestor == req.To {
		res := helper.BuildErrorResponse("Failed to process request", "Email requestor cannot be same with email to")
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res, err := c.model.Reject(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	context.JSON(http.StatusOK, res)
}

// Get implement net http handler
// @Tags List Friend
// @Summary List friend.
// @Accept json
// @Produce json
// @Param data body ListRequest true "request data"
// @Success 200 {object} ListResponse
// @Success 400 {object} helper.ResponseError
// @Router /api/v1/friend/list [post]
func (c *FriendsController) ListFriends(context *gin.Context) {
	var req ListRequest
	err := context.ShouldBindJSON(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if !ValidationEmail(req.Email) {
		res := helper.BuildErrorResponse("Failed to process request", "invalid email")
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res, err := c.model.ListFriends(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	context.JSON(http.StatusOK, res)
}

// Get implement net http handler
// @Tags List Common Friend
// @Summary List common friend between users.
// @Accept json
// @Produce json
// @Param data body CommonRequest true "request data"
// @Success 200 {object} CommonResponse
// @Success 400 {object} helper.ResponseError
// @Router /api/v1/friend/list/common [post]
func (c *FriendsController) ListCommonFriends(context *gin.Context) {
	var req CommonRequest
	err := context.ShouldBindJSON(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}
	for _, v := range req.Friends {
		if !ValidationEmail(v) {
			res := helper.BuildErrorResponse("Failed to process request", "please check the email list")
			context.JSON(http.StatusBadRequest, res)
			return
		}
	}
	res, err := c.model.ListCommonFriends(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	context.JSON(http.StatusOK, res)
}

// Get implement net http handler
// @Tags Block
// @Summary Block a friend.
// @Accept json
// @Produce json
// @Param data body FriendRequest true "request data"
// @Success 200 {object} FriendResponse
// @Success 400 {object} helper.ResponseError
// @Router /api/v1/friend/block [post]
func (c *FriendsController) Block(context *gin.Context) {
	var req BlockRequest
	err := context.ShouldBindJSON(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}

	if !ValidationEmail(req.Requestor) || !ValidationEmail(req.Block) {
		res := helper.BuildErrorResponse("Failed to process request", "invalid email")
		context.JSON(http.StatusBadRequest, res)
		return
	}

	if req.Requestor == req.Block {
		res := helper.BuildErrorResponse("Failed to process request", "Email requestor cannot be same with email to")
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res, err := c.model.Block(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error())
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	context.JSON(http.StatusOK, res)
}
