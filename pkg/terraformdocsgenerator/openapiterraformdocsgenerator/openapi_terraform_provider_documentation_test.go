package openapiterraformdocsgenerator

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"text/template"
)

func TestTerraformProviderDocumentation_RenderZendeskHTML(t *testing.T) {
	providerName := "openapi"
	terraformProviderDocumentation := TerraformProviderDocumentation{
		ProviderName:  providerName,
		ProviderNotes: []string{"provider custom note 1", "provider custom note 2"},
		ProviderInstallation: ProviderInstallation{
			ProviderName: providerName,
			Hostname:     "terraform.example.com",
			Namespace:    "examplecorp",
		},
		ProviderConfiguration: ProviderConfiguration{
			ProviderName: providerName,
		},
		ShowSpecialTermsDefinitions: true,
		ProviderResources: ProviderResources{
			ProviderName: providerName,
			Resources:    []Resource{},
		},
		DataSources: DataSources{
			ProviderName:        providerName,
			DataSourceInstances: []DataSource{},
			DataSources:         []DataSource{},
		},
	}
	var buf bytes.Buffer
	expectedHTML := `<p dir="ltr">
  This guide lists the configuration for 'openapi' Terraform provider
  resources that can be managed using
  <a href="https://www.hashicorp.com/blog/announcing-terraform-0-12/" target="_self">Terraform v0.12</a>.&nbsp;
</p>
<p><span class="wysiwyg-color-red">*Note: provider custom note 1</span></p>
<p><span class="wysiwyg-color-red">*Note: provider custom note 2</span></p>
<ul>
  <li>
    <a href="#provider_installation" target="_self">Provider Installation</a>
  </li>
  
    <li>
        <a href="#provider_resources" target="_self">Provider Resources</a>
        <ul>
            </ul>
    </li>
    <li>
        <a href="#provider_datasources" target="_self">Data Sources (using resource id)</a>
        <ul>
            </ul>
    </li>
    <li>
        <a href="#provider_datasources_filters" target="_self">Data Sources (using filters)</a>
        <ul>
            </ul>
    </li>


  <li>
    <a href="#special_terms_definitions" target="_self">Special Terms Definitions</a>
    <ul>

    </ul>
  </li>

</ul><h2 id="provider_installation">Provider Installation</h2>
<p>
  In order to provision 'openapi' Terraform resources, you need to first install the 'openapi'
  Terraform plugin by running&nbsp;the following command (you must be running Terraform &gt;= 0.12):
</p>
<pre></pre>
<p>
  <span></span>
</p>
<pre dir="ltr">
➜ ~ terraform init &amp;&amp; terraform plan
</pre>
<p>
<b>Note:</b> As of Terraform &gt;= 0.13 each Terraform module must declare which providers it requires, so that Terraform can install and use them. If you are using Terraform &gt;= 0.13, copy into your .tf file the 
following snippet already populated with the provider configuration: 
<pre dir="ltr">
terraform {
  required_providers {
    openapi = {
      source  = "terraform.example.com/examplecorp/openapi"
      version = ">= 2.0.1" 
    }
  }
}
</pre>
</p>
<h2 id="provider_resources">Provider Resources</h2>
	
<p>No resources are supported at the moment.</p> <h2 id="provider_datasources">Data Sources (using resource id)</h2>

No data sources using resource id are supported at the moment. 

<h2 id="provider_datasources_filters">Data Sources (using filters)</h2>

No data sources using filters are supported at the moment. 
<h2 id="special_terms_definitions">Special Terms Definitions</h2>
<p>
  This section describes specific terms used throughout this document to clarify their meaning in the context of Terraform.
</p>`
	err := terraformProviderDocumentation.RenderHTML(&buf)
	fmt.Println(buf.String())
	assert.Equal(t, expectedHTML, strings.Trim(buf.String(), "\n"))
	assert.Nil(t, err)
}

func TestTerraformProviderDocumentation_RenderZendeskHTML_Errors(t *testing.T) {
	testCases := []struct {
		name                            string
		tableOfContentsTemplate         string
		providerInstallTemplate         string
		providerConfigTemplate          string
		providerResourcesTemplate       string
		providerDataSourcesTemplate     string
		specialTermsDefinitionsTemplate string
		expectedErr                     error
	}{
		{
			name:                    "provider installation template error",
			providerInstallTemplate: `{{.nonExistentVariable}}`,
			expectedErr:             errors.New("template: ProviderInstallation:1:2: executing \"ProviderInstallation\" at <.nonExistentVariable>: can't evaluate field nonExistentVariable in type openapiterraformdocsgenerator.ProviderInstallation"),
		},
		{
			name:                   "provider configuration template error",
			providerConfigTemplate: `{{.nonExistentVariable}}`,
			expectedErr:            errors.New("template: ProviderConfiguration:1:2: executing \"ProviderConfiguration\" at <.nonExistentVariable>: can't evaluate field nonExistentVariable in type openapiterraformdocsgenerator.ProviderConfiguration"),
		},
		{
			name:                      "provider resources template error",
			providerResourcesTemplate: `{{.nonExistentVariable}}`,
			expectedErr:               errors.New("template: ProviderResources:1:2: executing \"ProviderResources\" at <.nonExistentVariable>: can't evaluate field nonExistentVariable in type openapiterraformdocsgenerator.ProviderResources"),
		},
		{
			name:                        "data sources template error",
			providerDataSourcesTemplate: `{{.nonExistentVariable}}`,
			expectedErr:                 errors.New("template: DataSources:1:2: executing \"DataSources\" at <.nonExistentVariable>: can't evaluate field nonExistentVariable in type openapiterraformdocsgenerator.DataSources"),
		},
		{
			name:                    "table of contents template error",
			tableOfContentsTemplate: `{{.nonExistentVariable}}`,
			expectedErr:             errors.New("template: TerraformProviderDocTableOfContents:1:2: executing \"TerraformProviderDocTableOfContents\" at <.nonExistentVariable>: can't evaluate field nonExistentVariable in type openapiterraformdocsgenerator.TerraformProviderDocumentation"),
		},
		{
			name:                            "special terms definitions template error",
			specialTermsDefinitionsTemplate: `{{.nonExistentVariable}}`,
			expectedErr:                     errors.New("template: TerraformProviderDocSpecialTermsDefinitions:1:2: executing \"TerraformProviderDocSpecialTermsDefinitions\" at <.nonExistentVariable>: can't evaluate field nonExistentVariable in type openapiterraformdocsgenerator.TerraformProviderDocumentation"),
		},
	}

	for _, tc := range testCases {
		var buf bytes.Buffer
		d := TerraformProviderDocumentation{}
		err := d.renderZendeskHTML(&buf, tc.tableOfContentsTemplate, tc.providerInstallTemplate, tc.providerConfigTemplate, tc.providerResourcesTemplate, tc.providerDataSourcesTemplate, tc.specialTermsDefinitionsTemplate)
		templateErr := err.(template.ExecError)
		assert.EqualError(t, templateErr.Err, tc.expectedErr.Error())
	}
}
