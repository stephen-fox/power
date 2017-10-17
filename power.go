package power

import (
	"errors"
	"os/exec"
	"os/user"
	"runtime"
)

// Power provides several methods for controlling a machine's power state.
type Power interface {
	// Restart restarts the machine.
	Restart() error

	// Shutdown turns the machine off.
	Shutdown() error

	// Sleep puts the machine to sleep.
	Sleep() error
}

// Get gets the Power implementation for the current operating system.
func Get() (Power, error) {
	switch runtime.GOOS {
	case "darwin":
		return new(darwin), nil
	case "linux":
		return new(linux), nil
	case "windows":
		return new(windows), nil
	}

	var power Power
	return power, errors.New("No Power implementation for this OS")
}

type darwin struct {
}

func (d *darwin) Restart() error {
	err := isRoot()
	if err != nil {
		return err
	}

	reboot := exec.Command("reboot")
	_, err = reboot.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (d *darwin) Shutdown() error {
	err := isRoot()
	if err != nil {
		return err
	}

	shutdown := exec.Command("shutdown", "-h", "now")
	_, err = shutdown.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (d *darwin) Sleep() error {
	err := isRoot()
	if err != nil {
		return err
	}

	pmset := exec.Command("pmset", "sleepnow")
	_, err = pmset.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

type linux struct {
}

func (l *linux) Restart() error {
	err := isRoot()
	if err != nil {
		return err
	}

	reboot := exec.Command("reboot")
	_, err = reboot.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (l *linux) Shutdown() error {
	err := isRoot()
	if err != nil {
		return err
	}

	poweroff := exec.Command("poweroff")
	_, err = poweroff.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (l *linux) Sleep() error {
	err := isRoot()
	if err != nil {
		return err
	}

	systemctl := exec.Command("systemctl")
	_, err = systemctl.CombinedOutput()
	if err == nil {
		// This system has 'systemctl'.
		suspend := exec.Command("systemctl", "suspend")
		_, err := suspend.CombinedOutput()
		if err != nil {
			return err
		}
	}

	pm := exec.Command("pm-suspend")
	_, err = pm.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func isRoot() error {
	currentUser, err := user.Current()
	if err != nil {
		return errors.New("Failed to check if current user is root: " + err.Error())
	}
	if currentUser.Username != "root" {
		return errors.New("The 'root' user is required for this action")
	}

	return nil
}

type windows struct {
}

func (w *windows) Restart() error {
	restart := exec.Command("shutdown", "/r")
	_, err := restart.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (w *windows) Shutdown() error {
	shutdown := exec.Command("shutdown", "/s")
	_, err := shutdown.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (w *windows) Sleep() error {
	rundll := exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", "0,1,0")
	_, err := rundll.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
