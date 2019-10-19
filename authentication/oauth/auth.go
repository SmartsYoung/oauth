package oauth

import (
	"errors"
	"github.com/SmartsYoung/authtest/authentication"
	"golang.org/x/oauth2"
)

type Auth struct {

}

func (a *Auth)Authorization(authOptions ...authentication.ConfigOption) (string, error){
	op := authentication.ToOption(authOptions...)
	if op.OpenID() != ""{
		return a.authorizationByOpenID(op.OpenID())
	}

	if op.Account() == "" || op.Secret() == ""{
		return "", errors.New("auth info must be not empty")
	}

	return a.authorizationByAccount(op.Account(), op.Secret())
}

func (a *Auth)Authentication(token string) error {
	return nil
}

func (a *Auth)authorizationByOpenID(openID string) (string, error) {
	return "", nil
}

func (a *Auth)authorizationByAccount(account string, secret string) (string, error) {
	conf := &oauth2.Config{

	}
}