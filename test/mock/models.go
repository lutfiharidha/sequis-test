package mock

import (
	"github.com/lutfiharidha/sequis-test/app"
	"github.com/stretchr/testify/mock"
)

type FriendModel struct {
	mock.Mock
}

func (c *FriendModel) Init() {}

func (c *FriendModel) FriendRequest(req app.FriendRequest) (res app.FriendResponse, err error) {
	res = app.FriendResponse{}

	call := c.Called()
	re := call.Get(0)
	er := call.Get(1)

	if er != nil {
		ree, ok := er.(error)
		if ok {
			return res, ree
		}
	}

	ree, ok := re.(app.FriendResponse)
	if ok {
		return ree, nil
	}

	return res, err
}

func (c *FriendModel) Approve(req app.FriendRequest) (res app.FriendResponse, err error) {

	res = app.FriendResponse{}

	call := c.Called()
	re := call.Get(0)
	er := call.Get(1)

	if er != nil {
		ree, ok := er.(error)
		if ok {
			return res, ree
		}
	}

	ree, ok := re.(app.FriendResponse)
	if ok {
		return ree, nil
	}

	return res, err
}

func (c *FriendModel) Reject(req app.FriendRequest) (res app.FriendResponse, err error) {

	res = app.FriendResponse{}

	call := c.Called()
	re := call.Get(0)
	er := call.Get(1)

	if er != nil {
		ree, ok := er.(error)
		if ok {
			return res, ree
		}
	}

	ree, ok := re.(app.FriendResponse)
	if ok {
		return ree, nil
	}

	return res, err
}

func (c *FriendModel) ListRequests(req app.ListRequest) (res app.ListFriendsRequestResponse, err error) {
	res = app.ListFriendsRequestResponse{}
	call := c.Called()
	re := call.Get(0)
	er := call.Get(1)

	if er != nil {
		ree, ok := er.(error)
		if ok {
			return res, ree
		}
	}

	ree, ok := re.(app.ListFriendsRequestResponse)
	if ok {
		return ree, nil
	}

	return res, err
}

func (c *FriendModel) ListFriends(req app.ListRequest) (res app.ListResponse, err error) {

	res = app.ListResponse{}

	call := c.Called()
	re := call.Get(0)
	er := call.Get(1)

	if er != nil {
		ree, ok := er.(error)
		if ok {
			return res, ree
		}
	}

	ree, ok := re.(app.ListResponse)
	if ok {
		return ree, nil
	}

	return res, err
}

func (c *FriendModel) ListCommonFriends(req app.CommonRequest) (res app.CommonResponse, err error) {

	res = app.CommonResponse{}

	call := c.Called()
	re := call.Get(0)
	er := call.Get(1)

	if er != nil {
		ree, ok := er.(error)
		if ok {
			return res, ree
		}
	}

	ree, ok := re.(app.CommonResponse)
	if ok {
		return ree, nil
	}

	return res, err
}

func (c *FriendModel) Block(req app.BlockRequest) (res app.FriendResponse, err error) {

	res = app.FriendResponse{}

	call := c.Called()
	re := call.Get(0)
	er := call.Get(1)

	if er != nil {
		ree, ok := er.(error)
		if ok {
			return res, ree
		}
	}

	ree, ok := re.(app.FriendResponse)
	if ok {
		return ree, nil
	}

	return res, err
}
