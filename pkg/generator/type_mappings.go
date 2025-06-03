package generator

// CommonTypeMapping represents a universal type mapping across all supported languages
type CommonTypeMapping struct {
	XSDType    string
	GoType     string
	JavaType   string
	CSharpType string
	PythonType string
	Comments   string // Documentation for this type mapping
}

// CommonTypeMappingRegistry holds all universal type mappings
type CommonTypeMappingRegistry struct {
	BuiltinMappings []CommonTypeMapping
	CustomMappings  []CommonTypeMapping
}

// NewCommonTypeMappingRegistry creates a new registry with predefined mappings
func NewCommonTypeMappingRegistry() *CommonTypeMappingRegistry {
	return &CommonTypeMappingRegistry{
		BuiltinMappings: getBuiltinTypeMappings(),
		CustomMappings:  getCustomTypeMappings(),
	}
}

// GetMappingsForLanguage returns all type mappings for a specific language
func (r *CommonTypeMappingRegistry) GetMappingsForLanguage(lang TargetLanguage) []TypeMapping {
	var mappings []TypeMapping

	// Add builtin mappings
	for _, mapping := range r.BuiltinMappings {
		targetType := r.getTargetTypeForLanguage(mapping, lang)
		if targetType != "" {
			mappings = append(mappings, TypeMapping{
				XSDType:    mapping.XSDType,
				TargetType: targetType,
			})
		}
	}

	// Add custom mappings
	for _, mapping := range r.CustomMappings {
		targetType := r.getTargetTypeForLanguage(mapping, lang)
		if targetType != "" {
			mappings = append(mappings, TypeMapping{
				XSDType:    mapping.XSDType,
				TargetType: targetType,
			})
		}
	}

	return mappings
}

// getTargetTypeForLanguage returns the target type for a specific language
func (r *CommonTypeMappingRegistry) getTargetTypeForLanguage(mapping CommonTypeMapping, lang TargetLanguage) string {
	switch lang {
	case LanguageGo:
		return mapping.GoType
	case LanguageJava:
		return mapping.JavaType
	case LanguageCSharp:
		return mapping.CSharpType
	case LanguagePython:
		return mapping.PythonType
	default:
		return ""
	}
}

// getBuiltinTypeMappings returns the universal builtin type mappings
func getBuiltinTypeMappings() []CommonTypeMapping {
	return []CommonTypeMapping{
		// String types
		{
			XSDType:    "string",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "Basic string type",
		},
		{
			XSDType:    "normalizedString",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "Normalized string (whitespace collapsed)",
		},
		{
			XSDType:    "token",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "Token string (no leading/trailing whitespace)",
		},
		{
			XSDType:    "anyURI",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "URI string",
		},
		{
			XSDType:    "language",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "Language identifier",
		},
		{
			XSDType:    "NMTOKEN",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "Name token",
		},
		{
			XSDType:    "Name",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "XML Name",
		},
		{
			XSDType:    "NCName",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "Non-colonized name",
		},
		{
			XSDType:    "ID",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "ID attribute type",
		},
		{
			XSDType:    "IDREF",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "ID reference",
		},
		{
			XSDType:    "ENTITY",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "Entity reference",
		},
		{
			XSDType:    "QName",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "Qualified name",
		},

		// Boolean type
		{
			XSDType:    "boolean",
			GoType:     "bool",
			JavaType:   "Boolean",
			CSharpType: "bool",
			PythonType: "bool",
			Comments:   "Boolean true/false value",
		},

		// Numeric types
		{
			XSDType:    "decimal",
			GoType:     "float64",
			JavaType:   "BigDecimal",
			CSharpType: "decimal",
			PythonType: "float",
			Comments:   "Decimal number with arbitrary precision",
		},
		{
			XSDType:    "float",
			GoType:     "float32",
			JavaType:   "Float",
			CSharpType: "float",
			PythonType: "float",
			Comments:   "Single precision floating point",
		},
		{
			XSDType:    "double",
			GoType:     "float64",
			JavaType:   "Double",
			CSharpType: "double",
			PythonType: "float",
			Comments:   "Double precision floating point",
		},

		// Integer types
		{
			XSDType:    "int",
			GoType:     "int32",
			JavaType:   "Integer",
			CSharpType: "int",
			PythonType: "int",
			Comments:   "32-bit signed integer",
		},
		{
			XSDType:    "integer",
			GoType:     "int64",
			JavaType:   "BigInteger",
			CSharpType: "long",
			PythonType: "int",
			Comments:   "Arbitrary precision integer",
		},
		{
			XSDType:    "long",
			GoType:     "int64",
			JavaType:   "Long",
			CSharpType: "long",
			PythonType: "int",
			Comments:   "64-bit signed integer",
		},
		{
			XSDType:    "short",
			GoType:     "int16",
			JavaType:   "Short",
			CSharpType: "short",
			PythonType: "int",
			Comments:   "16-bit signed integer",
		},
		{
			XSDType:    "byte",
			GoType:     "int8",
			JavaType:   "Byte",
			CSharpType: "sbyte",
			PythonType: "int",
			Comments:   "8-bit signed integer",
		},

		// Unsigned integer types
		{
			XSDType:    "unsignedLong",
			GoType:     "uint64",
			JavaType:   "BigInteger",
			CSharpType: "ulong",
			PythonType: "int",
			Comments:   "64-bit unsigned integer",
		},
		{
			XSDType:    "unsignedInt",
			GoType:     "uint32",
			JavaType:   "Long",
			CSharpType: "uint",
			PythonType: "int",
			Comments:   "32-bit unsigned integer",
		},
		{
			XSDType:    "unsignedShort",
			GoType:     "uint16",
			JavaType:   "Integer",
			CSharpType: "ushort",
			PythonType: "int",
			Comments:   "16-bit unsigned integer",
		},
		{
			XSDType:    "unsignedByte",
			GoType:     "uint8",
			JavaType:   "Short",
			CSharpType: "byte",
			PythonType: "int",
			Comments:   "8-bit unsigned integer",
		},
		{
			XSDType:    "nonNegativeInteger",
			GoType:     "uint64",
			JavaType:   "BigInteger",
			CSharpType: "ulong",
			PythonType: "int",
			Comments:   "Non-negative integer",
		},
		{
			XSDType:    "positiveInteger",
			GoType:     "uint64",
			JavaType:   "BigInteger",
			CSharpType: "ulong",
			PythonType: "int",
			Comments:   "Positive integer",
		},
		{
			XSDType:    "nonPositiveInteger",
			GoType:     "int64",
			JavaType:   "BigInteger",
			CSharpType: "long",
			PythonType: "int",
			Comments:   "Non-positive integer",
		},
		{
			XSDType:    "negativeInteger",
			GoType:     "int64",
			JavaType:   "BigInteger",
			CSharpType: "long",
			PythonType: "int",
			Comments:   "Negative integer",
		},

		// Date and time types
		{
			XSDType:    "dateTime",
			GoType:     "time.Time",
			JavaType:   "LocalDateTime",
			CSharpType: "DateTime",
			PythonType: "datetime",
			Comments:   "Date and time instant",
		},
		{
			XSDType:    "date",
			GoType:     "string",
			JavaType:   "LocalDate",
			CSharpType: "DateTime",
			PythonType: "date",
			Comments:   "Date without time",
		},
		{
			XSDType:    "time",
			GoType:     "string",
			JavaType:   "LocalTime",
			CSharpType: "TimeSpan",
			PythonType: "time",
			Comments:   "Time without date",
		},
		{
			XSDType:    "duration",
			GoType:     "string",
			JavaType:   "Duration",
			CSharpType: "TimeSpan",
			PythonType: "timedelta",
			Comments:   "Time duration",
		},
		{
			XSDType:    "gYearMonth",
			GoType:     "string",
			JavaType:   "YearMonth",
			CSharpType: "DateTime",
			PythonType: "str",
			Comments:   "Year and month",
		},
		{
			XSDType:    "gYear",
			GoType:     "string",
			JavaType:   "Year",
			CSharpType: "DateTime",
			PythonType: "str",
			Comments:   "Year",
		},
		{
			XSDType:    "gMonthDay",
			GoType:     "string",
			JavaType:   "MonthDay",
			CSharpType: "DateTime",
			PythonType: "str",
			Comments:   "Month and day",
		},
		{
			XSDType:    "gDay",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "Day",
		},
		{
			XSDType:    "gMonth",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "Month",
		},

		// Binary types
		{
			XSDType:    "base64Binary",
			GoType:     "[]byte",
			JavaType:   "byte[]",
			CSharpType: "byte[]",
			PythonType: "bytes",
			Comments:   "Base64 encoded binary data",
		},
		{
			XSDType:    "hexBinary",
			GoType:     "[]byte",
			JavaType:   "byte[]",
			CSharpType: "byte[]",
			PythonType: "bytes",
			Comments:   "Hex encoded binary data",
		},

		// Collection types
		{
			XSDType:    "NMTOKENS",
			GoType:     "[]string",
			JavaType:   "List<String>",
			CSharpType: "List<string>",
			PythonType: "List[str]",
			Comments:   "List of name tokens",
		},
		{
			XSDType:    "IDREFS",
			GoType:     "[]string",
			JavaType:   "List<String>",
			CSharpType: "List<string>",
			PythonType: "List[str]",
			Comments:   "List of ID references",
		},
		{
			XSDType:    "ENTITIES",
			GoType:     "[]string",
			JavaType:   "List<String>",
			CSharpType: "List<string>",
			PythonType: "List[str]",
			Comments:   "List of entities",
		},

		// Special types
		{
			XSDType:    "anyType",
			GoType:     "interface{}",
			JavaType:   "Object",
			CSharpType: "object",
			PythonType: "Any",
			Comments:   "Any type (generic object)",
		},
	}
}

// getCustomTypeMappings returns PLC/industrial type mappings
func getCustomTypeMappings() []CommonTypeMapping {
	return []CommonTypeMapping{
		// PLC Boolean types
		{
			XSDType:    "BOOL",
			GoType:     "bool",
			JavaType:   "Boolean",
			CSharpType: "bool",
			PythonType: "bool",
			Comments:   "PLC Boolean type",
		},

		// PLC Integer types
		{
			XSDType:    "SINT",
			GoType:     "int8",
			JavaType:   "Byte",
			CSharpType: "sbyte",
			PythonType: "int",
			Comments:   "PLC Small signed integer (-128 to 127)",
		},
		{
			XSDType:    "INT",
			GoType:     "int16",
			JavaType:   "Short",
			CSharpType: "short",
			PythonType: "int",
			Comments:   "PLC Signed integer (-32768 to 32767)",
		},
		{
			XSDType:    "DINT",
			GoType:     "int32",
			JavaType:   "Integer",
			CSharpType: "int",
			PythonType: "int",
			Comments:   "PLC Double signed integer (-2147483648 to 2147483647)",
		},
		{
			XSDType:    "LINT",
			GoType:     "int64",
			JavaType:   "Long",
			CSharpType: "long",
			PythonType: "int",
			Comments:   "PLC Long signed integer",
		},

		// PLC Unsigned integer types
		{
			XSDType:    "USINT",
			GoType:     "uint8",
			JavaType:   "Short",
			CSharpType: "byte",
			PythonType: "int",
			Comments:   "PLC Unsigned small integer (0-255)",
		},
		{
			XSDType:    "UINT",
			GoType:     "uint16",
			JavaType:   "Integer",
			CSharpType: "ushort",
			PythonType: "int",
			Comments:   "PLC Unsigned integer (0-65535)",
		},
		{
			XSDType:    "UDINT",
			GoType:     "uint32",
			JavaType:   "Long",
			CSharpType: "uint",
			PythonType: "int",
			Comments:   "PLC Unsigned double integer (0-4294967295)",
		},
		{
			XSDType:    "ULINT",
			GoType:     "uint64",
			JavaType:   "BigInteger",
			CSharpType: "ulong",
			PythonType: "int",
			Comments:   "PLC Unsigned long integer",
		},

		// PLC Bit string types
		{
			XSDType:    "BYTE",
			GoType:     "uint8",
			JavaType:   "Byte",
			CSharpType: "byte",
			PythonType: "int",
			Comments:   "PLC 8-bit string (0-255)",
		},
		{
			XSDType:    "WORD",
			GoType:     "uint16",
			JavaType:   "Integer",
			CSharpType: "ushort",
			PythonType: "int",
			Comments:   "PLC 16-bit string (0-65535)",
		},
		{
			XSDType:    "DWORD",
			GoType:     "uint32",
			JavaType:   "Long",
			CSharpType: "uint",
			PythonType: "int",
			Comments:   "PLC 32-bit string (0-4294967295)",
		},
		{
			XSDType:    "LWORD",
			GoType:     "uint64",
			JavaType:   "BigInteger",
			CSharpType: "ulong",
			PythonType: "int",
			Comments:   "PLC 64-bit string",
		},

		// PLC Floating point types
		{
			XSDType:    "REAL",
			GoType:     "float32",
			JavaType:   "Float",
			CSharpType: "float",
			PythonType: "float",
			Comments:   "PLC Single precision floating point",
		},
		{
			XSDType:    "LREAL",
			GoType:     "float64",
			JavaType:   "Double",
			CSharpType: "double",
			PythonType: "float",
			Comments:   "PLC Double precision floating point",
		},

		// PLC String types
		{
			XSDType:    "STRING",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "PLC String type",
		},
		{
			XSDType:    "WSTRING",
			GoType:     "string",
			JavaType:   "String",
			CSharpType: "string",
			PythonType: "str",
			Comments:   "PLC Wide string type",
		},

		// PLC Time types
		{
			XSDType:    "TIME",
			GoType:     "time.Duration",
			JavaType:   "Duration",
			CSharpType: "TimeSpan",
			PythonType: "timedelta",
			Comments:   "PLC Time duration",
		},
		{
			XSDType:    "LTIME",
			GoType:     "time.Duration",
			JavaType:   "Duration",
			CSharpType: "TimeSpan",
			PythonType: "timedelta",
			Comments:   "PLC Long time duration",
		},
		{
			XSDType:    "DATE",
			GoType:     "time.Time",
			JavaType:   "LocalDate",
			CSharpType: "DateTime",
			PythonType: "date",
			Comments:   "PLC Date",
		},
		{
			XSDType:    "TIME_OF_DAY",
			GoType:     "time.Time",
			JavaType:   "LocalTime",
			CSharpType: "TimeSpan",
			PythonType: "time",
			Comments:   "PLC Time of day",
		},
		{
			XSDType:    "TOD",
			GoType:     "time.Time",
			JavaType:   "LocalTime",
			CSharpType: "TimeSpan",
			PythonType: "time",
			Comments:   "PLC Time of day (short form)",
		},
		{
			XSDType:    "DATE_AND_TIME",
			GoType:     "time.Time",
			JavaType:   "LocalDateTime",
			CSharpType: "DateTime",
			PythonType: "datetime",
			Comments:   "PLC Date and time",
		},
		{
			XSDType:    "DT",
			GoType:     "time.Time",
			JavaType:   "LocalDateTime",
			CSharpType: "DateTime",
			PythonType: "datetime",
			Comments:   "PLC Date and time (short form)",
		},
	}
}
