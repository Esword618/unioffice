// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_AutoSortScope struct {
	// Auto Sort Scope
	PivotArea *CT_PivotArea
}

func NewCT_AutoSortScope() *CT_AutoSortScope {
	ret := &CT_AutoSortScope{}
	ret.PivotArea = NewCT_PivotArea()
	return ret
}

func (m *CT_AutoSortScope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sepivotArea := xml.StartElement{Name: xml.Name{Local: "ma:pivotArea"}}
	e.EncodeElement(m.PivotArea, sepivotArea)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AutoSortScope) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PivotArea = NewCT_PivotArea()
lCT_AutoSortScope:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pivotArea"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "pivotArea"}:
				if err := d.DecodeElement(m.PivotArea, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_AutoSortScope %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AutoSortScope
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AutoSortScope and its children
func (m *CT_AutoSortScope) Validate() error {
	return m.ValidateWithPath("CT_AutoSortScope")
}

// ValidateWithPath validates the CT_AutoSortScope and its children, prefixing error messages with path
func (m *CT_AutoSortScope) ValidateWithPath(path string) error {
	if err := m.PivotArea.ValidateWithPath(path + "/PivotArea"); err != nil {
		return err
	}
	return nil
}
