package dokku

import "fmt"

const (
	gitAllowHostCmd       = "git:allow-host <host>"
	gitAuthCmd            = "git:auth <host> [<username> <password>]"
	gitFromArchiveCmd     = "git:from-archive <app> <archive-url> [<git-username> <git-email>]"
	gitFromImageCmd       = "git:from-image <app> <docker-image> [<git-username> <git-email>]"
	gitInitializeCmd      = "git:initialize <app>"
	gitPublicKeyCmd       = "git:public-key"
	gitReportCmd          = "git:report [<app>] [<flag>]"
	gitSetCmd             = "git:set <app> <property> (<value>)"
	gitSyncCmd            = "git:sync %s %s"
	gitSyncWithOptionsCmd = "git:sync %s %s %s %s"
	gitUnlockCmd          = "git:unlock <app> [--force]"
)

type GitSyncOptions struct {
	Build  bool
	GitRef string
}

func (c *DefaultClient) GitSyncAppRepo(appName string, repo string, opt *GitSyncOptions) error {
	cmd := fmt.Sprintf(gitSyncCmd, appName, repo)
	if opt != nil {
		var buildFlag string
		if opt.Build {
			buildFlag = "--build"
		}
		cmd = fmt.Sprintf(gitSyncWithOptionsCmd, buildFlag, appName, repo, opt.GitRef)
	}
	_, err := c.Exec(cmd)
	return err
}
