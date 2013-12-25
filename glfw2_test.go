// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glfw2

import (
	gl "github.com/tHinqa/outside-opengl"
	"testing"
)

var mode VidMode

func TestInit(t *testing.T) {
	if Init() {
		GetDesktopMode(&mode)
		var major, minor, rev int
		//TODO(t) Find out why 0.0.0
		GetGLVersion(&major, &minor, &rev)
		t.Logf("GLFW Version: %v.%v.%v", major, minor, rev)
		t.Logf("%+v", mode)
		if OpenWindow(mode.Width/2, mode.Height/2,
			mode.RedBits, mode.GreenBits, mode.BlueBits,
			0, 0, 0, Window) {
			SetWindowTitle("Go GLFW 2")
			SetWindowPos(mode.Width/2-mode.Width/4, mode.Height/2-mode.Height/4)
			gl.Clear(gl.ColorBufferBit)
			SwapBuffers()
			Sleep(2.5) // seconds
			t.Logf("Elapsed (including 2.5s sleep): %fs", GetTime())
			GetTime() // discard
			s := GetTime()
			t.Logf("GetTime() call duration: ~%dns", int((GetTime()-s)*1e9))
		} else {
			t.Fatal("OpenWindow: Failed to Open Window")
		}
		Terminate()
	} else {
		t.Fatal("Init: Failed to initialize GLFW")
	}
}

//TODO(t): Credits
func TestMode(t *testing.T) {
	const MAX_NUM_MODES = 400
	var (
		dtmode VidMode
		modes  [MAX_NUM_MODES]VidMode
	)

	if !Init() {
		t.Fail()
	}

	GetDesktopMode(&dtmode)
	t.Logf("Desktop mode: %d x %d x %d\n",
		dtmode.Width, dtmode.Height, dtmode.RedBits+
			dtmode.GreenBits+dtmode.BlueBits)
	//TODO(t): make slice
	modecount := GetVideoModes(&modes[0], MAX_NUM_MODES)
	t.Log("Available modes:")
	for i := 0; i < modecount; i++ {
		t.Logf("%3d: %d x %d x %d\n", i,
			modes[i].Width, modes[i].Height, modes[i].RedBits+
				modes[i].GreenBits+modes[i].BlueBits)
	}
	Terminate()
}
