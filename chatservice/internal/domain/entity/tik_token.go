package entity

//go:generate mockgen -source=tik_token.go -destination=mocks/tik_token_go_mock.go -package=mocks entity TikToken

type TikToken interface {
	CountTokens(name, content string) int
}

type TikTokenImpl struct{}

func NewTikTokenImpl() *TikTokenImpl {
	return &TikTokenImpl{}
}
