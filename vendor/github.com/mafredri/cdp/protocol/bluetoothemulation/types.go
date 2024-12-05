// Code generated by cdpgen. DO NOT EDIT.

package bluetoothemulation

// CentralState Indicates the various states of Central.
type CentralState string

// CentralState as enums.
const (
	CentralStateNotSet     CentralState = ""
	CentralStateAbsent     CentralState = "absent"
	CentralStatePoweredOff CentralState = "powered-off"
	CentralStatePoweredOn  CentralState = "powered-on"
)

func (e CentralState) Valid() bool {
	switch e {
	case "absent", "powered-off", "powered-on":
		return true
	default:
		return false
	}
}

func (e CentralState) String() string {
	return string(e)
}

// ManufacturerData Stores the manufacturer data
type ManufacturerData struct {
	Key  int    `json:"key"`  // Company identifier https://bitbucket.org/bluetooth-SIG/public/src/main/assigned_numbers/company_identifiers/company_identifiers.yaml https://usb.org/developers
	Data []byte `json:"data"` // Manufacturer-specific data (Encoded as a base64 string when passed over JSON)
}

// ScanRecord Stores the byte data of the advertisement packet sent by a
// Bluetooth device.
type ScanRecord struct {
	Name             *string            `json:"name,omitempty"`             // No description.
	UUIDs            []string           `json:"uuids,omitempty"`            // No description.
	Appearance       *int               `json:"appearance,omitempty"`       // Stores the external appearance description of the device.
	TxPower          *int               `json:"txPower,omitempty"`          // Stores the transmission power of a broadcasting device.
	ManufacturerData []ManufacturerData `json:"manufacturerData,omitempty"` // Key is the company identifier and the value is an array of bytes of manufacturer specific data.
}

// ScanEntry Stores the advertisement packet information that is sent by a
// Bluetooth device.
type ScanEntry struct {
	DeviceAddress string     `json:"deviceAddress"` // No description.
	RSSI          int        `json:"rssi"`          // No description.
	ScanRecord    ScanRecord `json:"scanRecord"`    // No description.
}
