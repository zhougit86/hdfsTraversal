package models

type DelegationTokens struct {
	TokenIdent string `xorm:"not null pk VARCHAR(767)"`
	Token      string `xorm:"VARCHAR(767)"`
}
