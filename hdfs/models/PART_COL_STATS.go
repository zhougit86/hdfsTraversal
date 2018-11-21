package models

type PartColStats struct {
	CsId                int64   `xorm:"not null pk BIGINT(20)"`
	DbName              string  `xorm:"not null index(PCS_STATS_IDX) VARCHAR(128)"`
	TableName           string  `xorm:"not null index(PCS_STATS_IDX) VARCHAR(256)"`
	PartitionName       string  `xorm:"not null index(PCS_STATS_IDX) VARCHAR(767)"`
	ColumnName          string  `xorm:"not null index(PCS_STATS_IDX) VARCHAR(767)"`
	ColumnType          string  `xorm:"not null VARCHAR(128)"`
	PartId              int64   `xorm:"not null index BIGINT(20)"`
	LongLowValue        int64   `xorm:"BIGINT(20)"`
	LongHighValue       int64   `xorm:"BIGINT(20)"`
	DoubleHighValue     float64 `xorm:"DOUBLE(53,4)"`
	DoubleLowValue      float64 `xorm:"DOUBLE(53,4)"`
	BigDecimalLowValue  string  `xorm:"VARCHAR(4000)"`
	BigDecimalHighValue string  `xorm:"VARCHAR(4000)"`
	NumNulls            int64   `xorm:"not null BIGINT(20)"`
	NumDistincts        int64   `xorm:"BIGINT(20)"`
	AvgColLen           float64 `xorm:"DOUBLE(53,4)"`
	MaxColLen           int64   `xorm:"BIGINT(20)"`
	NumTrues            int64   `xorm:"BIGINT(20)"`
	NumFalses           int64   `xorm:"BIGINT(20)"`
	LastAnalyzed        int64   `xorm:"not null BIGINT(20)"`
}
