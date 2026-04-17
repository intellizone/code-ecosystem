package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceUrl       string
	RequiredServices []ServiceName
	ServiceUpdateUrl string
	HeartBeatURL     string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	TeacherPortal  = ServiceName("TeacherPortal")
)

type patchEntry struct {
	Name ServiceName
	Url  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
