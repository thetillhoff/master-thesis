package hardware_inspection

type lease struct {
	timestamp  uint64
	macAddress string
	ipAddress  string
}

var (
	// Contains a map of MAC to IP, where the MAC address acts as key
	// Separator of fields of MAC is colon!
	leaseTable map[string]lease = make(map[string]lease)
)
