package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
)

type gitManager interface {
	GitInitializeApp(appName string) error
	GitGetPublicKey() (string, error)
	GitSyncAppRepo(appName string, repo string, opt *GitSyncOptions) error
	GitCreateFromArchive(appName string, url string, opt *GitArchiveOptions) error
	GitCreateFromImage(appName string, image string, opt *GitImageOptions) error
	GitSetAuth(host string, username string, password string) error
	GitRemoveAuth(host string) error
	GitSetAppProperty(appName string, property GitProperty, val string) error
	GitRemoveAppProperty(appName string, property GitProperty) error
	GitAllowHost(host string) error
	GitUnlockApp(appName string, force bool) error
	GitGetAppReport(appName string) (*GitAppReport, error)
	GitGetReport() (GitReport, error)

	GitRunRepoGC(appName string) error
	GitPurgeRepoCache(appName string) error
}

type GitAppReport struct {
	DeployBranch       string `json:"deploy_branch" dokku:"Git deploy branch"`
	GlobalDeployBranch string `json:"global_deploy_branch" dokku:"Git global deploy branch"`
	KeepGitDir         bool   `json:"keep_git_dir" dokku:"Git keep git dir"`
	RevisionEnvVar     string `json:"rev_env_var" dokku:"Git rev env var"`
	SHA                string `json:"sha" dokku:"Git sha"`
	LastUpdatedAt      string `json:"last_updated_at" dokku:"Git last updated at"`
}

type GitReport map[string]*GitAppReport

type GitProperty string

const (
	GitPropertyDeployBranch = GitProperty("deploy-branch")
	GitPropertyRevEnvVar    = GitProperty("rev-env-var")
	GitPropertyKeepGitDir   = GitProperty("keep-git-dir")
)

const (
	gitAllowHostCmd       = "git:allow-host %s"
	gitAuthCmd            = "git:auth %s %s"
	gitFromArchiveCmd     = "git:from-archive --archive-type %s %s %s %s"
	gitFromImageCmd       = "git:from-image %s %s %s %s"
	gitInitializeCmd      = "git:initialize %s"
	gitPublicKeyCmd       = "git:public-key"
	gitReportCmd          = "git:report %s"
	gitSetCmd             = "git:set %s %s %s"
	gitSyncCmd            = "git:sync %s %s"
	gitSyncWithOptionsCmd = "git:sync %s %s %s %s"
	gitUnlockCmd          = "git:unlock %s %s"

	gitRepoGcCmd         = "repo:gc %s"
	gitRepoPurgeCacheCmd = "repo:purge-cache %s"
)

type GitArchiveOptions struct {
	ArchiveType   string
	AuthorDetails *GitAuthorDetails
}

type GitImageOptions struct {
	BuildDir      string
	AuthorDetails *GitAuthorDetails
}

type GitAuthorDetails struct {
	Username string
	Email    string
}

func (ad *GitAuthorDetails) String() string {
	return fmt.Sprintf("\"%s\" \"%s\"", ad.Username, ad.Email)
}

type GitSyncOptions struct {
	Build  bool
	GitRef string
}

func (c *BaseClient) GitInitializeApp(appName string) error {
	cmd := fmt.Sprintf(gitInitializeCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GitGetPublicKey() (string, error) {
	return c.Exec(gitPublicKeyCmd)
}

func (c *BaseClient) GitSyncAppRepo(appName string, repo string, opt *GitSyncOptions) error {
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

func (c *BaseClient) GitCreateFromArchive(appName string, url string, opt *GitArchiveOptions) error {
	var authorDetails string
	archiveType := "tar"
	if opt != nil {
		if opt.AuthorDetails != nil {
			authorDetails = opt.AuthorDetails.String()
		}
		if opt.ArchiveType != "" {
			archiveType = opt.ArchiveType
		}
	}
	cmd := fmt.Sprintf(gitFromArchiveCmd, archiveType, appName, url, authorDetails)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GitCreateFromImage(appName string, image string, opt *GitImageOptions) error {
	var authorDetails string
	buildDir := ""
	if opt != nil {
		if opt.AuthorDetails != nil {
			authorDetails = opt.AuthorDetails.String()
		}
		if opt.BuildDir != "" {
			buildDir = "--build-dir " + opt.BuildDir
		}
	}
	cmd := fmt.Sprintf(gitFromImageCmd, appName, image, buildDir, authorDetails)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GitSetAuth(host string, username string, password string) error {
	authDetails := fmt.Sprintf("%s %s", username, password)
	cmd := fmt.Sprintf(gitAuthCmd, host, authDetails)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GitRemoveAuth(host string) error {
	cmd := fmt.Sprintf(gitAuthCmd, host, "")
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GitSetAppProperty(appName string, property GitProperty, val string) error {
	cmd := fmt.Sprintf(gitSetCmd, appName, property, val)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GitRemoveAppProperty(appName string, property GitProperty) error {
	return c.GitSetAppProperty(appName, property, "")
}

func (c *BaseClient) GitAllowHost(host string) error {
	cmd := fmt.Sprintf(gitAllowHostCmd, host)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GitUnlockApp(appName string, force bool) error {
	var forceStr string
	if force {
		forceStr = "--force"
	}
	cmd := fmt.Sprintf(gitUnlockCmd, appName, forceStr)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GitGetAppReport(appName string) (*GitAppReport, error) {
	cmd := fmt.Sprintf(gitReportCmd, appName)
	output, err := c.Exec(cmd)

	var gitReport GitAppReport
	if err := reports.ParseInto(output, &gitReport); err != nil {
		return nil, fmt.Errorf("failed to parse report: %w", err)
	}

	return &gitReport, err
}

func (c *BaseClient) GitGetReport() (GitReport, error) {
	cmd := fmt.Sprintf(gitReportCmd, "")
	output, err := c.Exec(cmd)

	var gitReport GitReport
	if err := reports.ParseIntoMap(output, &gitReport); err != nil {
		return nil, fmt.Errorf("failed to parse report: %w", err)
	}

	return gitReport, err
}

func (c *BaseClient) GitRunRepoGC(appName string) error {
	cmd := fmt.Sprintf(gitRepoGcCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GitPurgeRepoCache(appName string) error {
	cmd := fmt.Sprintf(gitRepoPurgeCacheCmd, appName)
	_, err := c.Exec(cmd)
	return err
}
