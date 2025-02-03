package ooxml

import (
	"errors"
	"io"
	"log"

	"gooxml/opc"
)

const (
	ContentTypePresentationDocument    = "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml"
	ContentTypePresentationSlide       = "application/vnd.openxmlformats-officedocument.presentationml.slide+xml"
	ContentTypePresentationSlideMaster = "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"
	ContentTypePresentationSlideLayout = "application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml"
	ContentTypeImagePNG                = "image/png"
	ContentTypeImageGIF                = "image/gif"
	ContentTypeImageJPEG               = "image/jpeg"
	ContentTypeImageJPG                = "image/jpg"
)

func init() {
	opc.RegisterReadFormat(ContentTypePresentationDocument, buildPresentationDocument)
	opc.RegisterReadFormat(ContentTypePresentationSlide, buildPresentationSlide)
	opc.RegisterReadFormat(ContentTypePresentationSlideLayout, buildPresentationSlideLayout)
	opc.RegisterReadFormat(ContentTypePresentationSlideMaster, buildPresentationSlideMaster)
	opc.RegisterReadFormat(ContentTypeImagePNG, buildPresentationImagePNG)
	opc.RegisterReadFormat(ContentTypeImageGIF, buildPresentationImageGIF)
	opc.RegisterReadFormat(ContentTypeImageJPEG, buildPresentationImageJPEG)
	opc.RegisterReadFormat(ContentTypeImageJPG, buildPresentationImageJPG)
}

func ReadPresentationDocument(in io.Reader) (*PresentationDocument, error) {
	pkg, _ := opc.Read(in)
	return findPresentationDocument(pkg)
}

func OpenPresentationDocument(name string) (*PresentationDocument, error) {
	pkg, err := opc.Open(name)
	log.Println(err)
	return findPresentationDocument(pkg)
}

func findPresentationDocument(pkg *opc.Package) (*PresentationDocument, error) {
	parts := pkg.FindPartsByRelationOn(&pkg.Part, func(rel *opc.Relationship) bool { return rel.Type == RelationTypeOfficeDocument })
	if len(parts) != 1 || parts[0].Content == nil {
		return nil, errors.New("it is not a PresentationDocument")
	}
	return parts[0].Content.(*PresentationDocument), nil
}
