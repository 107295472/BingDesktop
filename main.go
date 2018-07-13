package main

import (
	"flag"

	"encoding/json"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Constants
const htmlAbout = `图片来源必应.`

// Vars
var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "enables the debug mode")
	w       *astilectron.Window
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()
	AppName = "BingDesktop"
	// Run bootstrap
	astilog.Debugf("Running app built at %s", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
			//BaseDirectoryPath:"D:/GOPATH/src/yin/BingDesktop",
		},
		Debug:         false,
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.PtrStr("文件"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astilectron.PtrStr("关于"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if err := bootstrap.SendMessage(w, "about", htmlAbout, func(m *bootstrap.MessageIn) {
							// Unmarshal payload
							var s string
							if err := json.Unmarshal(m.Payload, &s); err != nil {
								astilog.Error(errors.Wrap(err, "unmarshaling payload failed"))
								return
							}
							astilog.Info(s)
							astilog.Infof("关于 is %s!", s)
						}); err != nil {
							astilog.Error(errors.Wrap(err, "sending about event failed"))
						}
						return
					},
				},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			// go func() {
			// 	time.Sleep(5 * time.Second)
			// 	if err := bootstrap.SendMessage(w, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
			// 		astilog.Error(errors.Wrap(err, "sending check.out.menu event failed"))
			// 	}
			// }()
			w.OpenDevTools()
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: bingHandleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astilectron.PtrStr("#333"),
				Center:          astilectron.PtrBool(true),
				Height:          astilectron.PtrInt(600),
				Width:           astilectron.PtrInt(1000),
				//AlwaysOnTop:     astilectron.PtrBool(true),
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
