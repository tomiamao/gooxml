package ooxml

import (
	"gooxml/opc"
	"io"
	"log"
)

func buildPresentationImagePNG(pkg *opc.Package, partName string, in io.Reader) error {
	part := pkg.FindPart(partName)

	// max 20MB for images
	imgData := make([]byte, 20*1024*1024)
	n, err := in.Read(imgData)
	if err != nil {
		return err
	}

	log.Printf("********* IMAGE: %s - %d num bytes read\n", partName, n)

	part.Content = imgData[:n]
	return nil
}

func buildPresentationImageGIF(pkg *opc.Package, partName string, in io.Reader) error {
	part := pkg.FindPart(partName)

	// max 20MB for images
	imgData := make([]byte, 20*1024*1024)
	n, err := in.Read(imgData)
	if err != nil {
		return err
	}

	log.Printf("********* IMAGE: %s - %d num bytes read\n", partName, n)

	part.Content = imgData[:n]
	return nil
}

func buildPresentationImageJPEG(pkg *opc.Package, partName string, in io.Reader) error {
	part := pkg.FindPart(partName)

	// max 20MB for images
	imgData := make([]byte, 20*1024*1024)
	n, err := in.Read(imgData)
	if err != nil {
		return err
	}

	log.Printf("********* IMAGE: %s - %d num bytes read\n", partName, n)

	part.Content = imgData[:n]
	return nil
}

func buildPresentationImageJPG(pkg *opc.Package, partName string, in io.Reader) error {
	part := pkg.FindPart(partName)

	// max 20MB for images
	imgData := make([]byte, 20*1024*1024)
	n, err := in.Read(imgData)
	if err != nil {
		return err
	}

	log.Printf("********* IMAGE: %s - %d num bytes read\n", partName, n)

	part.Content = imgData[:n]
	return nil
}
