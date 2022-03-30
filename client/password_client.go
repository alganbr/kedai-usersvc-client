package client

import (
	"encoding/json"
	"github.com/alganbr/kedai-usersvc-client/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"net/http"
)

type IPasswordClient interface {
	Update(*models.UserPasswordRq) *errors.Error
	Validate(*models.ValidatePasswordRq) (*models.User, *errors.Error)
}

type PasswordClient struct {
	httpClient *rest.RequestBuilder
}

func (c *PasswordClient) Update(rq *models.UserPasswordRq) *errors.Error {
	rs := c.httpClient.Put("/usersvc/password", rq)
	if rs.Err != nil {
		return &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: rs.Err.Error(),
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return httpErr
	}
	return nil
}

func (c *PasswordClient) Validate(rq *models.ValidatePasswordRq) (*models.User, *errors.Error) {
	rs := c.httpClient.Post("/usersvc/password/validate", rq)
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
