package instance

type Instance struct {
	DisplayName string
	Path        string
	Version     string
	Id          string // should probably be a uuid
}
