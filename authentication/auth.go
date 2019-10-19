package authentication

type Auth interface {
	Authorization(authOptions ...ConfigOption) (string, error)
	Authentication(token string) error
}

type Option struct{
	account string
	secret string
	openID string
}

func (op *Option)Account() string {
	return op.account
}

func (op *Option)Secret() string {
	return op.account
}

func (op *Option)OpenID() string {
	return op.account
}

type ConfigOption func(option *Option)

func WithAccount(account string) ConfigOption {
	return func(option *Option) {
		option.account = account
	}
}

func WithSecret(secret string) ConfigOption {
	return func(option *Option) {
		option.secret = secret
	}
}

func WithOpenID(openID string) ConfigOption {
	return func(option *Option) {
		option.openID = openID
	}
}

func ToOption(opts ...ConfigOption) *Option {
	op := &Option{}
	for _, opt := range opts {
		opt(op)
	}
	return op
}