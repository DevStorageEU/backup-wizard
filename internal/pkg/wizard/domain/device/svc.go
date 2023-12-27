package device

type InspectionService interface {
	Inspect(ips []string) (*Inspection, error)
}
