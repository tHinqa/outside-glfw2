// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glfw2

import (
	// . "runtime"
	"fmt"
	gl "github.com/tHinqa/outside-opengl"
	"testing"
)

func aTestInit(t *testing.T) {
	if Init() {
		var major, minor, rev int
		//TODO(t) Find out why 0.0.0
		GetGLVersion(&major, &minor, &rev)
		t.Logf("GLFW Version: %v.%v.%v", major, minor, rev)
		var mode VidMode
		GetDesktopMode(&mode)
		t.Logf("%+v", mode)
		if OpenWindow(mode.Width/2, mode.Height/2,
			mode.RedBits, mode.GreenBits, mode.BlueBits,
			0, 0, 0, Window) {
			SetWindowTitle("Go GLFW 2")
			SetWindowPos(mode.Width/2-mode.Width/4, mode.Height/2-mode.Height/4)
			gl.Clear(gl.ColorBufferBit)
			SwapBuffers()
			Sleep(2.5) // seconds
			t.Log("Elapsed:", GetTime(), "seconds")
		} else {
			t.Fatal("OpenWindow: Failed to Open Window")
		}
		Terminate()
	} else {
		t.Fatal("Init: Failed to initialize GLFW")
	}
}

//TODO(t): Credits
func TestTriangles(t *testing.T) {
	var (
		width, height, frames, x, y int
		fps                         float64
		titlestr                    string
	)

	Init()
	defer Terminate()
	if !OpenWindow(640, 480, 0, 0, 0, 0, 0, 0, Window) {
		t.Fatal("OpenWindow failed")
	}
	Enable(StickyKeys)
	SwapInterval(0)

	running := true
	t0 := GetTime()
	for running {
		tt := GetTime()
		GetMousePos(&x, &y)
		if (tt-t0) > 1.0 || frames == 0 {
			fps = float64(frames) / (tt - t0)
			titlestr = fmt.Sprintf("Spinning Triangle (%.1f FPS)", fps)
			SetWindowTitle(titlestr)
			t0 = tt
			frames = 0
		}
		frames++
		GetWindowSize(&width, &height)
		if height <= 0 {
			height = 1
		}
		gl.Viewport(0, 0, gl.GLsizei(width), gl.GLsizei(height))
		gl.ClearColor(0, 0, 0, 0)
		gl.Clear(gl.ColorBufferBit)
		gl.MatrixMode(0x1701) //TODO(t): GL_PROJECTION
		gl.LoadIdentity()
		// gluPerspective(65, float32(width)/float32(height), 1, 100)
		gl.MatrixMode(0x1700) //TODO(t): GL_MODELVIEW
		gl.LoadIdentity()
		// gluLookAt(0, 1, 0, // Eye-position
		// 0, 20, 0, // View-point
		// 0, 0, 1) // Up-vector
		gl.Translatef(0, 14, 0)
		gl.Rotatef(0.3*gl.GLfloat(x)+gl.GLfloat(tt)*100, 0, 0, 1)
		gl.Begin(0x0004) //TODO(t): GL_TRIANGLES
		gl.Color3f(1, 0, 0)
		gl.Vertex3f(-5, 0, -4)
		gl.Color3f(0, 1, 0)
		gl.Vertex3f(5, 0, -4)
		gl.Color3f(0, 0, 1)
		gl.Vertex3f(0, 0, 6)
		gl.End()
		SwapBuffers()
		running = !GetKey(257) && //TODO(t): GLFW_KEY_ESC
			GetWindowParam(0x00020001) != 0 //TODO(t): GLFW_OPENED
	}
	Terminate()
	t.Log(titlestr)
}
