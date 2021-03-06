package organization

import (
	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/api/organizations"
	"github.com/cloudfoundry/cli/cf/command_metadata"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	. "github.com/cloudfoundry/cli/cf/i18n"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/codegangsta/cli"
)

type UnsharePrivateDomain struct {
	ui         terminal.UI
	config     core_config.Reader
	orgRepo    organizations.OrganizationRepository
	domainRepo api.DomainRepository
	orgReq     requirements.OrganizationRequirement
}

func NewUnsharePrivateDomain(ui terminal.UI, config core_config.Reader,
	orgRepo organizations.OrganizationRepository, domainRepo api.DomainRepository) *UnsharePrivateDomain {
	return &UnsharePrivateDomain{
		ui:         ui,
		config:     config,
		orgRepo:    orgRepo,
		domainRepo: domainRepo,
	}
}

func (cmd *UnsharePrivateDomain) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "unshare-private-domain",
		Description: T("Unshare a private domain with an org"),
		Usage:       T("CF_NAME unshare-private-domain ORG DOMAIN"),
	}
}

func (cmd *UnsharePrivateDomain) GetRequirements(requirementsFactory requirements.Factory, c *cli.Context) ([]requirements.Requirement, error) {
	if len(c.Args()) != 2 {
		cmd.ui.FailWithUsage(c)
	}

	cmd.orgReq = requirementsFactory.NewOrganizationRequirement(c.Args()[0])

	return []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
		cmd.orgReq,
	}, nil
}

func (cmd *UnsharePrivateDomain) Run(c *cli.Context) {
	org := cmd.orgReq.GetOrganization()
	domainName := c.Args()[1]
	domain, err := cmd.domainRepo.FindPrivateByName(domainName)

	if err != nil {
		cmd.ui.Failed(err.Error())
		return
	}

	cmd.ui.Say(T("Unsharing domain {{.DomainName}} from org {{.OrgName}} as {{.Username}}...",
		map[string]interface{}{
			"DomainName": terminal.EntityNameColor(domain.Name),
			"OrgName":    terminal.EntityNameColor(org.Name),
			"Username":   terminal.EntityNameColor(cmd.config.Username())}))

	err = cmd.orgRepo.UnsharePrivateDomain(org.Guid, domain.Guid)
	if err != nil {
		cmd.ui.Failed(err.Error())
		return
	}

	cmd.ui.Ok()
}
