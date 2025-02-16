package TradesHistoryTable

import (
	atas "moex/AbstractTableStruct"
	"moex/utils"
)

type TradesHistoryTable struct {
	i int
}

func (tbl TradesHistoryTable) SetXmlSourse() string {
	return utils.GetFile("C:\\\\Users\\\\Max\\\\Downloads\\\\trades-moex-stock-shares-2023-08-18xml\\trades.xml")
}

func (tbl TradesHistoryTable) GetCountRecordsInBasket() int {
	return 100_000
}

func (tbl TradesHistoryTable) GetBlockRowsDelimeters() string {
	st := "rows"
	return st
}

func (tbl TradesHistoryTable) GetTargetRecord() atas.AbstractRecord {
	return &TradesHistoryRecord{}
}

var mTable atas.Table

func Get() {

	mTable.GetRecordsFromZip(&TradesHistoryTable{})
}
