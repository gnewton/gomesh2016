
package gomesh2016

type SupplementalRoot struct {
	SupplementalRecordSet *SupplementalRecordSet `xml:"SupplementalRecordSet,omitempty" json:"SupplementalRecordSet,omitempty"`
}

type SupplementalRecordSet struct {
	Attr_LanguageCode string `xml:" LanguageCode,attr"  json:",omitempty"`
	SupplementalRecord []*SupplementalRecord `xml:"SupplementalRecord,omitempty" json:"SupplementalRecord,omitempty"`
}

type SupplementalRecord struct {
	Attr_SCRClass string `xml:" SCRClass,attr"  json:",omitempty"`
	ConceptList *ConceptList `xml:"ConceptList,omitempty" json:"ConceptList,omitempty"`
	DateCreated *DateCreated `xml:"DateCreated,omitempty" json:"DateCreated,omitempty"`
	DateRevised *DateRevised `xml:"DateRevised,omitempty" json:"DateRevised,omitempty"`
	Frequency *Frequency `xml:"Frequency,omitempty" json:"Frequency,omitempty"`
	HeadingMappedToList *HeadingMappedToList `xml:"HeadingMappedToList,omitempty" json:"HeadingMappedToList,omitempty"`
	IndexingInformationList *IndexingInformationList `xml:"IndexingInformationList,omitempty" json:"IndexingInformationList,omitempty"`
	Note *Note `xml:"Note,omitempty" json:"Note,omitempty"`
	PharmacologicalActionList *PharmacologicalActionList `xml:"PharmacologicalActionList,omitempty" json:"PharmacologicalActionList,omitempty"`
	PreviousIndexingList *PreviousIndexingList `xml:"PreviousIndexingList,omitempty" json:"PreviousIndexingList,omitempty"`
	SourceList *SourceList `xml:"SourceList,omitempty" json:"SourceList,omitempty"`
	SupplementalRecordName *SupplementalRecordName `xml:"SupplementalRecordName,omitempty" json:"SupplementalRecordName,omitempty"`
	SupplementalRecordUI *SupplementalRecordUI `xml:"SupplementalRecordUI,omitempty" json:"SupplementalRecordUI,omitempty"`
}

type Note struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Frequency struct {
	Text int16 `xml:",chardata" json:",omitempty"`
}

type SourceList struct {
	Source []*Source `xml:"Source,omitempty" json:"Source,omitempty"`
}

type Source struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IndexingInformationList struct {
	IndexingInformation []*IndexingInformation `xml:"IndexingInformation,omitempty" json:"IndexingInformation,omitempty"`
}

type IndexingInformation struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
	QualifierReferredTo *QualifierReferredTo `xml:"QualifierReferredTo,omitempty" json:"QualifierReferredTo,omitempty"`
}


type SupplementalRecordName struct {
	String *String `xml:"String,omitempty" json:"String,omitempty"`
}


type HeadingMappedToList struct {
	HeadingMappedTo []*HeadingMappedTo `xml:"HeadingMappedTo,omitempty" json:"HeadingMappedTo,omitempty"`
}

type HeadingMappedTo struct {
	DescriptorReferredTo *DescriptorReferredTo `xml:"DescriptorReferredTo,omitempty" json:"DescriptorReferredTo,omitempty"`
	QualifierReferredTo *QualifierReferredTo `xml:"QualifierReferredTo,omitempty" json:"QualifierReferredTo,omitempty"`
}



type SupplementalRecordUI struct {
	Text string `xml:",chardata" json:",omitempty"`
}



