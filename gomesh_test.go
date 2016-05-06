package gomesh2016

import (
	"testing"
		//	"fmt"
	//	"compress/gzip"
)

func makeNextSupplemental(suppChannel chan *SupplementalRecord) nextChannelItem {
	return func() bool {
		_, err := <-suppChannel
		return err
	}
}

func makeNextDescriptor(descChannel chan *DescriptorRecord) nextChannelItem {
	return func() bool {
		_, err := <-descChannel
		return err
	}
}

func makeNextQualifier(qualChannel chan *QualifierRecord) nextChannelItem {
	return func() bool {
		_, err := <-qualChannel
		return err
	}
}


func TestReadManyDescriptions(t *testing.T) {
	descFileName := "/home/gnewton/work/meshXml2016/desc2016.gz"
	descChannel, file, err := DescriptorChannelFromFile(descFileName)
	defer file.Close()

	if err != nil {
		t.Fatal("Error occured", err)
		return
	}
	if countChannel(makeNextDescriptor(descChannel)) != 27885 {
		t.Fail()
	}
}

func TestReadManyQualifiers(t *testing.T) {
	qualFileName := "/home/gnewton/work/meshXml2016/qual2016.xml"
	qualChannel, file, err := QualifierChannelFromFile(qualFileName)
	defer file.Close()

	if err != nil {
		t.Fatal("Error occured", err)
		return
	}
	n := countChannel(makeNextQualifier(qualChannel))

	if n != 82 {
		t.Fail()
	}
}


func TestReadManySupplementalRecords(t *testing.T) {
	suppFileName := "/home/gnewton/work/meshXml2016/supp2016.gz"
	suppChannel, file, err := SupplementalChannelFromFile(suppFileName)
	defer file.Close()

	if err != nil {
		t.Fatal("Error occured", err)
		return
	}
	if countChannel(makeNextSupplemental(suppChannel)) != 235000 {
		t.Fail()
	}
}


func countChannel(nextDescriptor nextChannelItem) int {
	counter := 0
	for {
		val := nextDescriptor()
		if !val {
			break
		}
		counter += 1
	}
	return counter
}
