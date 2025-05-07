//go:build !ci && !ios && !wasm && !test_web_driver && !mobile

package app










import "C"
import (
	"net/url"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
)

func (a *fyneApp) OpenURL(url *url.URL) error {
	cmd := exec.Command("open", url.String())
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return cmd.Run()
}



func (a *fyneApp) SetSystemTrayIcon(icon fyne.Resource) {
	a.Driver().(systrayDriver).SetSystemTrayIcon(icon)
}





func (a *fyneApp) SetSystemTrayMenu(menu *fyne.Menu) {
	if desk, ok := a.Driver().(systrayDriver); ok {
		desk.SetSystemTrayMenu(menu)
	}
}

func themeChanged() {
	fyne.CurrentApp().Settings().(*settings).setupTheme()
}

func watchTheme(_ *settings) {
	C.watchTheme()
}






func (a *fyneApp) registerRepositories() {
	// no-op
}
