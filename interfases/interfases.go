package interfases

import "encoding/xml"

type TableRecord interface {
	GetRecordFromXml(attrs *[]xml.Attr)
}
