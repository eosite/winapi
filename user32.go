// Copyright 2013 The winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

var (
	modUser32              = syscall.NewLazyDLL("user32.dll")
	procBeginPaint         = modUser32.NewProc("BeginPaint")
	procCreateDialogParam  = modUser32.NewProc("CreateDialogParamW")
	procCreateWindowEx     = modUser32.NewProc("CreateWindowExW")
	procDefWindowProc      = modUser32.NewProc("DefWindowProcW")
	procDestroyWindow      = modUser32.NewProc("DestroyWindow")
	procDialogBoxParam     = modUser32.NewProc("DialogBoxParamW")
	procDispatchMessage    = modUser32.NewProc("DispatchMessageW")
	procEndDialog          = modUser32.NewProc("EndDialog")
	procEndPaint           = modUser32.NewProc("EndPaint")
	procGetDC              = modUser32.NewProc("GetDC")
	procGetDlgItem         = modUser32.NewProc("GetDlgItem")
	procGetMessage         = modUser32.NewProc("GetMessageW")
	procGetWindowLong      = modUser32.NewProc("GetWindowLongW")
	procGetWindowLongPtr   = modUser32.NewProc("GetWindowLongPtrW")
	procLoadCursor         = modUser32.NewProc("LoadCursorW")
	procLoadIcon           = modUser32.NewProc("LoadIconW")
	procLoadMenu           = modUser32.NewProc("LoadMenuW")
	procLoadString         = modUser32.NewProc("LoadStringW")
	procMessageBox         = modUser32.NewProc("MessageBoxW")
	procUnregisterClass    = modUser32.NewProc("UnregisterClassW")
	procPostMessage        = modUser32.NewProc("PostMessageW")
	procPostQuitMessage    = modUser32.NewProc("PostQuitMessage")
	procRegisterClassEx    = modUser32.NewProc("RegisterClassExW")
	procReleaseDC          = modUser32.NewProc("ReleaseDC")
	procSendMessage        = modUser32.NewProc("SendMessageW")
	procSendDlgItemMessage = modUser32.NewProc("SendDlgItemMessageW")
	procSetMenu            = modUser32.NewProc("SetMenu")
	procSetWindowLong      = modUser32.NewProc("SetWindowLongW")
	procSetWindowLongPtr   = modUser32.NewProc("SetWindowLongPtrW")
	procShowWindow         = modUser32.NewProc("ShowWindow")
	procTranslateMessage   = modUser32.NewProc("TranslateMessage")
	procUpdateWindow       = modUser32.NewProc("UpdateWindow")
)

var is64Bit bool = false

type PaintStruct struct {
	Hdc         HDC
	FErase      BOOL
	RcPaint     Rect
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]byte
}

func BeginPaint(h HWND, ps *PaintStruct) HDC {
	var ret uintptr
	ret, _, lastError = procBeginPaint.Call(uintptr(h),
		uintptr(unsafe.Pointer(ps)))

	return HDC(ret)
}

func CreateDialogParam(instRes HINSTANCE, name string, parent HWND,
	proc uintptr, param uintptr) HWND {
	var ret uintptr
	ret, _, lastError = procCreateDialogParam.Call(uintptr(instRes),
		resourceNameToPtr(name), uintptr(parent), proc, param)

	return HWND(ret)
}

type CREATESTRUCT struct {
	CreateParams    uintptr
	Instance        HINSTANCE
	Menu            HMENU
	Parent          HWND
	Cy              int32
	Cx              int32
	Y               int32
	X               int32
	Style           int32
	Name, ClassName uintptr
	ExStyle         uint32
}

type CreateWindowExParam struct {
	ClassName, WindowName string
	Style, ExStyle        uint32
	X, Y, Width, Height   int
	Parent                HWND
	Menu                  HMENU
	Instance              HINSTANCE
	Param                 uintptr
}

func CreateWindowEx(param *CreateWindowExParam) HWND {
	var ret uintptr
	ret, _, lastError = procCreateWindowEx.Call(uintptr(param.ExStyle),
		StringToUintptr(param.ClassName), StringToUintptr(param.WindowName),
		uintptr(param.Style), uintptr(param.X), uintptr(param.Y),
		uintptr(param.Width), uintptr(param.Height), uintptr(param.Parent),
		uintptr(param.Menu), uintptr(param.Instance), param.Param)

	return HWND(ret)
}

func DefWindowProc(m *MSG) LRESULT {
	var ret uintptr
	ret, _, lastError = procDefWindowProc.Call(uintptr(m.HWnd), uintptr(m.Msg),
		uintptr(m.WParam), uintptr(m.LParam))

	return LRESULT(ret)
}

func DestroyWindow(h HWND) bool {
	var ret uintptr
	ret, _, lastError = procDestroyWindow.Call(uintptr(h))

	return PtrToBool(ret)
}

func DialogBoxParam(instRes HINSTANCE, name string, parent HWND,
	proc uintptr, param uintptr) int {
	var ret uintptr
	ret, _, lastError = procDialogBoxParam.Call(uintptr(instRes),
		resourceNameToPtr(name), uintptr(parent), proc, param)

	return int(ret)
}

func DispatchMessage(m *WinMSG) LRESULT {
	var ret uintptr
	ret, _, lastError = procDispatchMessage.Call(uintptr(unsafe.Pointer(m)))

	return LRESULT(ret)
}

func EndDialog(h HWND, result int) bool {
	var ret uintptr
	ret, _, lastError = procEndDialog.Call(uintptr(h), uintptr(result))

	return PtrToBool(ret)
}

func EndPaint(h HWND, ps *PaintStruct) bool {
	var ret uintptr
	ret, _, lastError = procEndPaint.Call(uintptr(h), uintptr(unsafe.Pointer(ps)))

	return PtrToBool(ret)
}

func GetDC(h HWND) HDC {
	var ret uintptr
	ret, _, lastError = procGetDC.Call(uintptr(h))

	return HDC(ret)
}

func GetDlgItem(h HWND, id int) HWND {
	var ret uintptr
	ret, _, lastError = procGetDlgItem.Call(uintptr(h), uintptr(id))

	return HWND(ret)
}

func GetMessage(m *WinMSG, h HWND, min, max UINT) bool {
	var ret uintptr
	ret, _, lastError = procGetMessage.Call(uintptr(unsafe.Pointer(m)), uintptr(h),
		uintptr(min), uintptr(max))

	return PtrToBool(ret)
}

func GetWindowLongPtr(h HWND, index int) (ret uintptr) {
	if is64Bit {
		ret, _, _ = procGetWindowLongPtr.Call(uintptr(h), uintptr(index))
	} else {
		ret, _, _ = procGetWindowLong.Call(uintptr(h), uintptr(index))
	}

	return
}

func LoadCursor(instRes HINSTANCE, name string) HCURSOR {
	var ret uintptr
	ret, _, lastError = procLoadCursor.Call(uintptr(instRes), resourceNameToPtr(name))

	return HCURSOR(ret)
}

func LoadIcon(instRes HINSTANCE, name string) HICON {
	var ret uintptr
	ret, _, lastError = procLoadIcon.Call(uintptr(instRes), resourceNameToPtr(name))

	return HICON(ret)
}

func LoadMenu(instRes HINSTANCE, name string) HMENU {
	var ret uintptr
	ret, _, lastError = procLoadMenu.Call(uintptr(instRes), resourceNameToPtr(name))

	return HMENU(ret)
}

func LoadString(inst HINSTANCE, id uint) (ret string) {
	var text [4096]uint16
	var r uintptr
	r, _, lastError = procLoadString.Call(uintptr(inst), uintptr(id),
		uintptr(unsafe.Pointer(&text[0])), 4096)

	if int(r) <= 0 {
		ret = ""
	} else {
		ret = string(utf16.Decode(text[0:r]))
	}

	return
}

func MessageBox(parent HWND, text, title string, boxType uint) int {
	var ret uintptr
	ret, _, lastError = procMessageBox.Call(uintptr(parent),
		StringToUintptr(text), StringToUintptr(title), uintptr(boxType))

	return int(ret)
}

func UnregisterClass(name string) bool {
	var ret uintptr
	ret, _, lastError = procUnregisterClass.Call(StringToUintptr(name), 0)

	return PtrToBool(ret)
}

func PostMessage(m *MSG) bool {
	var ret uintptr
	ret, _, lastError = procPostMessage.Call(uintptr(m.HWnd), uintptr(m.Msg),
		uintptr(m.WParam), uintptr(m.LParam))

	return PtrToBool(ret)
}

func PostQuitMessage(code int) {
	procPostQuitMessage.Call(uintptr(code))
}

type RegisterClassExParam struct {
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Icon       HICON
	Cursor     HCURSOR
	Background HBRUSH
	MenuName   string
	ClassName  string
	IconSm     HICON
}

func RegisterClassEx(p *RegisterClassExParam) bool {
	type WNDCLASSEX struct {
		size       uint32
		style      uint32
		wndProc    uintptr
		clsExtra   int32
		wndExtra   int32
		instance   HINSTANCE
		icon       HICON
		cursor     HCURSOR
		background HBRUSH
		menuName   uintptr
		className  uintptr
		iconSm     HICON
	}

	var v WNDCLASSEX
	v = WNDCLASSEX{
		size:       uint32(unsafe.Sizeof(v)),
		style:      p.Style,
		wndProc:    p.WndProc,
		clsExtra:   p.ClsExtra,
		wndExtra:   p.WndExtra,
		instance:   HINSTANCE(GetModuleHandle("")),
		icon:       p.Icon,
		cursor:     p.Cursor,
		background: p.Background,
		menuName:   StringToUintptr(p.MenuName),
		className:  StringToUintptr(p.ClassName),
		iconSm:     p.IconSm,
	}

	var ret uintptr
	ret, _, lastError = procRegisterClassEx.Call(uintptr(unsafe.Pointer(&v)))

	return ret != 0
}

func ReleaseDC(h HWND, hdc HDC) bool {
	var ret uintptr
	ret, _, lastError = procReleaseDC.Call(uintptr(h), uintptr(hdc))

	return ret == 1
}

func SendMessage(m *MSG) LRESULT {
	var ret uintptr
	ret, _, lastError = procSendMessage.Call(uintptr(m.HWnd), uintptr(m.Msg),
		uintptr(m.WParam), uintptr(m.LParam))

	return LRESULT(ret)
}

func SendDlgItemMessage(m *MSG, id int) LRESULT {
	var ret uintptr
	ret, _, lastError = procSendDlgItemMessage.Call(uintptr(m.HWnd), uintptr(id),
		uintptr(m.Msg), uintptr(m.WParam), uintptr(m.LParam))

	return LRESULT(ret)
}

func SetMenu(hwnd HWND, menu HMENU) bool {
	var ret uintptr
	ret, _, lastError = procSetMenu.Call(uintptr(hwnd), uintptr(menu))

	return PtrToBool(ret)
}

func SetWindowLongPtr(h HWND, index int, value uintptr) (ret uintptr) {
	if is64Bit {
		ret, _, _ = procSetWindowLongPtr.Call(uintptr(h), uintptr(index), value)
	} else {
		ret, _, _ = procSetWindowLong.Call(uintptr(h), uintptr(index), value)
	}

	return ret
}

func ShowWindow(h HWND, cmdShow uint) bool {
	var ret uintptr
	ret, _, lastError = procShowWindow.Call(uintptr(h), uintptr(cmdShow))

	return PtrToBool(ret)
}

func TranslateMessage(p *WinMSG) bool {
	var ret uintptr
	ret, _, lastError = procTranslateMessage.Call(uintptr(unsafe.Pointer(p)))

	return PtrToBool(ret)
}

func UpdateWindow(h HWND) bool {
	var ret uintptr
	ret, _, lastError = procUpdateWindow.Call(uintptr(h))

	return PtrToBool(ret)
}

func init() {
	is64Bit = unsafe.Sizeof(uintptr(0)) == 8
}

type MSG struct {
	HWnd   HWND
	Msg    UINT
	WParam WPARAM
	LParam LPARAM
}

type WinMSG struct {
	MSG
	Time DWORD
	Pt   Point
}

func LOWORD(v uintptr) UINT {
	return UINT(v & 0xffff)
}

func HIWORD(v uintptr) UINT {
	return UINT((v >> 16) & 0xffff)
}

type CMD struct {
	Cmd        UINT
	WParamHigh UINT
	LParam     LPARAM
}

type NMHDR struct {
	HWndFrom HWND
	IdFrom   uintptr
	Code     UINT
}

// Window message constants
const (
	WM_APP                    = 32768
	WM_ACTIVATE               = 6
	WM_ACTIVATEAPP            = 28
	WM_AFXFIRST               = 864
	WM_AFXLAST                = 895
	WM_ASKCBFORMATNAME        = 780
	WM_CANCELJOURNAL          = 75
	WM_CANCELMODE             = 31
	WM_CAPTURECHANGED         = 533
	WM_CHANGECBCHAIN          = 781
	WM_CHAR                   = 258
	WM_CHARTOITEM             = 47
	WM_CHILDACTIVATE          = 34
	WM_CLEAR                  = 771
	WM_CLOSE                  = 16
	WM_COMMAND                = 273
	WM_COMMNOTIFY             = 68 /* OBSOLETE */
	WM_COMPACTING             = 65
	WM_COMPAREITEM            = 57
	WM_CONTEXTMENU            = 123
	WM_COPY                   = 769
	WM_COPYDATA               = 74
	WM_CREATE                 = 1
	WM_CTLCOLORBTN            = 309
	WM_CTLCOLORDLG            = 310
	WM_CTLCOLOREDIT           = 307
	WM_CTLCOLORLISTBOX        = 308
	WM_CTLCOLORMSGBOX         = 306
	WM_CTLCOLORSCROLLBAR      = 311
	WM_CTLCOLORSTATIC         = 312
	WM_CUT                    = 768
	WM_DEADCHAR               = 259
	WM_DELETEITEM             = 45
	WM_DESTROY                = 2
	WM_DESTROYCLIPBOARD       = 775
	WM_DEVICECHANGE           = 537
	WM_DEVMODECHANGE          = 27
	WM_DISPLAYCHANGE          = 126
	WM_DRAWCLIPBOARD          = 776
	WM_DRAWITEM               = 43
	WM_DROPFILES              = 563
	WM_ENABLE                 = 10
	WM_ENDSESSION             = 22
	WM_ENTERIDLE              = 289
	WM_ENTERMENULOOP          = 529
	WM_ENTERSIZEMOVE          = 561
	WM_ERASEBKGND             = 20
	WM_EXITMENULOOP           = 530
	WM_EXITSIZEMOVE           = 562
	WM_FONTCHANGE             = 29
	WM_GETDLGCODE             = 135
	WM_GETFONT                = 49
	WM_GETHOTKEY              = 51
	WM_GETICON                = 127
	WM_GETMINMAXINFO          = 36
	WM_GETTEXT                = 13
	WM_GETTEXTLENGTH          = 14
	WM_HANDHELDFIRST          = 856
	WM_HANDHELDLAST           = 863
	WM_HELP                   = 83
	WM_HOTKEY                 = 786
	WM_HSCROLL                = 276
	WM_HSCROLLCLIPBOARD       = 782
	WM_ICONERASEBKGND         = 39
	WM_INITDIALOG             = 272
	WM_INITMENU               = 278
	WM_INITMENUPOPUP          = 279
	WM_INPUT                  = 0X00FF
	WM_INPUTLANGCHANGE        = 81
	WM_INPUTLANGCHANGEREQUEST = 80
	WM_KEYDOWN                = 256
	WM_KEYUP                  = 257
	WM_KILLFOCUS              = 8
	WM_MDIACTIVATE            = 546
	WM_MDICASCADE             = 551
	WM_MDICREATE              = 544
	WM_MDIDESTROY             = 545
	WM_MDIGETACTIVE           = 553
	WM_MDIICONARRANGE         = 552
	WM_MDIMAXIMIZE            = 549
	WM_MDINEXT                = 548
	WM_MDIREFRESHMENU         = 564
	WM_MDIRESTORE             = 547
	WM_MDISETMENU             = 560
	WM_MDITILE                = 550
	WM_MEASUREITEM            = 44
	WM_GETOBJECT              = 0X003D
	WM_CHANGEUISTATE          = 0X0127
	WM_UPDATEUISTATE          = 0X0128
	WM_QUERYUISTATE           = 0X0129
	WM_UNINITMENUPOPUP        = 0X0125
	WM_MENURBUTTONUP          = 290
	WM_MENUCOMMAND            = 0X0126
	WM_MENUGETOBJECT          = 0X0124
	WM_MENUDRAG               = 0X0123
	WM_APPCOMMAND             = 0X0319
	WM_MENUCHAR               = 288
	WM_MENUSELECT             = 287
	WM_MOVE                   = 3
	WM_MOVING                 = 534
	WM_NCACTIVATE             = 134
	WM_NCCALCSIZE             = 131
	WM_NCCREATE               = 129
	WM_NCDESTROY              = 130
	WM_NCHITTEST              = 132
	WM_NCLBUTTONDBLCLK        = 163
	WM_NCLBUTTONDOWN          = 161
	WM_NCLBUTTONUP            = 162
	WM_NCMBUTTONDBLCLK        = 169
	WM_NCMBUTTONDOWN          = 167
	WM_NCMBUTTONUP            = 168
	WM_NCXBUTTONDOWN          = 171
	WM_NCXBUTTONUP            = 172
	WM_NCXBUTTONDBLCLK        = 173
	WM_NCMOUSEHOVER           = 0X02A0
	WM_NCMOUSELEAVE           = 0X02A2
	WM_NCMOUSEMOVE            = 160
	WM_NCPAINT                = 133
	WM_NCRBUTTONDBLCLK        = 166
	WM_NCRBUTTONDOWN          = 164
	WM_NCRBUTTONUP            = 165
	WM_NEXTDLGCTL             = 40
	WM_NEXTMENU               = 531
	WM_NOTIFY                 = 78
	WM_NOTIFYFORMAT           = 85
	WM_NULL                   = 0
	WM_PAINT                  = 15
	WM_PAINTCLIPBOARD         = 777
	WM_PAINTICON              = 38
	WM_PALETTECHANGED         = 785
	WM_PALETTEISCHANGING      = 784
	WM_PARENTNOTIFY           = 528
	WM_PASTE                  = 770
	WM_PENWINFIRST            = 896
	WM_PENWINLAST             = 911
	WM_POWER                  = 72
	WM_POWERBROADCAST         = 536
	WM_PRINT                  = 791
	WM_PRINTCLIENT            = 792
	WM_QUERYDRAGICON          = 55
	WM_QUERYENDSESSION        = 17
	WM_QUERYNEWPALETTE        = 783
	WM_QUERYOPEN              = 19
	WM_QUEUESYNC              = 35
	WM_QUIT                   = 18
	WM_RENDERALLFORMATS       = 774
	WM_RENDERFORMAT           = 773
	WM_SETCURSOR              = 32
	WM_SETFOCUS               = 7
	WM_SETFONT                = 48
	WM_SETHOTKEY              = 50
	WM_SETICON                = 128
	WM_SETREDRAW              = 11
	WM_SETTEXT                = 12
	WM_SETTINGCHANGE          = 26
	WM_SHOWWINDOW             = 24
	WM_SIZE                   = 5
	WM_SIZECLIPBOARD          = 779
	WM_SIZING                 = 532
	WM_SPOOLERSTATUS          = 42
	WM_STYLECHANGED           = 125
	WM_STYLECHANGING          = 124
	WM_SYSCHAR                = 262
	WM_SYSCOLORCHANGE         = 21
	WM_SYSCOMMAND             = 274
	WM_SYSDEADCHAR            = 263
	WM_SYSKEYDOWN             = 260
	WM_SYSKEYUP               = 261
	WM_TCARD                  = 82
	WM_THEMECHANGED           = 794
	WM_TIMECHANGE             = 30
	WM_TIMER                  = 275
	WM_UNDO                   = 772
	WM_USER                   = 1024
	WM_USERCHANGED            = 84
	WM_VKEYTOITEM             = 46
	WM_VSCROLL                = 277
	WM_VSCROLLCLIPBOARD       = 778
	WM_WINDOWPOSCHANGED       = 71
	WM_WINDOWPOSCHANGING      = 70
	WM_WININICHANGE           = 26
	WM_KEYFIRST               = 256
	WM_KEYLAST                = 264
	WM_SYNCPAINT              = 136
	WM_MOUSEACTIVATE          = 33
	WM_MOUSEMOVE              = 512
	WM_LBUTTONDOWN            = 513
	WM_LBUTTONUP              = 514
	WM_LBUTTONDBLCLK          = 515
	WM_RBUTTONDOWN            = 516
	WM_RBUTTONUP              = 517
	WM_RBUTTONDBLCLK          = 518
	WM_MBUTTONDOWN            = 519
	WM_MBUTTONUP              = 520
	WM_MBUTTONDBLCLK          = 521
	WM_MOUSEWHEEL             = 522
	WM_MOUSEFIRST             = 512
	WM_XBUTTONDOWN            = 523
	WM_XBUTTONUP              = 524
	WM_XBUTTONDBLCLK          = 525
	WM_MOUSELAST              = 525
	WM_MOUSEHOVER             = 0X2A1
	WM_MOUSELEAVE             = 0X2A3
)

// mouse button constants
const (
	MK_CONTROL  = 0x0008
	MK_LBUTTON  = 0x0001
	MK_MBUTTON  = 0x0010
	MK_RBUTTON  = 0x0002
	MK_SHIFT    = 0x0004
	MK_XBUTTON1 = 0x0020
	MK_XBUTTON2 = 0x0040
)

// Window style constants
const (
	WS_OVERLAPPED       = 0X00000000
	WS_POPUP            = 0X80000000
	WS_CHILD            = 0X40000000
	WS_MINIMIZE         = 0X20000000
	WS_VISIBLE          = 0X10000000
	WS_DISABLED         = 0X08000000
	WS_CLIPSIBLINGS     = 0X04000000
	WS_CLIPCHILDREN     = 0X02000000
	WS_MAXIMIZE         = 0X01000000
	WS_CAPTION          = 0X00C00000
	WS_BORDER           = 0X00800000
	WS_DLGFRAME         = 0X00400000
	WS_VSCROLL          = 0X00200000
	WS_HSCROLL          = 0X00100000
	WS_SYSMENU          = 0X00080000
	WS_THICKFRAME       = 0X00040000
	WS_GROUP            = 0X00020000
	WS_TABSTOP          = 0X00010000
	WS_MINIMIZEBOX      = 0X00020000
	WS_MAXIMIZEBOX      = 0X00010000
	WS_TILED            = 0X00000000
	WS_ICONIC           = 0X20000000
	WS_SIZEBOX          = 0X00040000
	WS_OVERLAPPEDWINDOW = 0X00000000 | 0X00C00000 | 0X00080000 | 0X00040000 | 0X00020000 | 0X00010000
	WS_POPUPWINDOW      = 0X80000000 | 0X00800000 | 0X00080000
	WS_CHILDWINDOW      = 0X40000000
)

// Extended window style constants
const (
	WS_EX_DLGMODALFRAME    = 0X00000001
	WS_EX_NOPARENTNOTIFY   = 0X00000004
	WS_EX_TOPMOST          = 0X00000008
	WS_EX_ACCEPTFILES      = 0X00000010
	WS_EX_TRANSPARENT      = 0X00000020
	WS_EX_MDICHILD         = 0X00000040
	WS_EX_TOOLWINDOW       = 0X00000080
	WS_EX_WINDOWEDGE       = 0X00000100
	WS_EX_CLIENTEDGE       = 0X00000200
	WS_EX_CONTEXTHELP      = 0X00000400
	WS_EX_RIGHT            = 0X00001000
	WS_EX_LEFT             = 0X00000000
	WS_EX_RTLREADING       = 0X00002000
	WS_EX_LTRREADING       = 0X00000000
	WS_EX_LEFTSCROLLBAR    = 0X00004000
	WS_EX_RIGHTSCROLLBAR   = 0X00000000
	WS_EX_CONTROLPARENT    = 0X00010000
	WS_EX_STATICEDGE       = 0X00020000
	WS_EX_APPWINDOW        = 0X00040000
	WS_EX_OVERLAPPEDWINDOW = 0X00000100 | 0X00000200
	WS_EX_PALETTEWINDOW    = 0X00000100 | 0X00000080 | 0X00000008
	WS_EX_LAYERED          = 0X00080000
	WS_EX_NOINHERITLAYOUT  = 0X00100000
	WS_EX_LAYOUTRTL        = 0X00400000
	WS_EX_NOACTIVATE       = 0X08000000
)

// ShowWindow constants
const (
	SW_HIDE            = 0
	SW_NORMAL          = 1
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_MAXIMIZE        = 3
	SW_SHOWMAXIMIZED   = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
)

// MessageBox constants
const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND
	MB_DEFBUTTON1        = 0x00000000
	MB_DEFBUTTON2        = 0x00000100
	MB_DEFBUTTON3        = 0x00000200
	MB_DEFBUTTON4        = 0x00000300
)

// Button state constants
const (
	BST_CHECKED       = 1
	BST_INDETERMINATE = 2
	BST_UNCHECKED     = 0
	BST_FOCUS         = 8
	BST_PUSHED        = 4
)

// Predefined brushes constants
const (
	COLOR_3DDKSHADOW              = 21
	COLOR_3DFACE                  = 15
	COLOR_3DHILIGHT               = 20
	COLOR_3DHIGHLIGHT             = 20
	COLOR_3DLIGHT                 = 22
	COLOR_BTNHILIGHT              = 20
	COLOR_3DSHADOW                = 16
	COLOR_ACTIVEBORDER            = 10
	COLOR_ACTIVECAPTION           = 2
	COLOR_APPWORKSPACE            = 12
	COLOR_BACKGROUND              = 1
	COLOR_DESKTOP                 = 1
	COLOR_BTNFACE                 = 15
	COLOR_BTNHIGHLIGHT            = 20
	COLOR_BTNSHADOW               = 16
	COLOR_BTNTEXT                 = 18
	COLOR_CAPTIONTEXT             = 9
	COLOR_GRAYTEXT                = 17
	COLOR_HIGHLIGHT               = 13
	COLOR_HIGHLIGHTTEXT           = 14
	COLOR_INACTIVEBORDER          = 11
	COLOR_INACTIVECAPTION         = 3
	COLOR_INACTIVECAPTIONTEXT     = 19
	COLOR_INFOBK                  = 24
	COLOR_INFOTEXT                = 23
	COLOR_MENU                    = 4
	COLOR_MENUTEXT                = 7
	COLOR_SCROLLBAR               = 0
	COLOR_WINDOW                  = 5
	COLOR_WINDOWFRAME             = 6
	COLOR_WINDOWTEXT              = 8
	COLOR_HOTLIGHT                = 26
	COLOR_GRADIENTACTIVECAPTION   = 27
	COLOR_GRADIENTINACTIVECAPTION = 28
)

// Dialog box command ids
const (
	IDOK       = 1
	IDCANCEL   = 2
	IDABORT    = 3
	IDRETRY    = 4
	IDIGNORE   = 5
	IDYES      = 6
	IDNO       = 7
	IDCLOSE    = 8
	IDHELP     = 9
	IDTRYAGAIN = 10
	IDCONTINUE = 11
	IDTIMEOUT  = 32000
)

// System commands
const (
	SC_SIZE         = 0xF000
	SC_MOVE         = 0xF010
	SC_MINIMIZE     = 0xF020
	SC_MAXIMIZE     = 0xF030
	SC_NEXTWINDOW   = 0xF040
	SC_PREVWINDOW   = 0xF050
	SC_CLOSE        = 0xF060
	SC_VSCROLL      = 0xF070
	SC_HSCROLL      = 0xF080
	SC_MOUSEMENU    = 0xF090
	SC_KEYMENU      = 0xF100
	SC_ARRANGE      = 0xF110
	SC_RESTORE      = 0xF120
	SC_TASKLIST     = 0xF130
	SC_SCREENSAVE   = 0xF140
	SC_HOTKEY       = 0xF150
	SC_DEFAULT      = 0xF160
	SC_MONITORPOWER = 0xF170
	SC_CONTEXTHELP  = 0xF180
	SC_SEPARATOR    = 0xF00F
)

// Static control styles
const (
	SS_BITMAP          = 14
	SS_BLACKFRAME      = 7
	SS_BLACKRECT       = 4
	SS_CENTER          = 1
	SS_CENTERIMAGE     = 512
	SS_EDITCONTROL     = 0x2000
	SS_ENHMETAFILE     = 15
	SS_ETCHEDFRAME     = 18
	SS_ETCHEDHORZ      = 16
	SS_ETCHEDVERT      = 17
	SS_GRAYFRAME       = 8
	SS_GRAYRECT        = 5
	SS_ICON            = 3
	SS_LEFT            = 0
	SS_LEFTNOWORDWRAP  = 0xc
	SS_NOPREFIX        = 128
	SS_NOTIFY          = 256
	SS_OWNERDRAW       = 0xd
	SS_REALSIZECONTROL = 0x040
	SS_REALSIZEIMAGE   = 0x800
	SS_RIGHT           = 2
	SS_RIGHTJUST       = 0x400
	SS_SIMPLE          = 11
	SS_SUNKEN          = 4096
	SS_WHITEFRAME      = 9
	SS_WHITERECT       = 6
	SS_USERITEM        = 10
	SS_TYPEMASK        = 0x0000001F
	SS_ENDELLIPSIS     = 0x00004000
	SS_PATHELLIPSIS    = 0x00008000
	SS_WORDELLIPSIS    = 0x0000C000
	SS_ELLIPSISMASK    = 0x0000C000
)

// Button message constants
const (
	BM_CLICK    = 245
	BM_GETCHECK = 240
	BM_GETIMAGE = 246
	BM_GETSTATE = 242
	BM_SETCHECK = 241
	BM_SETIMAGE = 247
	BM_SETSTATE = 243
	BM_SETSTYLE = 244
)

// Button notifications
const (
	BN_CLICKED       = 0
	BN_PAINT         = 1
	BN_HILITE        = 2
	BN_PUSHED        = BN_HILITE
	BN_UNHILITE      = 3
	BN_UNPUSHED      = BN_UNHILITE
	BN_DISABLE       = 4
	BN_DOUBLECLICKED = 5
	BN_DBLCLK        = BN_DOUBLECLICKED
	BN_SETFOCUS      = 6
	BN_KILLFOCUS     = 7
)

// Button style constants
const (
	BS_3STATE          = 5
	BS_AUTO3STATE      = 6
	BS_AUTOCHECKBOX    = 3
	BS_AUTORADIOBUTTON = 9
	BS_BITMAP          = 128
	BS_BOTTOM          = 0X800
	BS_CENTER          = 0X300
	BS_CHECKBOX        = 2
	BS_DEFPUSHBUTTON   = 1
	BS_GROUPBOX        = 7
	BS_ICON            = 64
	BS_LEFT            = 256
	BS_LEFTTEXT        = 32
	BS_MULTILINE       = 0X2000
	BS_NOTIFY          = 0X4000
	BS_OWNERDRAW       = 0XB
	BS_PUSHBUTTON      = 0
	BS_PUSHLIKE        = 4096
	BS_RADIOBUTTON     = 4
	BS_RIGHT           = 512
	BS_RIGHTBUTTON     = 32
	BS_TEXT            = 0
	BS_TOP             = 0X400
	BS_USERBUTTON      = 8
	BS_VCENTER         = 0XC00
	BS_FLAT            = 0X8000
)

// Predefined icon constants
const (
	IDI_APPLICATION = 32512
	IDI_HAND        = 32513
	IDI_QUESTION    = 32514
	IDI_EXCLAMATION = 32515
	IDI_ASTERISK    = 32516
	IDI_WINLOGO     = 32517
	IDI_WARNING     = IDI_EXCLAMATION
	IDI_ERROR       = IDI_HAND
	IDI_INFORMATION = IDI_ASTERISK
)

// Predefined cursor constants
const (
	IDC_ARROW       = 32512
	IDC_IBEAM       = 32513
	IDC_WAIT        = 32514
	IDC_CROSS       = 32515
	IDC_UPARROW     = 32516
	IDC_SIZENWSE    = 32642
	IDC_SIZENESW    = 32643
	IDC_SIZEWE      = 32644
	IDC_SIZENS      = 32645
	IDC_SIZEALL     = 32646
	IDC_NO          = 32648
	IDC_HAND        = 32649
	IDC_APPSTARTING = 32650
	IDC_HELP        = 32651
	IDC_ICON        = 32641
	IDC_SIZE        = 32640
)

// GetSystemMetrics constants
const (
	SM_CXSCREEN             = 0
	SM_CYSCREEN             = 1
	SM_CXVSCROLL            = 2
	SM_CYHSCROLL            = 3
	SM_CYCAPTION            = 4
	SM_CXBORDER             = 5
	SM_CYBORDER             = 6
	SM_CXDLGFRAME           = 7
	SM_CYDLGFRAME           = 8
	SM_CYVTHUMB             = 9
	SM_CXHTHUMB             = 10
	SM_CXICON               = 11
	SM_CYICON               = 12
	SM_CXCURSOR             = 13
	SM_CYCURSOR             = 14
	SM_CYMENU               = 15
	SM_CXFULLSCREEN         = 16
	SM_CYFULLSCREEN         = 17
	SM_CYKANJIWINDOW        = 18
	SM_MOUSEPRESENT         = 19
	SM_CYVSCROLL            = 20
	SM_CXHSCROLL            = 21
	SM_DEBUG                = 22
	SM_SWAPBUTTON           = 23
	SM_RESERVED1            = 24
	SM_RESERVED2            = 25
	SM_RESERVED3            = 26
	SM_RESERVED4            = 27
	SM_CXMIN                = 28
	SM_CYMIN                = 29
	SM_CXSIZE               = 30
	SM_CYSIZE               = 31
	SM_CXFRAME              = 32
	SM_CYFRAME              = 33
	SM_CXMINTRACK           = 34
	SM_CYMINTRACK           = 35
	SM_CXDOUBLECLK          = 36
	SM_CYDOUBLECLK          = 37
	SM_CXICONSPACING        = 38
	SM_CYICONSPACING        = 39
	SM_MENUDROPALIGNMENT    = 40
	SM_PENWINDOWS           = 41
	SM_DBCSENABLED          = 42
	SM_CMOUSEBUTTONS        = 43
	SM_CXFIXEDFRAME         = SM_CXDLGFRAME
	SM_CYFIXEDFRAME         = SM_CYDLGFRAME
	SM_CXSIZEFRAME          = SM_CXFRAME
	SM_CYSIZEFRAME          = SM_CYFRAME
	SM_SECURE               = 44
	SM_CXEDGE               = 45
	SM_CYEDGE               = 46
	SM_CXMINSPACING         = 47
	SM_CYMINSPACING         = 48
	SM_CXSMICON             = 49
	SM_CYSMICON             = 50
	SM_CYSMCAPTION          = 51
	SM_CXSMSIZE             = 52
	SM_CYSMSIZE             = 53
	SM_CXMENUSIZE           = 54
	SM_CYMENUSIZE           = 55
	SM_ARRANGE              = 56
	SM_CXMINIMIZED          = 57
	SM_CYMINIMIZED          = 58
	SM_CXMAXTRACK           = 59
	SM_CYMAXTRACK           = 60
	SM_CXMAXIMIZED          = 61
	SM_CYMAXIMIZED          = 62
	SM_NETWORK              = 63
	SM_CLEANBOOT            = 67
	SM_CXDRAG               = 68
	SM_CYDRAG               = 69
	SM_SHOWSOUNDS           = 70
	SM_CXMENUCHECK          = 71
	SM_CYMENUCHECK          = 72
	SM_SLOWMACHINE          = 73
	SM_MIDEASTENABLED       = 74
	SM_MOUSEWHEELPRESENT    = 75
	SM_XVIRTUALSCREEN       = 76
	SM_YVIRTUALSCREEN       = 77
	SM_CXVIRTUALSCREEN      = 78
	SM_CYVIRTUALSCREEN      = 79
	SM_CMONITORS            = 80
	SM_SAMEDISPLAYFORMAT    = 81
	SM_IMMENABLED           = 82
	SM_CXFOCUSBORDER        = 83
	SM_CYFOCUSBORDER        = 84
	SM_TABLETPC             = 86
	SM_MEDIACENTER          = 87
	SM_STARTER              = 88
	SM_SERVERR2             = 89
	SM_CMETRICS             = 91
	SM_REMOTESESSION        = 0x1000
	SM_SHUTTINGDOWN         = 0x2000
	SM_REMOTECONTROL        = 0x2001
	SM_CARETBLINKINGENABLED = 0x2002
)

// Window class styles
const (
	CS_VREDRAW         = 0x00000001
	CS_HREDRAW         = 0x00000002
	CS_KEYCVTWINDOW    = 0x00000004
	CS_DBLCLKS         = 0x00000008
	CS_OWNDC           = 0x00000020
	CS_CLASSDC         = 0x00000040
	CS_PARENTDC        = 0x00000080
	CS_NOKEYCVT        = 0x00000100
	CS_NOCLOSE         = 0x00000200
	CS_SAVEBITS        = 0x00000800
	CS_BYTEALIGNCLIENT = 0x00001000
	CS_BYTEALIGNWINDOW = 0x00002000
	CS_GLOBALCLASS     = 0x00004000
	CS_IME             = 0x00010000
	CS_DROPSHADOW      = 0x00020000
)

const CW_USEDEFAULT = ^0x7fffffff

// GetWindowLong and GetWindowLongPtr constants
const (
	GWL_EXSTYLE     = -20
	GWL_STYLE       = -16
	GWL_WNDPROC     = -4
	GWLP_WNDPROC    = -4
	GWL_HINSTANCE   = -6
	GWLP_HINSTANCE  = -6
	GWL_HWNDPARENT  = -8
	GWLP_HWNDPARENT = -8
	GWL_ID          = -12
	GWLP_ID         = -12
	GWL_USERDATA    = -21
	GWLP_USERDATA   = -21
)
