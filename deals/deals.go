package deals

import (
	"bytes"
	"encoding/xml"
	"moex/AbstractTableStruct"
	"moex/TradesHistoryTable"
	"moex/utils"
	"strconv"
	"strings"
	"time"
)

const rowElementName = "row"
const countRecordsInBasket = 100_000

type Rows []AbstractTableStruct.AbstractRecord

type Baskets []Rows

/*
func DealRecordToStr(rec *TradesHistoryRecord) string {
		var rc *TradesHistoryRecord
		var RecStr string
		rc = rec
		return strconv.FormatInt(rc.TRADENO, 10)
}
*/

var ArrayRows Rows
var baskets Baskets

func GetDeals() int {

	xmlData := bytes.NewBufferString(utils.GetFile("C:\\\\Users\\\\Max\\\\Downloads\\\\trades-moex-stock-shares-2023-08-18xml\\trades.xml"))

	d := xml.NewDecoder(xmlData)
	var tRec AbstractTableStruct.AbstractRecord
	countAppendedRecords := 0
	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == rowElementName {
				tRec = &TradesHistoryTable.TradesHistoryRecord{}
				tRec.GetRecordFromXml(&se.Attr)
				countAppendedRecords++
				if countAppendedRecords <= countRecordsInBasket {
					ArrayRows = append(ArrayRows, tRec)
				} else {
					baskets = append(baskets, ArrayRows)
					ArrayRows = Rows{}
					countAppendedRecords = 0
				}
			}
		}
	}
	if countAppendedRecords > 0 {
		baskets = append(baskets, ArrayRows)
	}
	return len(baskets)
}
