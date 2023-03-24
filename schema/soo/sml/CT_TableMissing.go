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
)

type CT_TableMissing struct {
}

func NewCT_TableMissing() *CT_TableMissing {
	ret := &CT_TableMissing{}
	return ret
}

func (m *CT_TableMissing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableMissing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TableMissing: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TableMissing and its children
func (m *CT_TableMissing) Validate() error {
	return m.ValidateWithPath("CT_TableMissing")
}

// ValidateWithPath validates the CT_TableMissing and its children, prefixing error messages with path
func (m *CT_TableMissing) ValidateWithPath(path string) error {
	return nil
}
