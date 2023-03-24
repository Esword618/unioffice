// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_Layout struct {
	ManualLayout *CT_ManualLayout
	ExtLst       *CT_ExtensionList
}

func NewCT_Layout() *CT_Layout {
	ret := &CT_Layout{}
	return ret
}

func (m *CT_Layout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ManualLayout != nil {
		semanualLayout := xml.StartElement{Name: xml.Name{Local: "c:manualLayout"}}
		e.EncodeElement(m.ManualLayout, semanualLayout)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Layout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Layout:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "manualLayout"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "manualLayout"}:
				m.ManualLayout = NewCT_ManualLayout()
				if err := d.DecodeElement(m.ManualLayout, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Layout %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Layout
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Layout and its children
func (m *CT_Layout) Validate() error {
	return m.ValidateWithPath("CT_Layout")
}

// ValidateWithPath validates the CT_Layout and its children, prefixing error messages with path
func (m *CT_Layout) ValidateWithPath(path string) error {
	if m.ManualLayout != nil {
		if err := m.ManualLayout.ValidateWithPath(path + "/ManualLayout"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}