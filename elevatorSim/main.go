package main

import (
	"flag"

	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	astilog "github.com/asticode/go-astilog"
	"github.com/pkg/errors"
	//"encoding/json"
/*
	docker "elevatorSim/dockerRun"
	mng "elevatorSim/elevator"
	"elevatorSim/traffic"
*/
	"elevatorSim/clock"
	"time"

)

// Constants
const htmlAbout = `Welcome on <b>Astilectron</b> demo!<br>
This is using the bootstrap and the bundler.`

// Vars
var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "enables the debug mode")
	w       *astilectron.Window
)

const (
	NR_ELEVATOR  = 2
	TESTDURATION = 10
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	astilog.Debugf("Running app built at %s", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon/icon.icns",
			AppIconDefaultPath: "resources/icon/icon.png",
		},
		Debug: *debug,
		RestoreAssets: RestoreAssets,
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			// space for clock
			c := clock.NewClock()
			w = ws[0]
			w.OpenDevTools()

			ticker := time.NewTicker(200 * time.Millisecond)
			quit := make(chan struct{})
			go func() {
				for {
				select {
					case <- ticker.C:
						clockString := c.GetClock().Format("01-02 15:04:05")
						if err := bootstrap.SendMessage(w, "clock", clockString); err != nil {
							astilog.Error(errors.Wrap(err, "send clock failed"))
						}
					case <- quit:
						ticker.Stop()
						return
					}
				}
			}()
			return nil
		},
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				Center:          astilectron.PtrBool(true),
				Height:          astilectron.PtrInt(700),
				Width:           astilectron.PtrInt(900),
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}