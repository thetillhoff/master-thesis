package hardware_inspection

// This type is used to mirror the actual physical servers.
type Machine struct {
	MacAddress string // colon seperated
	IpAddress  string

	Ram   uint64
	Cores int
}
