package app_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lutfiharidha/sequis-test/app"
	"github.com/lutfiharidha/sequis-test/helper"
	mockmodel "github.com/lutfiharidha/sequis-test/test/mock"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestRequestFriend(t *testing.T) {
	convey.Convey("TestRequestFriend", t, func() {
		convey.Convey("Positive Scenario", func() {
			mockResponse := app.FriendResponse{Success: true}
			jsonMockRes, _ := json.Marshal(mockResponse)

			mc := &mockmodel.FriendModel{}
			ctrl := app.NewFriendsController()
			ctrl.Init(mc)
			mc.On("FriendRequest", mock.Anything, mock.Anything).Return(app.FriendResponse{
				Success: true,
			}, nil)

			r := SetUpRouter()
			r.POST("/api/v1/friend/request", ctrl.RequestFriend)

			request := app.FriendRequest{
				Requestor: "user1@mail.com",
				To:        "user2@mail.com",
			}
			jsonValue, _ := json.Marshal(request)
			req, _ := http.NewRequest("POST", "/api/v1/friend/request", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res, _ := ioutil.ReadAll(w.Body)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
		})

		convey.Convey("Negative Scenario", func() {
			convey.Convey("Error DB", func() {
				mockResponse := app.FriendResponse{Success: true}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("FriendRequest", mock.Anything, mock.Anything).Return(nil, errors.New("Failed to process request"))

				r := SetUpRouter()
				r.POST("/api/v1/friend/request", ctrl.RequestFriend)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
					To:        "user2@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/request", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusInternalServerError)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

			convey.Convey("Incomplete request", func() {
				mockResponse := app.FriendResponse{Success: true}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("FriendRequest", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/request", ctrl.RequestFriend)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/request", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

			convey.Convey("Invalid email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "invalid email",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("FriendRequest", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/request", ctrl.RequestFriend)

				request := app.FriendRequest{
					Requestor: "user1.com",
					To:        "user1.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/request", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Requestor and To have same email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "Email requestor cannot be same with email to",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("FriendRequest", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/request", ctrl.RequestFriend)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
					To:        "user1@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/request", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Already give a response to friend request", func() {
				mockResponse := app.FriendResponse{Success: false}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("FriendRequest", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/request", ctrl.RequestFriend)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
					To:        "user2@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/request", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})
		})
	})
}

func TestListRequests(t *testing.T) {
	convey.Convey("TestListRequests", t, func() {
		convey.Convey("Positive Scenario", func() {
			mockResponse := app.ListFriendsRequestResponse{
				Request: []app.FriendsRequestResponse{
					{
						Requestor: "user1@gmail.com",
						Status:    "pending",
					},
				},
			}

			mc := &mockmodel.FriendModel{}
			ctrl := app.NewFriendsController()
			ctrl.Init(mc)
			mc.On("ListRequests", mock.Anything, mock.Anything).Return(mockResponse, nil)

			r := SetUpRouter()
			r.POST("/api/v1/friend/list/request", ctrl.ListRequest)

			request := app.ListRequest{
				Email: "user2@mail.com",
			}
			jsonValue, _ := json.Marshal(request)
			req, _ := http.NewRequest("POST", "/api/v1/friend/list/request", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res, _ := ioutil.ReadAll(w.Body)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			convey.So(string(res), convey.ShouldNotBeEmpty)
		})

		convey.Convey("Negative Scenario", func() {
			convey.Convey("Error in Model", func() {
				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("ListRequests", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("Failed to process request"))

				r := SetUpRouter()
				r.POST("/api/v1/friend/list/request", ctrl.ListRequest)

				request := app.ListRequest{
					Email: "user1@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/list/request", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				fmt.Println("r", string(res))
				convey.So(w.Code, convey.ShouldEqual, http.StatusInternalServerError)
				convey.So(string(res), convey.ShouldNotBeNil)
			})

			convey.Convey("Invalid email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "invalid email",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("ListRequests", mock.Anything, mock.Anything).Return(app.ListFriendsRequestResponse{}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/list/request", ctrl.ListRequest)

				request := app.ListRequest{
					Email: "user1.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/list/request", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Incomplete request", func() {
				mockResponse := app.ListFriendsRequestResponse{}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("ListRequests", mock.Anything, mock.Anything).Return(mockResponse, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/list/request", ctrl.ListRequest)

				request := app.FriendRequest{}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/list/request", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

		})
	})
}

func TestApprove(t *testing.T) {
	convey.Convey("TestApprove", t, func() {
		convey.Convey("Positive Scenario", func() {
			mockResponse := app.FriendResponse{Success: true}
			jsonMockRes, _ := json.Marshal(mockResponse)

			mc := &mockmodel.FriendModel{}
			ctrl := app.NewFriendsController()
			ctrl.Init(mc)
			mc.On("Approve", mock.Anything, mock.Anything).Return(app.FriendResponse{
				Success: true,
			}, nil)

			r := SetUpRouter()
			r.POST("/api/v1/friend/approve", ctrl.Approve)

			request := app.FriendRequest{
				Requestor: "user1@mail.com",
				To:        "user2@mail.com",
			}
			jsonValue, _ := json.Marshal(request)
			req, _ := http.NewRequest("POST", "/api/v1/friend/approve", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res, _ := ioutil.ReadAll(w.Body)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
		})

		convey.Convey("Negative Scenario", func() {
			convey.Convey("Error DB", func() {
				mockResponse := app.FriendResponse{Success: true}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Approve", mock.Anything, mock.Anything).Return(nil, errors.New("Failed to process request"))

				r := SetUpRouter()
				r.POST("/api/v1/friend/approve", ctrl.Approve)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
					To:        "user2@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/approve", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusInternalServerError)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

			convey.Convey("Incomplete request", func() {
				mockResponse := app.FriendResponse{Success: true}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Approve", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/approve", ctrl.Approve)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/approve", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

			convey.Convey("Invalid email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "invalid email",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Approve", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/approve", ctrl.Approve)

				request := app.FriendRequest{
					Requestor: "user1.com",
					To:        "user1.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/approve", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Requestor and To have same email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "Email requestor cannot be same with email to",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Approve", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/approve", ctrl.Approve)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
					To:        "user1@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/approve", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Already give a response to friend request", func() {
				mockResponse := app.FriendResponse{Success: false}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Approve", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/approve", ctrl.Approve)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
					To:        "user2@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/approve", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})
		})
	})
}

func TestReject(t *testing.T) {
	convey.Convey("TestReject", t, func() {
		convey.Convey("Positive Scenario", func() {
			mockResponse := app.FriendResponse{Success: true}
			jsonMockRes, _ := json.Marshal(mockResponse)

			mc := &mockmodel.FriendModel{}
			ctrl := app.NewFriendsController()
			ctrl.Init(mc)
			mc.On("Reject", mock.Anything, mock.Anything).Return(app.FriendResponse{
				Success: true,
			}, nil)

			r := SetUpRouter()
			r.POST("/api/v1/friend/reject", ctrl.Reject)

			request := app.FriendRequest{
				Requestor: "user1@mail.com",
				To:        "user2@mail.com",
			}
			jsonValue, _ := json.Marshal(request)
			req, _ := http.NewRequest("POST", "/api/v1/friend/reject", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res, _ := ioutil.ReadAll(w.Body)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
		})

		convey.Convey("Negative Scenario", func() {
			convey.Convey("Error DB", func() {
				mockResponse := app.FriendResponse{Success: true}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Reject", mock.Anything, mock.Anything).Return(nil, errors.New("Failed to process request"))

				r := SetUpRouter()
				r.POST("/api/v1/friend/reject", ctrl.Reject)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
					To:        "user2@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/reject", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusInternalServerError)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

			convey.Convey("Incomplete request", func() {
				mockResponse := app.FriendResponse{Success: true}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Reject", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/reject", ctrl.Reject)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/reject", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

			convey.Convey("Invalid email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "invalid email",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Reject", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/reject", ctrl.Reject)

				request := app.FriendRequest{
					Requestor: "user1.com",
					To:        "user1.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/reject", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Requestor and To have same email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "Email requestor cannot be same with email to",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Reject", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/reject", ctrl.Reject)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
					To:        "user1@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/reject", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Already give a response to friend request", func() {
				mockResponse := app.FriendResponse{Success: false}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Reject", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/reject", ctrl.Reject)

				request := app.FriendRequest{
					Requestor: "user1@mail.com",
					To:        "user2@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/reject", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})
		})
	})
}

func TestListFriends(t *testing.T) {
	convey.Convey("TestListFriends", t, func() {
		convey.Convey("Positive Scenario", func() {
			mockResponse := app.ListResponse{
				Friends: []string{
					"user1@gmail.com",
				},
			}

			mc := &mockmodel.FriendModel{}
			ctrl := app.NewFriendsController()
			ctrl.Init(mc)
			mc.On("ListFriends", mock.Anything, mock.Anything).Return(mockResponse, nil)

			r := SetUpRouter()
			r.POST("/api/v1/friend/list", ctrl.ListFriends)

			request := app.ListRequest{
				Email: "user2@mail.com",
			}
			jsonValue, _ := json.Marshal(request)
			req, _ := http.NewRequest("POST", "/api/v1/friend/list", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res, _ := ioutil.ReadAll(w.Body)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			convey.So(string(res), convey.ShouldNotBeEmpty)
		})

		convey.Convey("Negative Scenario", func() {
			convey.Convey("Error in Model", func() {
				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("ListFriends", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("Failed to process request"))

				r := SetUpRouter()
				r.POST("/api/v1/friend/list", ctrl.ListFriends)

				request := app.ListRequest{
					Email: "user1@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/list", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				fmt.Println("r", string(res))
				convey.So(w.Code, convey.ShouldEqual, http.StatusInternalServerError)
				convey.So(string(res), convey.ShouldNotBeNil)
			})

			convey.Convey("Invalid email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "invalid email",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("ListFriends", mock.Anything, mock.Anything).Return(app.ListResponse{}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/list", ctrl.ListFriends)

				request := app.ListRequest{
					Email: "user1.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/list", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Incomplete request", func() {
				mockResponse := app.ListResponse{}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("ListFriends", mock.Anything, mock.Anything).Return(mockResponse, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/list", ctrl.ListFriends)

				request := app.FriendRequest{}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/list", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

		})
	})
}

func TestListCommonFriends(t *testing.T) {
	convey.Convey("TestListCommonFriends", t, func() {
		convey.Convey("Positive Scenario", func() {
			mockResponse := app.CommonResponse{
				Success: true,
				Friends: app.ListResponse{
					Friends: []string{
						"user2@example.com",
					},
				},
			}

			mc := &mockmodel.FriendModel{}
			ctrl := app.NewFriendsController()
			ctrl.Init(mc)
			mc.On("ListCommonFriends", mock.Anything, mock.Anything).Return(mockResponse, nil)

			r := SetUpRouter()
			r.POST("/api/v1/friend/list/common", ctrl.ListCommonFriends)

			request := app.CommonRequest{
				Friends: []string{
					"user1@example.com",
					"user3@example.com",
				},
			}
			jsonValue, _ := json.Marshal(request)
			req, _ := http.NewRequest("POST", "/api/v1/friend/list/common", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res, _ := ioutil.ReadAll(w.Body)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			convey.So(string(res), convey.ShouldNotBeEmpty)
		})

		convey.Convey("Negative Scenario", func() {
			convey.Convey("Error in Model", func() {
				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("ListCommonFriends", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("Failed to process request"))

				r := SetUpRouter()
				r.POST("/api/v1/friend/list/common", ctrl.ListCommonFriends)

				request := app.CommonRequest{
					Friends: []string{
						"user1@example.com",
						"user3@example.com",
					},
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/list/common", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusInternalServerError)
				convey.So(string(res), convey.ShouldNotBeNil)
			})

			convey.Convey("Invalid email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "please check the email list",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("ListCommonFriends", mock.Anything, mock.Anything).Return(app.CommonResponse{}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/list/common", ctrl.ListCommonFriends)

				request := app.CommonRequest{
					Friends: []string{
						"user.com",
					},
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/list/common", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Incomplete request", func() {
				mockResponse := app.ListResponse{}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("ListCommonFriends", mock.Anything, mock.Anything).Return(mockResponse, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/list/common", ctrl.ListCommonFriends)

				request := app.CommonRequest{}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/list/common", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

		})
	})
}

func TestBlock(t *testing.T) {
	convey.Convey("TestBlock", t, func() {
		convey.Convey("Positive Scenario", func() {
			mockResponse := app.FriendResponse{Success: true}
			jsonMockRes, _ := json.Marshal(mockResponse)

			mc := &mockmodel.FriendModel{}
			ctrl := app.NewFriendsController()
			ctrl.Init(mc)
			mc.On("Block", mock.Anything, mock.Anything).Return(app.FriendResponse{
				Success: true,
			}, nil)

			r := SetUpRouter()
			r.POST("/api/v1/friend/block", ctrl.Block)

			request := app.BlockRequest{
				Requestor: "user1@mail.com",
				Block:     "user2@mail.com",
			}
			jsonValue, _ := json.Marshal(request)
			req, _ := http.NewRequest("POST", "/api/v1/friend/block", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res, _ := ioutil.ReadAll(w.Body)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
		})

		convey.Convey("Negative Scenario", func() {
			convey.Convey("Error DB", func() {
				mockResponse := app.FriendResponse{Success: true}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Block", mock.Anything, mock.Anything).Return(nil, errors.New("Failed to process request"))

				r := SetUpRouter()
				r.POST("/api/v1/friend/block", ctrl.Block)

				request := app.BlockRequest{
					Requestor: "user1@mail.com",
					Block:     "user2@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/block", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusInternalServerError)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

			convey.Convey("Incomplete request", func() {
				mockResponse := app.FriendResponse{Success: true}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Block", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/block", ctrl.Block)

				request := app.BlockRequest{
					Requestor: "user1@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/block", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

			convey.Convey("Invalid email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "invalid email",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Block", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/block", ctrl.Block)

				request := app.BlockRequest{
					Requestor: "user1.com",
					Block:     "user1.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/block", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Requestor and To have same email", func() {
				mockResponse := helper.ResponseError{
					Success: false,
					Message: "Failed to process request",
					Errors:  "Email requestor cannot be same with email to",
				}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Block", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/block", ctrl.Block)

				request := app.BlockRequest{
					Requestor: "user1@mail.com",
					Block:     "user1@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/block", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Already block", func() {
				mockResponse := app.FriendResponse{Success: false}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.FriendModel{}
				ctrl := app.NewFriendsController()
				ctrl.Init(mc)
				mc.On("Block", mock.Anything, mock.Anything).Return(app.FriendResponse{
					Success: false,
				}, nil)

				r := SetUpRouter()
				r.POST("/api/v1/friend/block", ctrl.Block)

				request := app.BlockRequest{
					Requestor: "user1@mail.com",
					Block:     "user2@mail.com",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/friend/block", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})
		})
	})
}
