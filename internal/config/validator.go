package config

import (
	"errors"

	"github.com/koddr/gosl"
)

// Validate ...
func (c *Config) Validate() error {
	//
	var (
		// Project section errors.
		notSetProjectName = errors.New("project name is not set")

		// Backend section errors.
		notSetBackendTemplateOrRepository     = errors.New("'template' name or 'repository' settings in the 'backend' section is not set")
		notFoundBackendTemplateName           = errors.New("'template' name in the 'backend' section is not found in the supported list")
		collisionBackendTemplateAndRepository = errors.New("choose one: 'template' name or 'repository' settings in the 'backend' section")

		// Frontend section errors.
		notSetFrontendTemplateOrRepository     = errors.New("'template' name or 'repository' settings in the 'frontend' section is not set")
		notFoundFrontendTemplateName           = errors.New("'template' name in the 'frontend' section is not found in the supported list")
		collisionFrontendTemplateAndRepository = errors.New("choose one: 'template' name or 'repository' settings in the 'frontend' section")

		// Proxy section errors.
		notSetProxyTemplateOrRepository     = errors.New("'template' name or 'repository' settings in the 'proxy' section is not set")
		notFoundProxyTemplateName           = errors.New("'template' name in the 'proxy' section is not found in the supported list")
		collisionProxyTemplateAndRepository = errors.New("choose one: 'template' name or 'repository' settings in the 'proxy' section")

		// Containers section errors.
		notSetContainers = errors.New("'containers' section is not set, but 'deploy' section was configured")

		// Deploy section errors.
		notSetDeployAnsible = errors.New("'ansible' settings in the 'deploy' section is not set")
		notSetDeployDocker  = errors.New("'docker' settings in the 'deploy' section is not set")
	)

	// Create a new slice for join errors.
	errs := make([]error, 0)

	// Validate 'project' section.
	if c.Project.Name == "" {
		errs = append(errs, notSetProjectName)
	}

	// Validate 'backend' section.
	if c.Backend != nil {
		// Check, if 'backend' template name or repository is set.
		if c.Backend.Name == "" && c.Backend.Repository == nil {
			errs = append(errs, notSetBackendTemplateOrRepository)
		}

		// Check, if 'backend' template name or repository is not set both.
		if c.Backend.Name != "" && c.Backend.Repository != nil {
			errs = append(errs, collisionBackendTemplateAndRepository)
		}

		// Check, if 'backend' template name was set and not contains in the templates index.
		if c.Backend.Name != "" && !gosl.ContainsInSlice(BackendTemplates, c.Backend.Name) {
			errs = append(errs, notFoundBackendTemplateName)
		}
	}

	// Validate 'frontend' section.
	if c.Frontend != nil {
		// Check, if 'frontend' template name or repository is set.
		if c.Frontend.Name == "" && c.Frontend.Repository == nil {
			errs = append(errs, notSetFrontendTemplateOrRepository)
		}

		// Check, if 'frontend' template name or repository is not set both.
		if c.Frontend.Name != "" && c.Frontend.Repository != nil {
			errs = append(errs, collisionFrontendTemplateAndRepository)
		}

		// Check, if 'frontend' template name was set and not contains in the templates index.
		if c.Frontend.Name != "" && !gosl.ContainsInSlice(FrontendTemplates, c.Frontend.Name) {
			errs = append(errs, notFoundFrontendTemplateName)
		}
	}

	// Validate 'proxy' section.
	if c.Proxy != nil {
		// Check, if 'proxy' template name or repository is set.
		if c.Proxy.Name == "" && c.Proxy.Repository == nil {
			errs = append(errs, notSetProxyTemplateOrRepository)
		}

		// Check, if 'proxy' template name or repository is not set both.
		if c.Proxy.Name != "" && c.Proxy.Repository != nil {
			errs = append(errs, collisionProxyTemplateAndRepository)
		}

		// Check, if 'proxy' template name was set and not contains in the templates index.
		if c.Proxy.Name != "" && !gosl.ContainsInSlice(ProxyTemplates, c.Proxy.Name) {
			errs = append(errs, notFoundProxyTemplateName)
		}
	}

	// Validate 'deploy' section.
	if c.Deploy != nil {
		// Check, if 'containers' section is set.
		if c.Containers == nil {
			errs = append(errs, notSetContainers)
		}

		// Check, if 'ansible' settings in the 'deploy' section is set.
		if c.Deploy.Ansible == nil {
			errs = append(errs, notSetDeployAnsible)
		}

		// Check, if 'docker' settings in the 'deploy' section is set.
		if c.Deploy.Docker == nil {
			errs = append(errs, notSetDeployDocker)
		}
	}

	return errors.Join(errs...)
}
