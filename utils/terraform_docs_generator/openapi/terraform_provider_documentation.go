package openapi

type TerraformProviderDocumentation struct {
	ProviderName          string
	ProviderInstallation  ProviderInstallation
	ProviderConfiguration ProviderConfiguration
	ProviderResources     ProviderResources
	DataSources           DataSources
}

type ProviderInstallation struct {
	Description string
	Example     string
	Other       string
}

type ProviderConfiguration struct {
	ExampleUsage       []ExampleUsage
	ArgumentsReference ArgumentsReference
}

type ProviderResources struct {
	Resources []Resource
}

type DataSources struct {
	DataSources []DataSource
}

type DataSource struct {
	Name               string
	Description        string
	ExampleUsage       []ExampleUsage
	ArgumentsReference ArgumentsReference
}

type Resource struct {
	Name                string
	Description         string
	ExampleUsage        []ExampleUsage
	ArgumentsReference  ArgumentsReference
	AttributesReference AttributesReference
	Import              Import
}

type ExampleUsage struct {
	Description string
	Example     string
}

type ArgumentsReference struct {
	Description string
	Properties  []Property
	Notes       []string
}

type AttributesReference struct {
	Description string
	Properties  []Property
	Notes       []string
}

type Import struct {
	Description string
	Example     string
	Notes       []string
}

type Property struct {
	Name        string
	Type        string
	Required    bool
	Description string
	Schema      []Property // This is used to describe the schema for array of objects or object properties
}

func (t TerraformProviderDocumentation) renderMarkup() {

}