package client

import "github.com/mercadolibre/golang-restclient/rest"

type IUserSvcClient interface {
	User() IUserClient
	Password() IPasswordClient
}

type UserSvcClient struct {
	HttpClient *rest.RequestBuilder
}

func (c UserSvcClient) User() IUserClient {
	return &UserClient{c.HttpClient}
}

func (c UserSvcClient) Password() IPasswordClient {
	return &PasswordClient{c.HttpClient}
}
