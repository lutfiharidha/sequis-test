package app

import (
	"github.com/gin-gonic/gin"
)

func FriendsRoute(route *gin.Engine) {
	m := NewFriendModels()
	m.Init()
	c := NewFriendsController()
	c.Init(m)
	friendRoutes := route.Group("api/v1/friend")
	{
		friendRoutes.POST("/request", c.RequestFriend)
		friendRoutes.POST("/list/request", c.ListRequest)
		friendRoutes.POST("/approve", c.Approve)
		friendRoutes.POST("/reject", c.Reject)
		friendRoutes.POST("/list", c.ListFriends)
		friendRoutes.POST("/list/common", c.ListCommonFriends)
		friendRoutes.POST("/block", c.Block)
	}
}

func Router() {
	r := gin.New()
	r.Use(CORS())   //handle cors
	FriendsRoute(r) //call registered route

	r.Run(":8081")

}
