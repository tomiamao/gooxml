package ooxml

import (
	"encoding/xml"
	"gooxml/opc"
	"io"
)

type PresentationSlideMaster struct {
	XMLName   xml.Name `xml:"sldMaster"`
	SlideData SlideData
}

func buildPresentationSlideMaster(pkg *opc.Package, partName string, in io.Reader) error {
	slideLayout := &PresentationSlideMaster{}
	dec := xml.NewDecoder(in)
	if err := dec.Decode(slideLayout); err != nil {
		return err
	}
	part := pkg.FindPart(partName)
	if part == nil {
		panic(partName)
	}
	part.Content = slideLayout
	return nil
}
