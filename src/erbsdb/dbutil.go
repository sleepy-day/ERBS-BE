package erbsdb

type ApiKey struct {
	key string
}

func setApiKey(key string) ApiKey {
	return ApiKey{key: key}
}

func (api *ApiKey) Get() string {
	return api.key
}
