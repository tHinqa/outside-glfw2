// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glfw2

import (
	// . "runtime"
	"testing"
)

func TestInit(t *testing.T) {
	if GlfwInit() {
		//TODO(t) Find out why all 0
		// var major, minor, rev int
		// GlfwGetGLVersion(&major, &minor, &rev)
		// t.Log("GLFW Version:", major, minor, rev)
		var mode GlfwVidMode
		GlfwGetDesktopMode(&mode)
		t.Logf("%+v\n", mode)
		if GlfwOpenWindow(1024, 768,
			mode.RedBits, mode.GreenBits, mode.BlueBits,
			0, 0, 0, Window) {
			GlfwSetWindowTitle("Go GLFW 2")
			GlfwSetWindowPos(800, 0)
			GlfwSleep(0.5) // seconds
			t.Log(GlfwGetTime())
		} else {
			t.Fatal("GlfwOpenWindow: Failed to Open Window")
		}
		GlfwTerminate()
	} else {
		t.Fatal("GlfwInit: Failed to initialize GLFW")
	}
}
