package AbstractTableStruct

import "encoding/xml"

type AbstractRecord interface {
	GetRecordFromXml(attrs *[]xml.Attr)
}
