package structs

import "context"

type TokenAuth struct {
	Token string
}

func (t TokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + t.Token,
	}, nil
}

func (TokenAuth) RequireTransportSecurity() bool {
	return true
}
