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

type VidMode struct {
	Width, Height                int
	RedBits, BlueBits, GreenBits int
}

type Image struct {
	Width, Height int
	Format        int
	BytesPerPixel int
	Data          *byte
}

type WindowMode uint

const (
	Window WindowMode = iota + 0x00010001
	FullScreen
)

type EnableDisable uint

const (
	MouseCursor EnableDisable = iota + 0x00030001
	StickyKeys
	StickyMouseButtons
	SystemKeys
	KeyRepeat
	AutoPollEvents
)

type Cond *void
type Mutex *void
type Thread int

type CharFunc func(int, int)
type KeyFunc func(int, int)
type MouseButtonFunc func(int, int)
type MousePosFunc func(int, int)
type MouseWheelFunc func(int)
type ThreadFunc func(*void)
type WindowCloseFunc func() int
type WindowRefreshFunc func()
type WindowSizeFunc func(int, int)

var BroadcastCond func(cond Cond)
var CloseWindow func()
var CreateCond func() Cond
var CreateMutex func() Mutex
var CreateThread func(fun ThreadFunc, arg *void) Thread
var DestroyCond func(cond Cond)
var DestroyMutex func(mutex Mutex)
var DestroyThread func(ID Thread)
var Disable func(token EnableDisable)
var Enable func(token EnableDisable)
var ExtensionSupported func(extension string) int
var FreeImage func(img *Image)
var GetDesktopMode func(mode *VidMode)
var GetGLVersion func(major, minor, rev *int)
var GetJoystickButtons func(joy int, buttons *byte, numbuttons int) int
var GetJoystickParam func(joy, param int) int
var GetJoystickPos func(joy int, pos *float32, numaxes int) int
var GetKey func(key int) bool
var GetMouseButton func(button int) int
var GetMousePos func(xpos, ypos *int)
var GetMouseWheel func() int
var GetNumberOfProcessors func() int
var GetProcAddress func(procname string) *void
var GetThreadID func() Thread
var GetTime func() float64
var GetVersion func(major, minor, rev *int)
var GetVideoModes func(list *VidMode, maxcount int) int
var GetWindowParam func(param int) int
var GetWindowSize func(width, height *int)
var IconifyWindow func()
var Init func() bool
var LoadMemoryTexture2D func(data *void, size long, flags int) int
var LoadTexture2D func(name string, flags int) int
var LoadTextureImage2D func(img *Image, flags int) int
var LockMutex func(mutex Mutex)
var OpenWindow func(width, height, redbits, greenbits, bluebits, alphabits, depthbits, stencilbits int, mode WindowMode) bool
var OpenWindowHint func(target, hint int)
var PollEvents func()
var ReadImage func(name string, img *Image, flags int) int
var ReadMemoryImage func(data *void, size long, img *Image, flags int) int
var RestoreWindow func()
var SetCharCallback func(cbfun CharFunc)
var SetKeyCallback func(cbfun KeyFunc)
var SetMouseButtonCallback func(cbfun MouseButtonFunc)
var SetMousePos func(xpos, ypos int)
var SetMousePosCallback func(cbfun MousePosFunc)
var SetMouseWheel func(pos int)
var SetMouseWheelCallback func(cbfun MouseWheelFunc)
var SetTime func(time float64)
var SetWindowCloseCallback func(cbfun WindowCloseFunc)
var SetWindowPos func(x, y int)
var SetWindowRefreshCallback func(cbfun WindowRefreshFunc)
var SetWindowSize func(width, height int)
var SetWindowSizeCallback func(cbfun WindowSizeFunc)
var SetWindowTitle func(title string)
var SignalCond func(cond Cond)
var Sleep func(time float64)
var SwapBuffers func()
var SwapInterval func(interval int)
var Terminate func()
var UnlockMutex func(mutex Mutex)
var WaitCond func(cond Cond, mutex Mutex, timeout float64)
var WaitEvents func()
var WaitThread func(ID Thread, waitmode int) int

var allApis = outside.Apis{
	{"glfwBroadcastCond", &BroadcastCond},
	{"glfwCloseWindow", &CloseWindow},
	{"glfwCreateCond", &CreateCond},
	{"glfwCreateMutex", &CreateMutex},
	{"glfwCreateThread", &CreateThread},
	{"glfwDestroyCond", &DestroyCond},
	{"glfwDestroyMutex", &DestroyMutex},
	{"glfwDestroyThread", &DestroyThread},
	{"glfwDisable", &Disable},
	{"glfwEnable", &Enable},
	{"glfwExtensionSupported", &ExtensionSupported},
	{"glfwFreeImage", &FreeImage},
	{"glfwGetDesktopMode", &GetDesktopMode},
	{"glfwGetGLVersion", &GetGLVersion},
	{"glfwGetJoystickButtons", &GetJoystickButtons},
	{"glfwGetJoystickParam", &GetJoystickParam},
	{"glfwGetJoystickPos", &GetJoystickPos},
	{"glfwGetKey", &GetKey},
	{"glfwGetMouseButton", &GetMouseButton},
	{"glfwGetMousePos", &GetMousePos},
	{"glfwGetMouseWheel", &GetMouseWheel},
	{"glfwGetNumberOfProcessors", &GetNumberOfProcessors},
	{"glfwGetProcAddress", &GetProcAddress},
	{"glfwGetThreadID", &GetThreadID},
	{"glfwGetTime", &GetTime},
	{"glfwGetVersion", &GetVersion},
	{"glfwGetVideoModes", &GetVideoModes},
	{"glfwGetWindowParam", &GetWindowParam},
	{"glfwGetWindowSize", &GetWindowSize},
	{"glfwIconifyWindow", &IconifyWindow},
	{"glfwInit", &Init},
	{"glfwLoadMemoryTexture2D", &LoadMemoryTexture2D},
	{"glfwLoadTexture2D", &LoadTexture2D},
	{"glfwLoadTextureImage2D", &LoadTextureImage2D},
	{"glfwLockMutex", &LockMutex},
	{"glfwOpenWindow", &OpenWindow},
	{"glfwOpenWindowHint", &OpenWindowHint},
	{"glfwPollEvents", &PollEvents},
	{"glfwReadImage", &ReadImage},
	{"glfwReadMemoryImage", &ReadMemoryImage},
	{"glfwRestoreWindow", &RestoreWindow},
	{"glfwSetCharCallback", &SetCharCallback},
	{"glfwSetKeyCallback", &SetKeyCallback},
	{"glfwSetMouseButtonCallback", &SetMouseButtonCallback},
	{"glfwSetMousePos", &SetMousePos},
	{"glfwSetMousePosCallback", &SetMousePosCallback},
	{"glfwSetMouseWheel", &SetMouseWheel},
	{"glfwSetMouseWheelCallback", &SetMouseWheelCallback},
	{"glfwSetTime", &SetTime},
	{"glfwSetWindowCloseCallback", &SetWindowCloseCallback},
	{"glfwSetWindowPos", &SetWindowPos},
	{"glfwSetWindowRefreshCallback", &SetWindowRefreshCallback},
	{"glfwSetWindowSize", &SetWindowSize},
	{"glfwSetWindowSizeCallback", &SetWindowSizeCallback},
	{"glfwSetWindowTitle", &SetWindowTitle},
	{"glfwSignalCond", &SignalCond},
	{"glfwSleep", &Sleep},
	{"glfwSwapBuffers", &SwapBuffers},
	{"glfwSwapInterval", &SwapInterval},
	{"glfwTerminate", &Terminate},
	{"glfwUnlockMutex", &UnlockMutex},
	{"glfwWaitCond", &WaitCond},
	{"glfwWaitEvents", &WaitEvents},
	{"glfwWaitThread", &WaitThread},
}
