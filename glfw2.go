// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package glfw2 provides API definitions for accessing the
//glfw dll. Definitions based on GLFW 2.7.9
package glfw2

//MOTE(t): all callbacks and calls are stdcall

import (
	"github.com/tHinqa/outside"
	. "runtime"
)

func init() {
	var dll string
	if GOOS == "windows" {
		dll = "glfw.dll"
	}
	outside.AddDllApis(dll, false, allApis)
}

type (
	void struct{}
	long int32
)

type GlfwVidMode struct {
	Width, Height                int
	RedBits, BlueBits, GreenBits int
}

type GlfwImage struct {
	Width, Height int
	Format        int
	BytesPerPixel int
	Data          *byte
}

type WindowMode int

const (
	Window WindowMode = iota + 0x00010001
	FullScreen
)

type GlfwCond *void
type GlfwMutex *void
type GlfwThread int

type GlfwCharFunc func(int, int)
type GlfwKeyFunc func(int, int)
type GlfwMouseButtonFunc func(int, int)
type GlfwMousePosFunc func(int, int)
type GlfwMouseWheelFunc func(int)
type GlfwThreadFunc func(*void)
type GlfwWindowCloseFunc func() int
type GlfwWindowRefreshFunc func()
type GlfwWindowSizeFunc func(int, int)

var GlfwBroadcastCond func(cond GlfwCond)
var GlfwCloseWindow func()
var GlfwCreateCond func() GlfwCond
var GlfwCreateMutex func() GlfwMutex
var GlfwCreateThread func(fun GlfwThreadFunc, arg *void) GlfwThread
var GlfwDestroyCond func(cond GlfwCond)
var GlfwDestroyMutex func(mutex GlfwMutex)
var GlfwDestroyThread func(ID GlfwThread)
var GlfwDisable func(token int)
var GlfwEnable func(token int)
var GlfwExtensionSupported func(extension string) int
var GlfwFreeImage func(img *GlfwImage)
var GlfwGetDesktopMode func(mode *GlfwVidMode)
var GlfwGetGLVersion func(major, minor, rev *int)
var GlfwGetJoystickButtons func(joy int, buttons *byte, numbuttons int) int
var GlfwGetJoystickParam func(joy, param int) int
var GlfwGetJoystickPos func(joy int, pos *float32, numaxes int) int
var GlfwGetKey func(key int) int
var GlfwGetMouseButton func(button int) int
var GlfwGetMousePos func(xpos, ypos *int)
var GlfwGetMouseWheel func() int
var GlfwGetNumberOfProcessors func() int
var GlfwGetProcAddress func(procname string) *void
var GlfwGetThreadID func() GlfwThread
var GlfwGetTime func() float64
var GlfwGetVersion func(major, minor, rev *int)
var GlfwGetVideoModes func(list *GlfwVidMode, maxcount int) int
var GlfwGetWindowParam func(param int) int
var GlfwGetWindowSize func(width, height *int)
var GlfwIconifyWindow func()
var GlfwInit func() bool
var GlfwLoadMemoryTexture2D func(data *void, size long, flags int) int
var GlfwLoadTexture2D func(name string, flags int) int
var GlfwLoadTextureImage2D func(img *GlfwImage, flags int) int
var GlfwLockMutex func(mutex GlfwMutex)
var GlfwOpenWindow func(width, height, redbits, greenbits, bluebits, alphabits, depthbits, stencilbits int, mode WindowMode) bool
var GlfwOpenWindowHint func(target, hint int)
var GlfwPollEvents func()
var GlfwReadImage func(name string, img *GlfwImage, flags int) int
var GlfwReadMemoryImage func(data *void, size long, img *GlfwImage, flags int) int
var GlfwRestoreWindow func()
var GlfwSetCharCallback func(cbfun GlfwCharFunc)
var GlfwSetKeyCallback func(cbfun GlfwKeyFunc)
var GlfwSetMouseButtonCallback func(cbfun GlfwMouseButtonFunc)
var GlfwSetMousePos func(xpos, ypos int)
var GlfwSetMousePosCallback func(cbfun GlfwMousePosFunc)
var GlfwSetMouseWheel func(pos int)
var GlfwSetMouseWheelCallback func(cbfun GlfwMouseWheelFunc)
var GlfwSetTime func(time float64)
var GlfwSetWindowCloseCallback func(cbfun GlfwWindowCloseFunc)
var GlfwSetWindowPos func(x, y int)
var GlfwSetWindowRefreshCallback func(cbfun GlfwWindowRefreshFunc)
var GlfwSetWindowSize func(width, height int)
var GlfwSetWindowSizeCallback func(cbfun GlfwWindowSizeFunc)
var GlfwSetWindowTitle func(title string)
var GlfwSignalCond func(cond GlfwCond)
var GlfwSleep func(time float64)
var GlfwSwapBuffers func()
var GlfwSwapInterval func(interval int)
var GlfwTerminate func()
var GlfwUnlockMutex func(mutex GlfwMutex)
var GlfwWaitCond func(cond GlfwCond, mutex GlfwMutex, timeout float64)
var GlfwWaitEvents func()
var GlfwWaitThread func(ID GlfwThread, waitmode int) int

var allApis = outside.Apis{
	{"glfwBroadcastCond", &GlfwBroadcastCond},
	{"glfwCloseWindow", &GlfwCloseWindow},
	{"glfwCreateCond", &GlfwCreateCond},
	{"glfwCreateMutex", &GlfwCreateMutex},
	{"glfwCreateThread", &GlfwCreateThread},
	{"glfwDestroyCond", &GlfwDestroyCond},
	{"glfwDestroyMutex", &GlfwDestroyMutex},
	{"glfwDestroyThread", &GlfwDestroyThread},
	{"glfwDisable", &GlfwDisable},
	{"glfwEnable", &GlfwEnable},
	{"glfwExtensionSupported", &GlfwExtensionSupported},
	{"glfwFreeImage", &GlfwFreeImage},
	{"glfwGetDesktopMode", &GlfwGetDesktopMode},
	{"glfwGetGLVersion", &GlfwGetGLVersion},
	{"glfwGetJoystickButtons", &GlfwGetJoystickButtons},
	{"glfwGetJoystickParam", &GlfwGetJoystickParam},
	{"glfwGetJoystickPos", &GlfwGetJoystickPos},
	{"glfwGetKey", &GlfwGetKey},
	{"glfwGetMouseButton", &GlfwGetMouseButton},
	{"glfwGetMousePos", &GlfwGetMousePos},
	{"glfwGetMouseWheel", &GlfwGetMouseWheel},
	{"glfwGetNumberOfProcessors", &GlfwGetNumberOfProcessors},
	{"glfwGetProcAddress", &GlfwGetProcAddress},
	{"glfwGetThreadID", &GlfwGetThreadID},
	{"glfwGetTime", &GlfwGetTime},
	{"glfwGetVersion", &GlfwGetVersion},
	{"glfwGetVideoModes", &GlfwGetVideoModes},
	{"glfwGetWindowParam", &GlfwGetWindowParam},
	{"glfwGetWindowSize", &GlfwGetWindowSize},
	{"glfwIconifyWindow", &GlfwIconifyWindow},
	{"glfwInit", &GlfwInit},
	{"glfwLoadMemoryTexture2D", &GlfwLoadMemoryTexture2D},
	{"glfwLoadTexture2D", &GlfwLoadTexture2D},
	{"glfwLoadTextureImage2D", &GlfwLoadTextureImage2D},
	{"glfwLockMutex", &GlfwLockMutex},
	{"glfwOpenWindow", &GlfwOpenWindow},
	{"glfwOpenWindowHint", &GlfwOpenWindowHint},
	{"glfwPollEvents", &GlfwPollEvents},
	{"glfwReadImage", &GlfwReadImage},
	{"glfwReadMemoryImage", &GlfwReadMemoryImage},
	{"glfwRestoreWindow", &GlfwRestoreWindow},
	{"glfwSetCharCallback", &GlfwSetCharCallback},
	{"glfwSetKeyCallback", &GlfwSetKeyCallback},
	{"glfwSetMouseButtonCallback", &GlfwSetMouseButtonCallback},
	{"glfwSetMousePos", &GlfwSetMousePos},
	{"glfwSetMousePosCallback", &GlfwSetMousePosCallback},
	{"glfwSetMouseWheel", &GlfwSetMouseWheel},
	{"glfwSetMouseWheelCallback", &GlfwSetMouseWheelCallback},
	{"glfwSetTime", &GlfwSetTime},
	{"glfwSetWindowCloseCallback", &GlfwSetWindowCloseCallback},
	{"glfwSetWindowPos", &GlfwSetWindowPos},
	{"glfwSetWindowRefreshCallback", &GlfwSetWindowRefreshCallback},
	{"glfwSetWindowSize", &GlfwSetWindowSize},
	{"glfwSetWindowSizeCallback", &GlfwSetWindowSizeCallback},
	{"glfwSetWindowTitle", &GlfwSetWindowTitle},
	{"glfwSignalCond", &GlfwSignalCond},
	{"glfwSleep", &GlfwSleep},
	{"glfwSwapBuffers", &GlfwSwapBuffers},
	{"glfwSwapInterval", &GlfwSwapInterval},
	{"glfwTerminate", &GlfwTerminate},
	{"glfwUnlockMutex", &GlfwUnlockMutex},
	{"glfwWaitCond", &GlfwWaitCond},
	{"glfwWaitEvents", &GlfwWaitEvents},
	{"glfwWaitThread", &GlfwWaitThread},
}
