package deals

import (
	"bytes"
	"encoding/xml"
	inter "moex/interfases"
	"moex/utils"
	"strconv"
	"strings"
	"time"
)

const rowElementName = "row"
const countRecordsInBasket = 100_000

type Rows []inter.TableRecord

type Baskets []Rows

type Row struct {
	TRADENO        int64
	TRADEDATE      time.Time
	TRADETIME      time.Time
	SECID          string
	BOARDID        string
	PRICE          float64
	QUANTITY       int64
	VALUE          float64
	TYPE           string
	BUYSELL        string
	TRADINGSESSION int32
}

/*
func DealRecordToStr(rec *Row) string {
		var rc *Row
		var RecStr string
		rc = rec
		return strconv.FormatInt(rc.TRADENO, 10)
}
*/

var ArrayRows Rows
var baskets Baskets

func (row *Row) GetRecordFromXml(attrs *[]xml.Attr) {

	var tradeDate string

	var layout = "2006-01-02"
	var layoutTime = "2006-01-02 15:04:05"
	var Value string

	for idx := range len(*attrs) {

		Value = (*attrs)[idx].Value

		switch (*attrs)[idx].Name.Local {
		case "TRADENO":
			row.TRADENO, _ = strconv.ParseInt(Value, 10, 64)
		case "TRADEDATE":
			row.TRADEDATE, _ = time.Parse(layout, Value)
			tradeDate = Value
		case "TRADETIME":
			tradeTime := tradeDate + " " + strings.Trim(Value, " ")
			row.TRADETIME, _ = time.Parse(layoutTime, tradeTime)
		case "SECID":
			row.SECID = Value
		case "BOARDID":
			row.BOARDID = Value
		case "PRICE":
			row.PRICE, _ = strconv.ParseFloat(Value, 64)
		case "QUANTITY":
			row.QUANTITY, _ = strconv.ParseInt(Value, 10, 64)
		case "VALUE":
			row.VALUE, _ = strconv.ParseFloat(Value, 64)
		case "TYPE":
			row.TYPE = Value
		case "BUYSELL":
			row.BUYSELL = Value
		case "TRADINGSESSION":
			t64, _ := strconv.ParseInt(Value, 10, 64)
			row.TRADINGSESSION = int32(t64)
		}
	}
}

func GetDeals() int {

	xmlData := bytes.NewBufferString(utils.GetFile("C:\\\\Users\\\\Max\\\\Downloads\\\\trades-moex-stock-shares-2023-08-18xml\\trades.xml"))

	d := xml.NewDecoder(xmlData)
	var tRec inter.TableRecord
	countAppendedRecords := 0
	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == rowElementName {
				tRec = &Row{}
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
