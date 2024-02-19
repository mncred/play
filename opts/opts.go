package opts

// --ldflags vars
var (
	// BuildTime - compilation time
	BuildTime = ""
	// Origin - compilation git origin
	Origin = ""
	// Branch - compilation git branch
	Branch = ""
	// Commit - compilation git commit hash
	Commit = ""
	// CommitAuthor - compilation git commit author name
	CommitAuthor = ""
	// CommitEmail - compilation git commit author email
	CommitEmail = ""
	// CommitTag - compilation git commit tag
	CommitTag = ""
	// Compiler - compilation go version
	Compiler = ""
	// Wails - compilation wails version
	Wails = ""
	// OS - compilation OS target
	OS = ""
	// Arch - compilation arch target
	Arch = ""
	// LogLevel - default log level (debug, info, warn, error, etc)
	LogLevel = ""
	// LogsDir - default location for log files
	LogsDir = ""
	// AppId - unique app id
	AppId = ""
	// WindowTitle - app window title
	WindowTitle = ""
	// WindowWidth - app window width
	WindowWidth = ""
	// WindowHeight - app window height
	WindowHeight = ""
	// WindowWidthMin - app window min width
	WindowWidthMin = ""
	// WindowHeightMax - app window max height
	WindowWidthMax = ""
	// WindowHeightMin - app window min height
	WindowHeightMin = ""
	// WindowHeightMax - app window max height
	WindowHeightMax = ""
	// WindowDisableResize - app window disable resizing
	WindowDisableResize = ""
	// WindowFrameless - app window frameless
	WindowFrameless = ""
	// WindowDisableMaximize - app window disable maximize button
	WindowDisableMaximize = ""
	// WindowDisableMinimize - app window disable minimize button
	WindowDisableMinimize = ""
	// WindowDisableClose - app window disable close button
	WindowDisableClose = ""
)

type Options struct {
	BuildTime             string `json:"buildTime"`
	Origin                string `json:"origin"`
	Branch                string `json:"branch"`
	Commit                string `json:"commit"`
	CommitAuthor          string `json:"commitAuthor"`
	CommitEmail           string `json:"commitEmail"`
	CommitTag             string `json:"commitTag"`
	Compiler              string `json:"compiler"`
	Wails                 string `json:"wails"`
	OS                    string `json:"os"`
	Arch                  string `json:"arch"`
	LogLevel              string `json:"logLevel"`
	LogsDir               string `json:"logsDir"`
	AppId                 string `json:"appId"`
	WindowTitle           string `json:"windowTitle"`
	WindowWidth           string `json:"windowWidth"`
	WindowHeight          string `json:"windowHeight"`
	WindowWidthMin        string `json:"windowWidthMin"`
	WindowWidthMax        string `json:"windowWidthMax"`
	WindowHeightMin       string `json:"windowHeightMin"`
	WindowHeightMax       string `json:"windowHeightMax"`
	WindowDisableResize   string `json:"windowDisableResize"`
	WindowFrameless       string `json:"windowFrameless"`
	WindowDisableMaximize string `json:"windowDisableMaximize"`
	WindowDisableMinimize string `json:"windowDisableMinimize"`
	WindowDisableClose    string `json:"windowDisableClose"`
}

// Opts provides global list of the app options
var Opts = Options{
	BuildTime:             BuildTime,
	Origin:                Origin,
	Branch:                Branch,
	Commit:                Commit,
	CommitAuthor:          CommitAuthor,
	CommitEmail:           CommitEmail,
	CommitTag:             CommitTag,
	Compiler:              Compiler,
	Wails:                 Wails,
	OS:                    OS,
	Arch:                  Arch,
	LogLevel:              LogLevel,
	LogsDir:               LogsDir,
	AppId:                 AppId,
	WindowTitle:           WindowTitle,
	WindowWidth:           WindowWidth,
	WindowHeight:          WindowHeight,
	WindowWidthMin:        WindowWidthMin,
	WindowWidthMax:        WindowWidthMax,
	WindowHeightMin:       WindowHeightMin,
	WindowHeightMax:       WindowHeightMax,
	WindowDisableResize:   WindowDisableResize,
	WindowFrameless:       WindowFrameless,
	WindowDisableMaximize: WindowDisableMaximize,
	WindowDisableMinimize: WindowDisableMinimize,
	WindowDisableClose:    WindowDisableClose,
}

// Get returns self options.
// This is a wrapper to pass into Wails Front-End.
func (o Options) Get() Options {
	return o
}
