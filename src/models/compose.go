package models

type Compose struct {
	Services map[string]ComposeService
	Volumes  map[string]ComposeVolumeDeclaration
	Networks map[string]ComposeNetworkDeclaration
}

type ComposeNetworkDeclaration struct {
}

type ComposeVolumeDeclaration struct {
}

type ComposeService struct {
	Image       string
	Ports       []string
	Networks    []string
	Volumes     []string
	Environment map[string]string
}
