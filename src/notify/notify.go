package notify

import "os/exec"

func Show(title, description string) {
	exec.Command("notify-send", title, description).Run()
}
