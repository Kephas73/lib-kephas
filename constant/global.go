package constant

// Device
const (
	KDeviceKindPC     int = 1
	KDeviceKindMobile int = 2
	KDeviceKindWeb    int = 3
)

const (
	ValueEmpty int    = 0
	StrEmpty   string = ""
)

const (
	ModeTrimLeadingSpace bool   = true // Ignore whitespace
	ModeComma            string = ","  // Separator character
	ModeLazyQuotes       bool   = true // If LazyQuotes is true, a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
	ModeFieldsPerRecord  int    = 0    // Read sets it to the number of fields in the first record
)

const (
	AccessUUID string = "AccessUUID"
	DeviceKind string = "DeviceKind"
	DeviceIP   string = "DeviceIP"
	UserID     string = "UserID"
)
