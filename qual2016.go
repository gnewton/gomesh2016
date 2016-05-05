
package gomesh2016

type QualifierRoot struct {
	QualifierRecordSet *QualifierRecordSet `xml:"QualifierRecordSet,omitempty" json:"QualifierRecordSet,omitempty"`
}

type QualifierRecordSet struct {
	Attr_LanguageCode string `xml:" LanguageCode,attr"  json:",omitempty"`
	QualifierRecord []*QualifierRecord `xml:"QualifierRecord,omitempty" json:"QualifierRecord,omitempty"`
}

type QualifierRecord struct {
	Annotation *Annotation `xml:"Annotation,omitempty" json:"Annotation,omitempty"`
	ConceptList *ConceptList `xml:"ConceptList,omitempty" json:"ConceptList,omitempty"`
	DateCreated *DateCreated `xml:"DateCreated,omitempty" json:"DateCreated,omitempty"`
	DateEstablished *DateEstablished `xml:"DateEstablished,omitempty" json:"DateEstablished,omitempty"`
	DateRevised *DateRevised `xml:"DateRevised,omitempty" json:"DateRevised,omitempty"`
	HistoryNote *HistoryNote `xml:"HistoryNote,omitempty" json:"HistoryNote,omitempty"`
	OnlineNote *OnlineNote `xml:"OnlineNote,omitempty" json:"OnlineNote,omitempty"`
	QualifierName *QualifierName `xml:"QualifierName,omitempty" json:"QualifierName,omitempty"`
	QualifierUI *QualifierUI `xml:"QualifierUI,omitempty" json:"QualifierUI,omitempty"`
	TreeNumberList *TreeNumberList `xml:"TreeNumberList,omitempty" json:"TreeNumberList,omitempty"`
}

type QualifierUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type QualifierName struct {
	String *String `xml:"String,omitempty" json:"String,omitempty"`
}







