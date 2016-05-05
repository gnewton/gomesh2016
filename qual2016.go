
package gomesh2016

type Root struct {
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

type String struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DateEstablished struct {
	Day *Day `xml:"Day,omitempty" json:"Day,omitempty"`
	Month *Month `xml:"Month,omitempty" json:"Month,omitempty"`
	Year *Year `xml:"Year,omitempty" json:"Year,omitempty"`
}

type Month struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

type Day struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

type Year struct {
	Text int16 `xml:",chardata" json:",omitempty"`
}

type OnlineNote struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DateCreated struct {
	Day *Day `xml:"Day,omitempty" json:"Day,omitempty"`
	Month *Month `xml:"Month,omitempty" json:"Month,omitempty"`
	Year *Year `xml:"Year,omitempty" json:"Year,omitempty"`
}

type DateRevised struct {
	Day *Day `xml:"Day,omitempty" json:"Day,omitempty"`
	Month *Month `xml:"Month,omitempty" json:"Month,omitempty"`
	Year *Year `xml:"Year,omitempty" json:"Year,omitempty"`
}

type Annotation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HistoryNote struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TreeNumberList struct {
	TreeNumber []*TreeNumber `xml:"TreeNumber,omitempty" json:"TreeNumber,omitempty"`
}

type TreeNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ConceptList struct {
	Concept []*Concept `xml:"Concept,omitempty" json:"Concept,omitempty"`
}

type Concept struct {
	Attr_PreferredConceptYN string `xml:" PreferredConceptYN,attr"  json:",omitempty"`
	ConceptName *ConceptName `xml:"ConceptName,omitempty" json:"ConceptName,omitempty"`
	ConceptRelationList *ConceptRelationList `xml:"ConceptRelationList,omitempty" json:"ConceptRelationList,omitempty"`
	ConceptUI *ConceptUI `xml:"ConceptUI,omitempty" json:"ConceptUI,omitempty"`
	ScopeNote *ScopeNote `xml:"ScopeNote,omitempty" json:"ScopeNote,omitempty"`
	TermList *TermList `xml:"TermList,omitempty" json:"TermList,omitempty"`
}

type ConceptUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ConceptName struct {
	String *String `xml:"String,omitempty" json:"String,omitempty"`
}

type ScopeNote struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ConceptRelationList struct {
	ConceptRelation []*ConceptRelation `xml:"ConceptRelation,omitempty" json:"ConceptRelation,omitempty"`
}

type ConceptRelation struct {
	Attr_RelationName string `xml:" RelationName,attr"  json:",omitempty"`
	Concept1UI *Concept1UI `xml:"Concept1UI,omitempty" json:"Concept1UI,omitempty"`
	Concept2UI *Concept2UI `xml:"Concept2UI,omitempty" json:"Concept2UI,omitempty"`
}

type Concept1UI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Concept2UI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TermList struct {
	Term []*Term `xml:"Term,omitempty" json:"Term,omitempty"`
}

type Term struct {
	Attr_ConceptPreferredTermYN string `xml:" ConceptPreferredTermYN,attr"  json:",omitempty"`
	Attr_IsPermutedTermYN string `xml:" IsPermutedTermYN,attr"  json:",omitempty"`
	Attr_LexicalTag string `xml:" LexicalTag,attr"  json:",omitempty"`
	Attr_RecordPreferredTermYN string `xml:" RecordPreferredTermYN,attr"  json:",omitempty"`
	Abbreviation *Abbreviation `xml:"Abbreviation,omitempty" json:"Abbreviation,omitempty"`
	DateCreated *DateCreated `xml:"DateCreated,omitempty" json:"DateCreated,omitempty"`
	EntryVersion *EntryVersion `xml:"EntryVersion,omitempty" json:"EntryVersion,omitempty"`
	SortVersion *SortVersion `xml:"SortVersion,omitempty" json:"SortVersion,omitempty"`
	String *String `xml:"String,omitempty" json:"String,omitempty"`
	TermUI *TermUI `xml:"TermUI,omitempty" json:"TermUI,omitempty"`
}

type EntryVersion struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SortVersion struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TermUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Abbreviation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

