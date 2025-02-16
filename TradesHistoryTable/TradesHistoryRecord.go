package TradesHistoryTable

import (
	"encoding/xml"
	"moex/AbstractTableStruct"
	"strconv"
	"strings"
	"time"
)

type TradesHistoryRecord struct {
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

func (row *TradesHistoryRecord) GetRecordFromXml(attrs *[]xml.Attr) {

	var tradeDate string
	layout := "2006-01-02"
	layoutTime := "2006-01-02 15:04:05"
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
