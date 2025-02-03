package ooxml

import (
	"encoding/xml"
	"io"

	"github.com/tomiamao/gooxml/opc"
)

type PresentationSlideLayout struct {
	XMLName   xml.Name `xml:"sldLayout"`
	SlideData SlideData
}

func buildPresentationSlideLayout(pkg *opc.Package, partName string, in io.Reader) error {
	slideLayout := &PresentationSlideLayout{}
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
