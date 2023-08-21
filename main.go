package main

import (
	"context"
	"embed"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/mncred/mncplay/internal/build"
	"github.com/mncred/mncplay/internal/config"
	"github.com/mncred/mncplay/internal/logger"
	"github.com/mncred/mncplay/internal/system"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Read the App config
	configProvider := config.Provider{}
	if err := configProvider.LoadDefault(); err != nil {
		fmt.Printf("unable to load config: %v\n", err)
		os.Exit(1)
	}
	config := configProvider.Get()

	// Wails App Context
	var wailsCtx context.Context

	// Provide a logger
	lg := logger.New().
		WithName(config.Id).
		WithLevel(logger.ParseLevel(config.Debug.Log.Level)).
		WithOutput(config.Debug.Log.Output).
		WithMiddleware(FrontendMiddleware(func(script string) {
			if wailsCtx != nil {
				runtime.WindowExecJS(wailsCtx, script)
			}
		}))

	// Create application with options
	if err := wails.Run(&options.App{
		Logger:           lg.WithName("wails"),
		Title:            config.Window.Header.Title,
		Width:            config.Window.Width.Default,
		Height:           config.Window.Height.Default,
		MinWidth:         config.Window.Width.Min,
		MinHeight:        config.Window.Height.Min,
		MaxWidth:         config.Window.Width.Max,
		MaxHeight:        config.Window.Height.Max,
		DisableResize:    !config.Window.Resizable,
		Frameless:        !config.Window.Header.Native,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Debug: options.Debug{
			OpenInspectorOnStartup: config.Debug.Inspector,
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
			lg.WithName("front").Pointer(),
			&configProvider,
			&build.BuildInfo{},
			&system.Info{},
			&system.System{},
		},
		OnStartup: func(ctx context.Context) {
			wailsCtx = ctx

			// Basic log info
			lg.Print("log started: " + time.Now().Format(time.RFC1123Z))
			lg.Print("build time: " + build.BuildTime)
			lg.Print("build origin: " + build.BuildOrigin)
			lg.Print("build branch: " + build.BuildBranch)
			lg.Print("build commit: " + build.BuildCommit)
			lg.Print("build commit author: " + build.BuildCommitAuthor)
			lg.Print("build commit email: " + build.BuildCommitEmail)
			lg.Print("build compiler: " + build.BuildCompiler)
			lg.Print("build wails: " + build.BuildWails)

			runtime.EventsOn(ctx, "window:unminimise", func(data ...any) {
				runtime.WindowUnminimise(ctx)
			})
		},
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: config.Id,
			OnSecondInstanceLaunch: func(secondInstanceData options.SecondInstanceData) {
				runtime.WindowUnminimise(wailsCtx)
				runtime.Show(wailsCtx)
			},
		},
	}); err != nil {
		lg.Info("please, use Wails to build: https://wails.io")
		lg.Fatal(fmt.Sprintf("unable to launch app: %v\n", err))
	}
}

// FrontendMiddleware is the frontend logger proxy with formatters.
func FrontendMiddleware(handler func(script string)) func(level logger.Level, line string) string {
	return func(level logger.Level, line string) string {
		badgeStyleTrace := `color: white; background: black; margin: 0 4px 0 0; border-radius: 2px; font-weight: bold;`
		badgeStyleDebug := `color: white; background: gray; margin: 0 4px 0 0; border-radius: 2px; font-weight: bold;`
		badgeStyleInfo := `color: white; background: blue; margin: 0 4px 0 0; border-radius: 2px; font-weight: bold;`
		badgeStyleWarning := `color: white; background: orange; margin: 0 4px 0 0; border-radius: 2px; font-weight: bold;`
		badgeStyleError := `color: white; background: red; margin: 0 4px 0 0; border-radius: 2px; font-weight: bold;`
		badgeStyleFatal := `color: white; background: red; margin: 0 4px 0 0; border-radius: 2px; font-weight: bold;`
		badgeStylePrint := `color: white; background: gray; margin: 0 4px 0 0; border-radius: 2px; font-weight: bold;`

		badgeStyle := ""
		badgeText := ""
		script := "console."
		switch level {
		case logger.LevelTrace:
			script += "debug"
			badgeStyle = badgeStyleTrace
			badgeText = "TRC"
		case logger.LevelDebug:
			script += "debug"
			badgeStyle = badgeStyleDebug
			badgeText = "DBG"
		case logger.LevelInfo:
			script += "info"
			badgeStyle = badgeStyleInfo
			badgeText = "INF"
		case logger.LevelWarning:
			script += "warn"
			badgeStyle = badgeStyleWarning
			badgeText = "WRN"
		case logger.LevelError:
			script += "error"
			badgeStyle = badgeStyleError
			badgeText = "ERR"
		case logger.LevelFatal:
			script += "error"
			badgeStyle = badgeStyleFatal
			badgeText = "FTL"
		case logger.LevelPrint:
			script += "log"
			badgeStyle = badgeStylePrint
			badgeText = "PRN"
		}
		script += fmt.Sprintf("('%%c%%s%%c%%s', '%s', '%s', '', %s)", badgeStyle, badgeText, strconv.Quote(line))
		handler(script)
		return line
	}
}
