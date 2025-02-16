package AbstractTableStruct

import (
	"bytes"
	"encoding/xml"
)

type Rows []AbstractRecord
type Baskets []Rows

type AbstractTable interface {
	GetBlockRowsDelimeters() string
	GetCountRecordsInBasket() int
	SetXmlSourse() string
	GetTargetRecord() AbstractRecord
}

type (
	Table struct {
		rows    Rows
		baskets Baskets
	}
)

func (tbl Table) GetRecordsFromZip(srcTable AbstractTable) int {

	xmlData := bytes.NewBufferString((srcTable).SetXmlSourse())
	rowElementName := (srcTable).GetBlockRowsDelimeters()
	countRecordsInBasket := (srcTable).GetCountRecordsInBasket()
	d := xml.NewDecoder(xmlData)
	var tRec AbstractRecord
	countAppendedRecords := 0
	countRecordAppended := 0
	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == rowElementName {
				tRec = (srcTable).GetTargetRecord()
				tRec.GetRecordFromXml(&se.Attr)
				countAppendedRecords++
				countRecordAppended++
				if countAppendedRecords <= countRecordsInBasket {
					tbl.rows = append(tbl.rows, tRec)
				} else {
					tbl.baskets = append(tbl.baskets, tbl.rows)
					tbl.rows = Rows{}
					countAppendedRecords = 0
				}
			}
		}
	}
	if countAppendedRecords > 0 {
		tbl.baskets = append(tbl.baskets, tbl.rows)
	}
	return countRecordAppended
}
