package main

import (
	"context"
	"embed"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	logrusAdapter "mncplay/internal/logger/adapter/logrus"
	"mncplay/internal/system"
	"mncplay/opts"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Configure global logger
	initLogger(opts.LogsDir, opts.LogLevel)

	// Configure local loggers
	mainLogger := logrusAdapter.Provide(logrus.WithField("module", "main"))
	wailsLogger := logrusAdapter.Provide(logrus.WithField("module", "wails"))
	frontendLogger := logrusAdapter.Provide(logrus.WithField("module", "frontend"))

	// Wails App Context
	var wailsCtx context.Context

	// Create application with options
	if err := wails.Run(&options.App{
		Logger:           wailsLogger,
		Title:            opts.WindowTitle,
		Width:            func() int { v, _ := strconv.ParseInt(opts.WindowWidth, 10, 32); return int(v) }(),
		Height:           func() int { v, _ := strconv.ParseInt(opts.WindowHeight, 10, 32); return int(v) }(),
		MinWidth:         func() int { v, _ := strconv.ParseInt(opts.WindowWidthMin, 10, 32); return int(v) }(),
		MinHeight:        func() int { v, _ := strconv.ParseInt(opts.WindowHeightMin, 10, 32); return int(v) }(),
		MaxWidth:         func() int { v, _ := strconv.ParseInt(opts.WindowWidthMax, 10, 32); return int(v) }(),
		MaxHeight:        func() int { v, _ := strconv.ParseInt(opts.WindowHeightMax, 10, 32); return int(v) }(),
		DisableResize:    func() bool { v, _ := strconv.ParseBool(opts.WindowHeightMax); return v }(),
		Frameless:        func() bool { v, _ := strconv.ParseBool(opts.WindowFrameless); return v }(),
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
		Mac: &mac.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
			IsZoomControlEnabled: false,
			ZoomFactor:           1,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: false,
			Icon:                icon,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			frontendLogger,
			&opts.Opts,
			&system.Info{},
			&system.System{},
		},
		OnStartup: func(ctx context.Context) {
			wailsCtx = ctx

			// Basic log info
			mainLogger.Info("log started: " + time.Now().Format(time.RFC1123Z))
			mainLogger.Info("build OS: " + opts.OS)
			mainLogger.Info("build arch: " + opts.Arch)
			mainLogger.Info("build time: " + opts.BuildTime)
			mainLogger.Info("build origin: " + opts.Origin)
			mainLogger.Info("build branch: " + opts.Branch)
			mainLogger.Info("build commit: " + opts.Commit)
			mainLogger.Info("build commit author: " + opts.CommitAuthor)
			mainLogger.Info("build commit email: " + opts.CommitEmail)
			mainLogger.Info("build compiler: " + opts.Compiler)
			mainLogger.Info("build wails: " + opts.Wails)
			mainLogger.Info("build log level: " + opts.LogLevel)

			runtime.EventsOn(ctx, "window:unminimise", func(data ...any) {
				runtime.WindowUnminimise(ctx)
			})
		},
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: opts.AppId,
			OnSecondInstanceLaunch: func(secondInstanceData options.SecondInstanceData) {
				runtime.WindowUnminimise(wailsCtx)
				runtime.Show(wailsCtx)
			},
		},
	}); err != nil {
		mainLogger.Info("please, use Wails to build: https://wails.io")
		mainLogger.Fatal(fmt.Sprintf("unable to launch app: %v\n", err))
	}
}

func initLogger(logsDir string, level string) {
	// log formatter
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// log level
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Printf("build log level not defined, use debug level: %v\n", err)
		lvl = logrus.DebugLevel
	}
	logrus.SetLevel(lvl)

	// log outputs (stderr + file / stderr only)
	var wr io.Writer = os.Stderr
	if len(logsDir) > 0 {
		os.MkdirAll(logsDir, os.ModePerm)

		// rewrite latest log file
		latestFile, err := os.OpenFile(
			path.Join(logsDir, "latest.log"),
			os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
			os.ModePerm,
		)
		if err != nil {
			fmt.Printf("unable to create latest log file, logs won't be saved: %v\n", err)
			return
		}

		// timestamped log file
		timestampedFile, err := os.OpenFile(
			path.Join(logsDir, strconv.FormatInt(time.Now().Unix(), 10)+".log"),
			os.O_CREATE|os.O_WRONLY,
			os.ModePerm,
		)
		if err != nil {
			fmt.Printf("unable to create timestamped log file, logs won't be saved: %v\n", err)
			return
		}

		wr = io.MultiWriter(os.Stderr, latestFile, timestampedFile)
	} else {
		fmt.Printf("build logs dir not defined, logs won't be created\n")
	}
	logrus.SetOutput(wr)
}
