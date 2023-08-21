package build

var (
	BuildTime         = ""
	BuildOrigin       = ""
	BuildBranch       = ""
	BuildCommit       = ""
	BuildCommitAuthor = ""
	BuildCommitEmail  = ""
	BuildTag          = ""
	BuildCompiler     = ""
	BuildWails        = ""
)

type BuildInfo struct{}

func (b *BuildInfo) BuildTime() string {
	return BuildTime
}

func (b *BuildInfo) BuildOrigin() string {
	return BuildOrigin
}

func (b *BuildInfo) BuildCommit() string {
	return BuildCommit
}

func (b *BuildInfo) BuildBranch() string {
	return BuildBranch
}

func (b *BuildInfo) BuildCommitAuthor() string {
	return BuildCommitAuthor
}

func (b *BuildInfo) BuildCommitEmail() string {
	return BuildCommitEmail
}

func (b *BuildInfo) BuildTag() string {
	return BuildTag
}

func (b *BuildInfo) BuildCompiler() string {
	return BuildCompiler
}

func (b *BuildInfo) BuildWails() string {
	return BuildWails
}
