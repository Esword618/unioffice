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
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_VolTopic struct {
	// Type
	TAttr ST_VolValueType
	// Topic Value
	V string
	// Strings in Subtopic
	Stp []string
	// References
	Tr []*CT_VolTopicRef
}

func NewCT_VolTopic() *CT_VolTopic {
	ret := &CT_VolTopic{}
	return ret
}

func (m *CT_VolTopic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TAttr != ST_VolValueTypeUnset {
		attr, err := m.TAttr.MarshalXMLAttr(xml.Name{Local: "t"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	sev := xml.StartElement{Name: xml.Name{Local: "ma:v"}}
	unioffice.AddPreserveSpaceAttr(&sev, m.V)
	e.EncodeElement(m.V, sev)
	if m.Stp != nil {
		sestp := xml.StartElement{Name: xml.Name{Local: "ma:stp"}}
		for _, c := range m.Stp {
			e.EncodeElement(c, sestp)
		}
	}
	setr := xml.StartElement{Name: xml.Name{Local: "ma:tr"}}
	for _, c := range m.Tr {
		e.EncodeElement(c, setr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_VolTopic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "t" {
			m.TAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_VolTopic:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "v"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "v"}:
				if err := d.DecodeElement(&m.V, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "stp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "stp"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Stp = append(m.Stp, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "tr"}:
				tmp := NewCT_VolTopicRef()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tr = append(m.Tr, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_VolTopic %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_VolTopic
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_VolTopic and its children
func (m *CT_VolTopic) Validate() error {
	return m.ValidateWithPath("CT_VolTopic")
}

// ValidateWithPath validates the CT_VolTopic and its children, prefixing error messages with path
func (m *CT_VolTopic) ValidateWithPath(path string) error {
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	for i, v := range m.Tr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}