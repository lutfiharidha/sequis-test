package app

import (
	"fmt"

	"gorm.io/gorm"
)

type FriendModel interface {
	Init()
	Approve(req FriendRequest) (res FriendResponse, err error)
	Reject(req FriendRequest) (res FriendResponse, err error)
	FriendRequest(req FriendRequest) (res FriendResponse, err error)
	ListRequests(req ListRequest) (res ListFriendsRequestResponse, err error)
	ListFriends(req ListRequest) (res ListResponse, err error)
	ListCommonFriends(req CommonRequest) (res CommonResponse, err error)
	Block(req BlockRequest) (res FriendResponse, err error)
}

type friendModel struct {
	db *gorm.DB
}

func NewFriendModels() FriendModel {
	return &friendModel{}
}

func (m *friendModel) Init() {
	db := NewSQL().SetupDatabaseConnection()
	m.db = db //calling database connection
}

func (m *friendModel) FriendRequest(req FriendRequest) (res FriendResponse, err error) {
	reqUser := []string{
		req.Requestor,
		req.To,
	}
	tx1 := m.db.Debug().Where("requestor in ?", reqUser).Where("friends.to in ?", reqUser).Where("status", "pending").Find(&Friend{})
	if tx1.Error != nil {
		return res, tx1.Error
	}
	if tx1.RowsAffected == 0 { //check if request beetwen user still pending
		tx := m.db.Debug().Create(&Friend{
			Requestor: req.Requestor,
			To:        req.To,
			Status:    "pending",
		})
		if tx.Error != nil {
			return res, tx.Error
		}
		res.Success = true
		return res, err
	}
	res.Success = false
	return res, err
}

func (m *friendModel) Approve(req FriendRequest) (res FriendResponse, err error) {
	reqUser := []string{
		req.Requestor,
		req.To,
	}
	tx1 := m.db.Debug().Where("requestor in ?", reqUser).Where("friends.to in ?", reqUser).Where("status", "pending").Find(&Friend{})
	if tx1.RowsAffected > 0 { //check if request beetwen user still pending
		tx := m.db.Debug().Model(&Friend{}).Where("requestor in ?", reqUser).Where("friends.to in ?", reqUser).Update("status", "accepted")
		if tx.Error != nil {
			return res, tx.Error
		}
		res.Success = true
		return res, err
	}

	res.Success = false
	return res, err
}

func (m *friendModel) Reject(req FriendRequest) (res FriendResponse, err error) {
	reqUser := []string{
		req.Requestor,
		req.To,
	}
	tx1 := m.db.Debug().Where("requestor in ?", reqUser).Where("friends.to in ?", reqUser).Where("status", "pending").Find(&Friend{})
	if tx1.RowsAffected > 0 { //check if request beetwen user still pending
		tx := m.db.Debug().Model(&Friend{}).Where("requestor in ?", reqUser).Where("friends.to in ?", reqUser).Update("status", "rejected")
		if tx.Error != nil {
			return res, tx.Error
		}
		res.Success = true
		return res, err
	}

	res.Success = false
	return res, err
}

func (m *friendModel) ListRequests(req ListRequest) (res ListFriendsRequestResponse, err error) {
	var listReq []Friend
	tx := m.db.Debug().Where("to", req.Email).Find(&listReq)
	if tx.Error != nil {
		return res, tx.Error
	}
	for _, v := range listReq {
		res.Request = append(res.Request, FriendsRequestResponse{
			Requestor: v.Requestor,
			Status:    v.Status,
		})
	}
	return res, err
}

func (m *friendModel) ListFriends(req ListRequest) (res ListResponse, err error) {
	var listReq []Friend
	tx := m.db.Debug().Raw(`SELECT * FROM friends WHERE (requestor = ? OR friends.to = ?) AND status =  'accepted'`, req.Email, req.Email).Scan(&listReq)
	if tx.Error != nil {
		return res, tx.Error
	}
	var before []string
	for _, v := range listReq {
		before = append(before, v.Requestor)
		before = append(before, v.To)
	}

	for _, v := range before {
		if v != req.Email {
			res.Friends = append(res.Friends, v)
		}
	}
	return res, err
}

func (m *friendModel) ListCommonFriends(req CommonRequest) (res CommonResponse, err error) {
	var listReq []Friend
	fmt.Println("req", req)
	tx := m.db.Debug().Raw(`SELECT * from friends where (friends.to in ? OR requestor in ? ) AND status = 'accepted'`, req.Friends, req.Friends).Scan(&listReq)
	if tx.Error != nil {
		return res, tx.Error
	}
	var before []string
	for _, v := range listReq {
		for _, b := range req.Friends {
			if b != v.Requestor && b != v.To {
				if b == v.Requestor {
					before = append(before, v.Requestor)
				} else {
					before = append(before, v.To)
				}
			}
		}
	}

	keys := make(map[string]bool)
	list := []string{}

	for _, v := range before {
		if _, value := keys[v]; !value {
			keys[v] = true
			list = append(list, v)
		}
	}
	res.Friends = ListResponse{
		Friends: list,
	}
	res.Success = true
	res.Count = len(list)
	return res, err
}

func (m *friendModel) Block(req BlockRequest) (res FriendResponse, err error) {
	tx := m.db.Debug().Model(&Friend{}).Where("requestor", req.Requestor).Where("to", req.Block).Update("status", "blocked")
	if tx.Error != nil {
		return res, tx.Error
	}
	if tx.RowsAffected == 0 {
		res.Success = false
		return res, err
	}
	res.Success = true
	return res, err
}
