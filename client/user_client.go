package client

import (
	"encoding/json"
	"fmt"
	"github.com/alganbr/kedai-usersvc-client/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"net/http"
)

type IUserClient interface {
	Get(int64) (*models.User, *errors.Error)
	Create(*models.UserRq) (*models.User, *errors.Error)
	Update(int64, *models.UserRq) (*models.User, *errors.Error)
	Patch(int64, *models.UserRq) (*models.User, *errors.Error)
}

type UserClient struct {
	httpClient *rest.RequestBuilder
}

func (c *UserClient) Get(id int64) (*models.User, *errors.Error) {
	rs := c.httpClient.Get(fmt.Sprintf("/usersvc/user/%d", id))
	if rs.Err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: rs.Err.Error(),
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return nil, httpErr
	}
	var user *models.User
	if err := json.Unmarshal(rs.Bytes(), &user); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error when unmarshalling response body",
		}
	}
	return user, nil
}

func (c *UserClient) Create(rq *models.UserRq) (*models.User, *errors.Error) {
	rs := c.httpClient.Post("/usersvc/user", rq)
	if rs.Response != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: rs.Err.Error(),
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return nil, httpErr
	}
	var user *models.User
	if err := json.Unmarshal(rs.Bytes(), &user); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error when unmarshalling response body",
		}
	}
	return user, nil
}

func (c *UserClient) Update(id int64, rq *models.UserRq) (*models.User, *errors.Error) {
	rs := c.httpClient.Put(fmt.Sprintf("/usersvc/user/%d", id), rq)
	if rs == nil || rs.Response == nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error getting http response",
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return nil, httpErr
	}
	var user *models.User
	if err := json.Unmarshal(rs.Bytes(), &user); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error when unmarshalling response body",
		}
	}
	return user, nil
}

func (c *UserClient) Patch(id int64, rq *models.UserRq) (*models.User, *errors.Error) {
	rs := c.httpClient.Patch(fmt.Sprintf("/usersvc/user/%d", id), rq)
	if rs.Err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: rs.Err.Error(),
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return nil, httpErr
	}
	var user *models.User
	if err := json.Unmarshal(rs.Bytes(), &user); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error when unmarshalling response body",
		}
	}
	return user, nil
}
