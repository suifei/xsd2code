package types

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// XSDSchema represents the root XSD schema element
type XSDSchema struct {
	XMLName              xml.Name            `xml:"http://www.w3.org/2001/XMLSchema schema"`
	TargetNamespace      string              `xml:"targetNamespace,attr"`
	ElementFormDefault   string              `xml:"elementFormDefault,attr"`
	AttributeFormDefault string              `xml:"attributeFormDefault,attr"`
	Xmlns                map[string]string   `xml:"-"`
	Elements             []XSDElement        `xml:"element"`
	ComplexTypes         []XSDComplexType    `xml:"complexType"`
	SimpleTypes          []XSDSimpleType     `xml:"simpleType"`
	Groups               []XSDGroup          `xml:"group"`
	AttributeGroups      []XSDAttributeGroup `xml:"attributeGroup"`
	Imports              []XSDImport         `xml:"import"`
	Includes             []XSDInclude        `xml:"include"`
	Annotations          []XSDAnnotation     `xml:"annotation"`
}

// XSDElement represents an XSD element
type XSDElement struct {
	XMLName     xml.Name        `xml:"element"`
	Name        string          `xml:"name,attr"`
	Type        string          `xml:"type,attr"`
	Ref         string          `xml:"ref,attr"`
	MinOccurs   string          `xml:"minOccurs,attr"`
	MaxOccurs   string          `xml:"maxOccurs,attr"`
	Default     string          `xml:"default,attr"`
	Fixed       string          `xml:"fixed,attr"`
	Nillable    string          `xml:"nillable,attr"`
	ComplexType *XSDComplexType `xml:"complexType"`
	SimpleType  *XSDSimpleType  `xml:"simpleType"`
	Annotation  *XSDAnnotation  `xml:"annotation"`
}

// XSDComplexType represents an XSD complex type
type XSDComplexType struct {
	XMLName         xml.Name               `xml:"complexType"`
	Name            string                 `xml:"name,attr"`
	Mixed           string                 `xml:"mixed,attr"`
	Abstract        string                 `xml:"abstract,attr"`
	Sequence        *XSDSequence           `xml:"sequence"`
	Choice          *XSDChoice             `xml:"choice"`
	All             *XSDAll                `xml:"all"`
	Group           *XSDGroupRef           `xml:"group"`
	Attributes      []XSDAttribute         `xml:"attribute"`
	AttributeGroups []XSDAttributeGroupRef `xml:"attributeGroup"`
	SimpleContent   *XSDSimpleContent      `xml:"simpleContent"`
	ComplexContent  *XSDComplexContent     `xml:"complexContent"`
	Annotation      *XSDAnnotation         `xml:"annotation"`
}

// XSDSimpleType represents an XSD simple type
type XSDSimpleType struct {
	XMLName     xml.Name        `xml:"simpleType"`
	Name        string          `xml:"name,attr"`
	Restriction *XSDRestriction `xml:"restriction"`
	Union       *XSDUnion       `xml:"union"`
	List        *XSDList        `xml:"list"`
	Annotation  *XSDAnnotation  `xml:"annotation"`
}

// XSDSequence represents an XSD sequence
type XSDSequence struct {
	XMLName   xml.Name      `xml:"sequence"`
	MinOccurs string        `xml:"minOccurs,attr"`
	MaxOccurs string        `xml:"maxOccurs,attr"`
	Elements  []XSDElement  `xml:"element"`
	Groups    []XSDGroupRef `xml:"group"`
	Choices   []XSDChoice   `xml:"choice"`
	Sequences []XSDSequence `xml:"sequence"`
}

// XSDChoice represents an XSD choice
type XSDChoice struct {
	XMLName   xml.Name      `xml:"choice"`
	MinOccurs string        `xml:"minOccurs,attr"`
	MaxOccurs string        `xml:"maxOccurs,attr"`
	Elements  []XSDElement  `xml:"element"`
	Groups    []XSDGroupRef `xml:"group"`
	Choices   []XSDChoice   `xml:"choice"`
	Sequences []XSDSequence `xml:"sequence"`
}

// XSDAll represents an XSD all
type XSDAll struct {
	XMLName   xml.Name     `xml:"all"`
	MinOccurs string       `xml:"minOccurs,attr"`
	MaxOccurs string       `xml:"maxOccurs,attr"`
	Elements  []XSDElement `xml:"element"`
}

// XSDAttribute represents an XSD attribute
type XSDAttribute struct {
	XMLName    xml.Name       `xml:"attribute"`
	Name       string         `xml:"name,attr"`
	Type       string         `xml:"type,attr"`
	Ref        string         `xml:"ref,attr"`
	Use        string         `xml:"use,attr"`
	Default    string         `xml:"default,attr"`
	Fixed      string         `xml:"fixed,attr"`
	Form       string         `xml:"form,attr"`
	SimpleType *XSDSimpleType `xml:"simpleType"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDGroup represents an XSD group
type XSDGroup struct {
	XMLName    xml.Name       `xml:"group"`
	Name       string         `xml:"name,attr"`
	Ref        string         `xml:"ref,attr"`
	MinOccurs  string         `xml:"minOccurs,attr"`
	MaxOccurs  string         `xml:"maxOccurs,attr"`
	Sequence   *XSDSequence   `xml:"sequence"`
	Choice     *XSDChoice     `xml:"choice"`
	All        *XSDAll        `xml:"all"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDGroupRef represents an XSD group reference
type XSDGroupRef struct {
	XMLName    xml.Name       `xml:"group"`
	Ref        string         `xml:"ref,attr"`
	MinOccurs  string         `xml:"minOccurs,attr"`
	MaxOccurs  string         `xml:"maxOccurs,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDAttributeGroup represents an XSD attribute group
type XSDAttributeGroup struct {
	XMLName         xml.Name               `xml:"attributeGroup"`
	Name            string                 `xml:"name,attr"`
	Ref             string                 `xml:"ref,attr"`
	Attributes      []XSDAttribute         `xml:"attribute"`
	AttributeGroups []XSDAttributeGroupRef `xml:"attributeGroup"`
	Annotation      *XSDAnnotation         `xml:"annotation"`
}

// XSDAttributeGroupRef represents an XSD attribute group reference
type XSDAttributeGroupRef struct {
	XMLName    xml.Name       `xml:"attributeGroup"`
	Ref        string         `xml:"ref,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDRestriction represents an XSD restriction
type XSDRestriction struct {
	XMLName         xml.Name               `xml:"restriction"`
	Base            string                 `xml:"base,attr"`
	Enumerations    []XSDEnumeration       `xml:"enumeration"`
	Pattern         *XSDPattern            `xml:"pattern"`
	MinLength       *XSDMinLength          `xml:"minLength"`
	MaxLength       *XSDMaxLength          `xml:"maxLength"`
	MinInclusive    *XSDMinInclusive       `xml:"minInclusive"`
	MaxInclusive    *XSDMaxInclusive       `xml:"maxInclusive"`
	MinExclusive    *XSDMinExclusive       `xml:"minExclusive"`
	MaxExclusive    *XSDMaxExclusive       `xml:"maxExclusive"`
	TotalDigits     *XSDTotalDigits        `xml:"totalDigits"`
	FractionDigits  *XSDFractionDigits     `xml:"fractionDigits"`
	Attributes      []XSDAttribute         `xml:"attribute"`
	AttributeGroups []XSDAttributeGroupRef `xml:"attributeGroup"`
	Sequence        *XSDSequence           `xml:"sequence"`
	Choice          *XSDChoice             `xml:"choice"`
	All             *XSDAll                `xml:"all"`
	Group           *XSDGroupRef           `xml:"group"`
	Annotation      *XSDAnnotation         `xml:"annotation"`
}

// XSDEnumeration represents an XSD enumeration
type XSDEnumeration struct {
	XMLName    xml.Name       `xml:"enumeration"`
	Value      string         `xml:"value,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDPattern represents an XSD pattern
type XSDPattern struct {
	XMLName    xml.Name       `xml:"pattern"`
	Value      string         `xml:"value,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDMinLength represents an XSD minLength
type XSDMinLength struct {
	XMLName    xml.Name       `xml:"minLength"`
	Value      string         `xml:"value,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDMaxLength represents an XSD maxLength
type XSDMaxLength struct {
	XMLName    xml.Name       `xml:"maxLength"`
	Value      string         `xml:"value,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDMinInclusive represents an XSD minInclusive
type XSDMinInclusive struct {
	XMLName    xml.Name       `xml:"minInclusive"`
	Value      string         `xml:"value,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDMaxInclusive represents an XSD maxInclusive
type XSDMaxInclusive struct {
	XMLName    xml.Name       `xml:"maxInclusive"`
	Value      string         `xml:"value,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDMinExclusive represents an XSD minExclusive
type XSDMinExclusive struct {
	XMLName    xml.Name       `xml:"minExclusive"`
	Value      string         `xml:"value,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDMaxExclusive represents an XSD maxExclusive
type XSDMaxExclusive struct {
	XMLName    xml.Name       `xml:"maxExclusive"`
	Value      string         `xml:"value,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDTotalDigits represents an XSD totalDigits
type XSDTotalDigits struct {
	XMLName    xml.Name       `xml:"totalDigits"`
	Value      string         `xml:"value,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDFractionDigits represents an XSD fractionDigits
type XSDFractionDigits struct {
	XMLName    xml.Name       `xml:"fractionDigits"`
	Value      string         `xml:"value,attr"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDUnion represents an XSD union
type XSDUnion struct {
	XMLName     xml.Name        `xml:"union"`
	MemberTypes string          `xml:"memberTypes,attr"`
	SimpleTypes []XSDSimpleType `xml:"simpleType"`
	Annotation  *XSDAnnotation  `xml:"annotation"`
}

// XSDList represents an XSD list
type XSDList struct {
	XMLName    xml.Name       `xml:"list"`
	ItemType   string         `xml:"itemType,attr"`
	SimpleType *XSDSimpleType `xml:"simpleType"`
	Annotation *XSDAnnotation `xml:"annotation"`
}

// XSDSimpleContent represents an XSD simpleContent
type XSDSimpleContent struct {
	XMLName     xml.Name        `xml:"simpleContent"`
	Extension   *XSDExtension   `xml:"extension"`
	Restriction *XSDRestriction `xml:"restriction"`
	Annotation  *XSDAnnotation  `xml:"annotation"`
}

// XSDComplexContent represents an XSD complexContent
type XSDComplexContent struct {
	XMLName     xml.Name        `xml:"complexContent"`
	Mixed       string          `xml:"mixed,attr"`
	Extension   *XSDExtension   `xml:"extension"`
	Restriction *XSDRestriction `xml:"restriction"`
	Annotation  *XSDAnnotation  `xml:"annotation"`
}

// XSDExtension represents an XSD extension
type XSDExtension struct {
	XMLName         xml.Name               `xml:"extension"`
	Base            string                 `xml:"base,attr"`
	Sequence        *XSDSequence           `xml:"sequence"`
	Choice          *XSDChoice             `xml:"choice"`
	All             *XSDAll                `xml:"all"`
	Group           *XSDGroupRef           `xml:"group"`
	Attributes      []XSDAttribute         `xml:"attribute"`
	AttributeGroups []XSDAttributeGroupRef `xml:"attributeGroup"`
	Annotation      *XSDAnnotation         `xml:"annotation"`
}

// XSDImport represents an XSD import
type XSDImport struct {
	XMLName        xml.Name       `xml:"import"`
	Namespace      string         `xml:"namespace,attr"`
	SchemaLocation string         `xml:"schemaLocation,attr"`
	Annotation     *XSDAnnotation `xml:"annotation"`
}

// XSDInclude represents an XSD include
type XSDInclude struct {
	XMLName        xml.Name       `xml:"include"`
	SchemaLocation string         `xml:"schemaLocation,attr"`
	Annotation     *XSDAnnotation `xml:"annotation"`
}

// XSDAnnotation represents an XSD annotation
type XSDAnnotation struct {
	XMLName       xml.Name           `xml:"annotation"`
	Documentation []XSDDocumentation `xml:"documentation"`
	AppInfo       []XSDAppInfo       `xml:"appinfo"`
}

// XSDDocumentation represents an XSD documentation
type XSDDocumentation struct {
	XMLName xml.Name `xml:"documentation"`
	Lang    string   `xml:"lang,attr"`
	Source  string   `xml:"source,attr"`
	Content string   `xml:",chardata"`
}

// XSDAppInfo represents an XSD appinfo
type XSDAppInfo struct {
	XMLName xml.Name `xml:"appinfo"`
	Source  string   `xml:"source,attr"`
	Content string   `xml:",chardata"`
}

// GoType represents a generated Go type
type GoType struct {
	Name      string
	Package   string
	XMLName   string
	Namespace string
	Fields    []GoField
	Constants []GoConstant
	Comment   string
	IsEnum    bool
	BaseType  string
}

// GoField represents a field in a Go struct
type GoField struct {
	Name        string
	Type        string
	XMLTag      string
	JSONTag     string
	Comment     string
	IsAttribute bool
	IsElement   bool
	IsOptional  bool
	IsArray     bool
	MinOccurs   int
	MaxOccurs   int // -1 for unbounded
}

// GoConstant represents a Go constant (for enums)
type GoConstant struct {
	Name    string
	Value   string
	Comment string
}

// TypeMapping represents XSD to Go type mapping
type TypeMapping struct {
	XSDType string
	GoType  string
}

// GetBuiltinTypeMappings returns the mapping from XSD built-in types to Go types
func GetBuiltinTypeMappings() []TypeMapping {
	return []TypeMapping{
		// Standard XSD types
		{"string", "string"},
		{"normalizedString", "string"},
		{"token", "string"},
		{"anyURI", "string"},
		{"language", "string"},
		{"NMTOKEN", "string"},
		{"NMTOKENS", "[]string"},
		{"Name", "string"},
		{"NCName", "string"},
		{"ID", "string"},
		{"IDREF", "string"},
		{"IDREFS", "[]string"},
		{"ENTITY", "string"},
		{"ENTITIES", "[]string"},
		{"QName", "string"},

		{"boolean", "bool"},

		{"decimal", "float64"},
		{"float", "float32"},
		{"double", "float64"},

		{"duration", "string"}, // Could be time.Duration in future
		{"dateTime", "time.Time"},
		{"time", "string"},
		{"date", "string"},
		{"gYearMonth", "string"},
		{"gYear", "string"},
		{"gMonthDay", "string"},
		{"gDay", "string"},
		{"gMonth", "string"},

		{"hexBinary", "[]byte"},
		{"base64Binary", "[]byte"},

		{"integer", "int64"},
		{"nonPositiveInteger", "int64"},
		{"negativeInteger", "int64"},
		{"long", "int64"},
		{"int", "int32"},
		{"short", "int16"},
		{"byte", "int8"},
		{"nonNegativeInteger", "uint64"},
		{"unsignedLong", "uint64"},
		{"unsignedInt", "uint32"},
		{"unsignedShort", "uint16"},
		{"unsignedByte", "uint8"},
		{"positiveInteger", "uint64"},

		{"anyType", "interface{}"},

		// PLC Open IEC 61131-3 elementary types - these are empty elements in choice
		{"BOOL", "*struct{}"},
		{"BYTE", "*struct{}"},
		{"WORD", "*struct{}"},
		{"DWORD", "*struct{}"},
		{"LWORD", "*struct{}"},
		{"SINT", "*struct{}"},
		{"INT", "*struct{}"},
		{"DINT", "*struct{}"},
		{"LINT", "*struct{}"},
		{"USINT", "*struct{}"},
		{"UINT", "*struct{}"},
		{"UDINT", "*struct{}"},
		{"ULINT", "*struct{}"},
		{"REAL", "*struct{}"},
		{"LREAL", "*struct{}"},
		{"TIME", "*struct{}"},
		{"DATE", "*struct{}"},
		{"DT", "*struct{}"},
		{"TOD", "*struct{}"},
	}
}

// ToGoTypeName converts an XSD type name to a valid Go type name
func ToGoTypeName(xsdType string) string {
	// Remove namespace prefix
	if colonIndex := strings.LastIndex(xsdType, ":"); colonIndex != -1 {
		xsdType = xsdType[colonIndex+1:]
	}

	// Convert to PascalCase
	return ToPascalCase(xsdType)
}

// ToGoFieldName converts an XSD element/attribute name to a valid Go field name
func ToGoFieldName(name string) string {
	return ToPascalCase(name)
}

// ToPascalCase converts a string to PascalCase by making the first letter uppercase
func ToPascalCase(s string) string {
	if s == "" {
		return s
	}

	// Simply capitalize the first letter, keep the rest as-is
	return strings.ToUpper(s[:1]) + s[1:]
}

// ToSnakeCase returns the original string as-is for JSON tags
// No case conversion needed - use XSD original text directly
func ToSnakeCase(s string) string {
	return s
}

// ParseOccurs parses minOccurs and maxOccurs attributes
func ParseOccurs(minOccurs, maxOccurs string) (min int, max int) {
	min = 1
	max = 1

	if minOccurs != "" {
		if minOccurs == "0" {
			min = 0
		} else {
			min = 1
		}
	}

	if maxOccurs != "" {
		if maxOccurs == "unbounded" {
			max = -1
		} else if maxOccurs == "0" {
			max = 0
		} else {
			max = 1
		}
	}

	return min, max
}

// IsOptional determines if a field is optional based on minOccurs
func IsOptional(minOccurs string) bool {
	return minOccurs == "0"
}

// IsArray determines if a field should be an array based on maxOccurs
func IsArray(maxOccurs string) bool {
	return maxOccurs == "unbounded" || (maxOccurs != "" && maxOccurs != "0" && maxOccurs != "1")
}

// GetDocumentation extracts documentation from XSD annotation
func GetDocumentation(annotation *XSDAnnotation) string {
	if annotation == nil {
		return ""
	}

	var docs []string
	for _, doc := range annotation.Documentation {
		if content := strings.TrimSpace(doc.Content); content != "" {
			docs = append(docs, content)
		}
	}

	return strings.Join(docs, " ")
}

// ValidateGoIdentifier checks if a string is a valid Go identifier
func ValidateGoIdentifier(name string) error {
	if name == "" {
		return fmt.Errorf("identifier cannot be empty")
	}

	// Check first character
	first := rune(name[0])
	if !((first >= 'a' && first <= 'z') || (first >= 'A' && first <= 'Z') || first == '_') {
		return fmt.Errorf("identifier must start with letter or underscore: %s", name)
	}

	// Check remaining characters
	for _, r := range name[1:] {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_') {
			return fmt.Errorf("invalid character in identifier: %s", name)
		}
	}

	return nil
}
