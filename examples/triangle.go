package main

import (
	"fmt"
	. "github.com/tHinqa/outside-glfw2"
	gl "github.com/tHinqa/outside-opengl"
)

func main() {
	var (
		width, height, frames, x, y int
		fps                         float64
		titlestr                    string
	)

	var mode VidMode

	Init()
	defer Terminate()
	GetDesktopMode(&mode)
	if !OpenWindow(640, 480, 0, 0, 0, 0, 0, 0, Window) {
		panic("OpenWindow failed")
	}
	SetWindowPos(mode.Width/2 /*-mode.Width/4*/, mode.Height/2-mode.Height/4)
	Enable(StickyKeys)
	SwapInterval(0)

	running := true
	t0 := GetTime()
	for running {
		tt := GetTime()
		GetMousePos(&x, &y)
		if (tt-t0) > 1.0 || frames == 0 {
			fps = float64(frames) / (tt - t0)
			titlestr = fmt.Sprintf("Spinning Triangle (%.1f FPS %ds)", fps, int(tt))
			SetWindowTitle(titlestr)
			t0 = tt
			frames = 0
		}
		frames++
		GetWindowSize(&width, &height)
		if height <= 0 {
			height = 1
		}
		gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
		gl.ClearColor(0, 0, 0, 0)
		gl.Clear(gl.ColorBufferBit)
		gl.MatrixMode(gl.Projection)
		gl.LoadIdentity()
		gl.Perspective(65, float64(width)/float64(height), 1, 100)
		gl.MatrixMode(gl.ModelView)
		gl.LoadIdentity()
		gl.LookAt(
			0, 1, 0, // Eye-position
			0, 20, 0, // View-point
			0, 0, 1) // Up-vector
		gl.Translatef(0, 14, 0)
		gl.Rotatef(0.3*float32(x)+float32(tt)*100, 0, 0, 1)
		gl.Begin(gl.Triangles)
		gl.Color3f(1, 0, 0)
		gl.Vertex3f(-5, 0, -4)
		gl.Color3f(0, 1, 0)
		gl.Vertex3f(5, 0, -4)
		gl.Color3f(0, 0, 1)
		gl.Vertex3f(0, 0, 6)
		gl.End()
		SwapBuffers()
		Sleep(0.008) // not in original (limits ~1300fps sys to ~100fps)
		running = !GetKey(KeyEsc) &&
			GetWindowParam(0x00020001) != 0 //TODO(t): GLFW_OPENED
	}
}
