package main

import (
	"log"
	"os"
	"strings"

	"gooxml/ooxml"
	"gooxml/opc"
)

func main() {
	var doc *ooxml.PresentationDocument
	var err error
	if len(os.Args) > 1 {
		doc, err = ooxml.OpenPresentationDocument(os.Args[1])
	} else {
		doc, err = ooxml.ReadPresentationDocument(os.Stdin)
	}
	if err != nil || doc == nil {
		log.Fatal(err)
	}

	for _, v := range doc.Package.Parts {
		if v.Name == "/docProps/thumbnail.jpeg" {
			err := os.WriteFile(strings.ReplaceAll(v.Name, "/", "_"), v.Content.([]byte), 0644)
			if err != nil {
				log.Println(err)
			}
		} else if v.ContentType == ooxml.ContentTypePresentationSlide &&
			(v.Name == "/ppt/slides/slide1.xml" ||
				v.Name == "/ppt/slides/slide2.xml") {
			log.Printf("%s - %s\n", v.ContentType, v.Name)

			log.Println("Relationships")
			for _, vv := range v.Relationships {
				log.Printf("%s - %s - %s\n", vv.ID, vv.Target, vv.Type)

				if vv.Type == "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image" {
					// find part
					parts := doc.Package.FindPartsByRelationOn(v, func(rel *opc.Relationship) bool {
						log.Printf("%s ------ %s ------ %s\n", rel.ID, rel.Type, rel.Target)
						return rel.ID == vv.ID && rel.Type == vv.Type && rel.Target == vv.Target
					})
					for _, part := range parts {
						err := os.WriteFile(strings.ReplaceAll(part.Name, "/", "_"), part.Content.([]byte), 0644)
						if err != nil {
							log.Println(err)
						}
					}
				}

				if vv.Type == "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideLayout" {
					// find part
					parts := doc.Package.FindPartsByRelationOn(v, func(rel *opc.Relationship) bool {
						return rel.ID == vv.ID && rel.Type == vv.Type && rel.Target == vv.Target
					})
					for _, part := range parts {
						vvv := part.Content.(*ooxml.PresentationSlideLayout)
						log.Printf("***** %s\n", vvv.XMLName)
						log.Printf("\t ********* %s\n \t\t %s\n", vvv.SlideData.XMLName, vvv.SlideData.ShapeTree.String())
					}
				}

				if vv.Type == "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster" {
					// find part
					parts := doc.Package.FindPartsByRelationOn(v, func(rel *opc.Relationship) bool {
						return rel.ID == vv.ID && rel.Type == vv.Type && rel.Target == vv.Target
					})
					for _, part := range parts {
						vvv := part.Content.(*ooxml.PresentationSlideMaster)
						log.Printf("***** %s\n", vvv.XMLName)
						log.Printf("\t ********* %s\n \t\t %s\n", vvv.SlideData.XMLName, vvv.SlideData.ShapeTree.String())
					}
				}
			}
		}
	}

	/*
		slides := doc.Slides()
		for _, slide := range slides {
			fmt.Println(slide.String())
		}
	*/
}
