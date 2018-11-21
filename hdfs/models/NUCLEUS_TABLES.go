package models

type NucleusTables struct {
	ClassName     string `xorm:"not null pk VARCHAR(128)"`
	TableName     string `xorm:"not null VARCHAR(128)"`
	Type          string `xorm:"not null VARCHAR(4)"`
	Owner         string `xorm:"not null VARCHAR(2)"`
	Version       string `xorm:"not null VARCHAR(20)"`
	InterfaceName string `xorm:"VARCHAR(255)"`
}
