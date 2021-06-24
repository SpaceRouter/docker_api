package models

type Compose struct {
	Services map[string]ComposeService
	Volumes  map[string]ComposeVolumeDeclaration
	Networks map[string]ComposeNetworkDeclaration
}

type ComposeNetworkDeclaration struct {
	External bool
}

type ComposeVolumeDeclaration struct {
}

type ComposeService struct {
	Image       string
	Networks    []string
	Volumes     []string
	Labels      map[string]string
	Environment map[string]string
}

type ComposeOutput struct {
	Name     string
	Services []interface{}
	Owner    string
}
