package deals

import (
	"bytes"
	"encoding/xml"
	atas "moex/AbstractTableStruct"
	"moex/TradesHistoryTable"
	"moex/utils"
)

const rowElementName = "row"
const countRecordsInBasket = 100_000

var ArrayRows atas.Rows
var baskets atas.Baskets

func GetDeals() int {

	xmlData := bytes.NewBufferString(utils.GetFile("C:\\\\Users\\\\Max\\\\Downloads\\\\trades-moex-stock-shares-2023-08-18xml\\trades.xml"))

	d := xml.NewDecoder(xmlData)
	var tRec atas.AbstractRecord
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
					ArrayRows = atas.Rows{}
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
