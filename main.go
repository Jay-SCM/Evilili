package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type CustomLabel struct {
	widget.Label
	isDragging bool
	startX     float32
	startY     float32
}

func (l *CustomLabel) MouseDown(e *desktop.MouseEvent) {
	l.isDragging = true
	l.startX = e.Position.X
	l.startY = e.Position.Y
}

func (l *CustomLabel) MouseUp(e *desktop.MouseEvent) {
	l.isDragging = false
}

func (l *CustomLabel) MouseMoved(e *desktop.MouseEvent) {
	if l.isDragging {
		deltaX := e.Position.X - l.startX
		deltaY := e.Position.Y - l.startY
		l.Move(l.Position().Add(fyne.NewPos(deltaX, deltaY)))
		l.startX = e.Position.X
		l.startY = e.Position.Y
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("System Stats Monitor")

	cpuLabel := widget.NewLabel("CPU Usage: Loading...")
	memLabel := widget.NewLabel("Memory Usage: Loading...")

	draggable := &CustomLabel{}
	draggable.SetText("Drag me")
	draggable.Resize(fyne.NewSize(100, 50))
	draggable.Move(fyne.NewPos(100, 100)) // Initial position

	content := container.NewVBox(cpuLabel, memLabel)
	mainContainer := container.NewWithoutLayout(draggable, content)

	w.SetContent(mainContainer)
	w.Resize(fyne.NewSize(800, 600))
	w.Show()

	go func() {
		for {
			time.Sleep(1 * time.Second)

			cpuPercent, err := cpu.Percent(0, false)
			if err == nil {
				cpuLabel.SetText(fmt.Sprintf("CPU Usage: %.2f%%", cpuPercent[0]))
			}

			memStats, err := mem.VirtualMemory()
			if err == nil {
				memLabel.SetText(fmt.Sprintf("Memory Usage: %.2f%% (%.2f GB used out of %.2f GB)",
					memStats.UsedPercent, float64(memStats.Used)/1e9, float64(memStats.Total)/1e9))
			}
		}
	}()

	a.Run()
}

// package main

// import (
// 	"fmt"
// 	"time"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/driver/desktop"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/shirou/gopsutil/cpu"
// 	"github.com/shirou/gopsutil/mem"
// )

// type CustomLabel struct {
// 	widget.Label
// 	isDragging bool
// 	startX     float32
// 	startY     float32
// }

// func (l *CustomLabel) MouseDown(e *desktop.MouseEvent) {
// 	l.isDragging = true
// 	l.startX = e.Position.X
// 	l.startY = e.Position.Y
// }

// func (l *CustomLabel) MouseUp(e *desktop.MouseEvent) {
// 	l.isDragging = false
// }

// func (l *CustomLabel) MouseMoved(e *desktop.MouseEvent) {
// 	if l.isDragging {
// 		deltaX := e.Position.X - l.startX
// 		deltaY := e.Position.Y - l.startY
// 		l.Move(l.Position().Add(fyne.NewPos(deltaX, deltaY)))
// 		l.startX = e.Position.X
// 		l.startY = e.Position.Y
// 	}
// }

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("System Stats Monitor")

// 	cpuLabel := widget.NewLabel("CPU Usage: Loading...")
// 	memLabel := widget.NewLabel("Memory Usage: Loading...")

// 	draggable := &CustomLabel{}
// 	draggable.SetText("Drag me")
// 	draggable.Resize(fyne.NewSize(100, 50))
// 	draggable.Move(fyne.NewPos(100, 100)) // Initial position

// 	content := container.NewVBox(cpuLabel, memLabel)
// 	mainContainer := container.NewWithoutLayout(draggable, content)

// 	w.SetContent(mainContainer)
// 	w.Resize(fyne.NewSize(800, 600))
// 	w.ShowAndRun()

// 	go func() {
// 		for {
// 			cpuPercent, _ := cpu.Percent(0, false)
// 			cpuLabel.SetText(fmt.Sprintf("CPU Usage: %.2f%%", cpuPercent[0]))

// 			memStats, _ := mem.VirtualMemory()
// 			memLabel.SetText(fmt.Sprintf("Memory Usage: %.2f%% (%.2f GB used out of %.2f GB)",
// 				memStats.UsedPercent, float64(memStats.Used)/1e9, float64(memStats.Total)/1e9))

// 			time.Sleep(1 * time.Second)
// 		}
// 	}()
// }

// package main

// import (
// 	"fmt"
// 	"os/exec"
// 	"runtime"
// 	"time"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/driver/desktop"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/shirou/gopsutil/cpu"
// 	"github.com/shirou/gopsutil/mem"
// 	"golang.org/x/sys/windows"
// )

// var (
// 	user32                         = windows.NewLazySystemDLL("user32.dll")
// 	procSetWindowLong              = user32.NewProc("SetWindowLongW")
// 	procGetWindowLong              = user32.NewProc("GetWindowLongW")
// 	procSetLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes")
// )

// const (
// 	GWL_EXSTYLE       = 0xFFFFFFFF // Use hexadecimal for negative constants
// 	WS_EX_LAYERED     = 0x00080000
// 	WS_EX_TRANSPARENT = 0x00000020
// 	LWA_ALPHA         = 0x00000002
// )

// // CustomLabel is a custom widget that handles mouse events.
// type CustomLabel struct {
// 	widget.Label
// 	isDragging  bool
// 	lastMouseX  int
// 	lastMouseY  int
// }

// func (l *CustomLabel) MouseDown(e *desktop.MouseEvent) {
// 	l.isDragging = true
// 	l.lastMouseX = int(e.Position.X)
// 	l.lastMouseY = int(e.Position.Y)
// 	println("Mouse down")
// }

// func (l *CustomLabel) MouseUp(e *desktop.MouseEvent) {
// 	l.isDragging = false
// 	println("Mouse up")
// }

// func (l *CustomLabel) MouseMoved(e *desktop.MouseEvent) {
// 	if l.isDragging {
// 		dx := int(e.Position.X) - l.lastMouseX
// 		dy := int(e.Position.Y) - l.lastMouseY
// 		if w, ok := l.Parent().(*fyne.Window); ok {
// 			newPos := fyne.NewPos(int(w.Canvas().Size().Width)+dx, int(w.Canvas().Size().Height)+dy)
// 			w.Canvas().SetPosition(newPos)
// 			l.lastMouseX, l.lastMouseY = int(e.Position.X), int(e.Position.Y)
// 		}
// 		println("Mouse moved")
// 	}
// }

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("System Stats Monitor")

// 	// Create labels for stats
// 	cpuLabel := widget.NewLabel("CPU Usage: Loading...")
// 	memLabel := widget.NewLabel("Memory Usage: Loading...")

// 	// Create the layout and add the labels
// 	content := container.NewVBox(cpuLabel, memLabel)
// 	draggable := &CustomLabel{}
// 	draggable.SetText("Drag me")

// 	// Wrap the custom label in a container to enable mouse event handling
// 	eventHandler := container.NewMax(draggable)
// 	eventHandler.OnMouseDown = func(e *fyne.PointEvent) {
// 		draggable.MouseDown(&desktop.MouseEvent{
// 			Position: fyne.NewPos(float32(e.Position.X), float32(e.Position.Y)),
// 			Button:   e.Button,
// 		})
// 	}

// 	eventHandler.OnMouseUp = func(e *fyne.PointEvent) {
// 		draggable.MouseUp(&desktop.MouseEvent{
// 			Position: fyne.NewPos(float32(e.Position.X), float32(e.Position.Y)),
// 			Button:   e.Button,
// 		})
// 	}

// 	eventHandler.OnMouseMoved = func(e *fyne.PointEvent) {
// 		draggable.MouseMoved(&desktop.MouseEvent{
// 			Position: fyne.NewPos(float32(e.Position.X), float32(e.Position.Y)),
// 			Button:   e.Button,
// 		})
// 	}

// 	mainContainer := container.NewBorder(nil, nil, nil, nil, eventHandler, content)

// 	w.SetContent(mainContainer)
// 	w.Resize(fyne.NewSize(800, 600))
// 	w.ShowAndRun()

// 	// Fetch and update system stats periodically
// 	go func() {
// 		for {
// 			// CPU usage
// 			cpuPercent, _ := cpu.Percent(0, false)
// 			cpuLabel.SetText(fmt.Sprintf("CPU Usage: %.2f%%", cpuPercent[0]))

// 			// Memory usage
// 			memStats, _ := mem.VirtualMemory()
// 			memLabel.SetText(fmt.Sprintf("Memory Usage: %.2f%% (%.2f GB used out of %.2f GB)",
// 				memStats.UsedPercent, float64(memStats.Used)/1e9, float64(memStats.Total)/1e9))

// 			// Sleep before updating stats again
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	// Enable click-through functionality (OS-specific)
// 	go enableClickThrough(w)
// }

// // Platform-specific code to enable click-through
// func enableClickThrough(w fyne.Window) {
// 	switch runtime.GOOS {
// 	case "windows":
// 		enableClickThroughWindows(w)
// 	case "linux":
// 		enableClickThroughLinux(w)
// 	default:
// 		fmt.Println("Click-through functionality is not supported on this OS")
// 	}
// }

// // Windows-specific click-through implementation using Windows API
// func enableClickThroughWindows(w fyne.Window) {
// 	hwnd := getHWND(w) // Fetch the window handle
// 	if hwnd == 0 {
// 		fmt.Println("Failed to get window handle")
// 		return
// 	}
// 	// Set the window to allow click-through
// 	style, _, _ := procGetWindowLong.Call(hwnd, uintptr(GWL_EXSTYLE))
// 	procSetWindowLong.Call(hwnd, uintptr(GWL_EXSTYLE), style|WS_EX_LAYERED|WS_EX_TRANSPARENT)
// 	procSetLayeredWindowAttributes.Call(hwnd, 0, 0xFF, LWA_ALPHA) // Optional: Set transparency level (0xFF for full opacity)
// }

// // Linux-specific click-through using xprop and Xlib
// func enableClickThroughLinux(w fyne.Window) {
// 	winID := getLinuxWindowID(w)
// 	if winID == "" {
// 		fmt.Println("Failed to get window ID")
// 		return
// 	}
// 	// Use xprop to set the window to allow click-through
// 	exec.Command("xprop", "-id", winID, "-f", "_NET_WM_WINDOW_TYPE", "32a", "-set", "_NET_WM_WINDOW_TYPE", "_NET_WM_WINDOW_TYPE_DOCK").Run()
// }

// // Fetch the HWND for Windows
// func getHWND(w fyne.Window) uintptr {
// 	// Windows-specific code to get the HWND
// 	// Fyne does not expose the HWND directly; you may need to use other methods or libraries.
// 	return 0
// }

// // Fetch the window ID for Linux
// func getLinuxWindowID(w fyne.Window) string {
// 	// Use xwininfo or xprop to retrieve the window ID (on Linux X11).
// 	winID, err := exec.Command("xwininfo", "-name", "System Stats Monitor").Output()
// 	if err != nil {
// 		fmt.Println("Failed to get window ID:", err)
// 		return ""
// 	}
// 	return string(winID)
// }

// package main

// import (
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/driver/desktop"
// 	"fyne.io/fyne/v2/widget"
// )

// // CustomLabel is a custom widget that handles mouse events.
// type CustomLabel struct {
// 	widget.Label
// }

// func (l *CustomLabel) MouseDown(e *desktop.MouseEvent) {
// 	// Handle mouse down event
// 	println("Mouse down")
// }

// func (l *CustomLabel) MouseUp(e *desktop.MouseEvent) {
// 	// Handle mouse up event
// 	println("Mouse up")
// }

// func (l *CustomLabel) MouseMoved(e *desktop.MouseEvent) {
// 	// Handle mouse move event
// 	println("Mouse moved")
// }

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("Demo")

// 	content := &CustomLabel{}
// 	content.SetText("Drag me")

// 	container := container.NewCenter(content)

// 	w.SetContent(container)
// 	w.Resize(fyne.NewSize(800, 600))
// 	w.ShowAndRun()
// }

// package main

// import (
// 	"fmt"
// 	"os/exec"
// 	"runtime"
// 	"time"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/shirou/gopsutil/cpu"
// 	"github.com/shirou/gopsutil/mem"
// 	"golang.org/x/sys/windows"
// )

// var (
// 	user32                         = windows.NewLazySystemDLL("user32.dll")
// 	procSetWindowLong              = user32.NewProc("SetWindowLongW")
// 	procGetWindowLong              = user32.NewProc("GetWindowLongW")
// 	procSetLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes")
// )

// const (
// 	GWL_EXSTYLE       = 0xFFFFFFFF // Use hexadecimal for negative constants
// 	WS_EX_LAYERED     = 0x00080000
// 	WS_EX_TRANSPARENT = 0x00000020
// 	LWA_ALPHA         = 0x00000002
// )

// // Draggable window implementation
// var isDragging bool
// var lastMouseX, lastMouseY int

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("System Stats Monitor")

// 	// Create labels for stats
// 	cpuLabel := widget.NewLabel("CPU Usage: Loading...")
// 	memLabel := widget.NewLabel("Memory Usage: Loading...")

// 	// Create the layout and add the labels
// 	content := container.NewVBox(cpuLabel, memLabel)
// 	draggable := container.NewBorder(nil, nil, nil, nil, content)

// 	// Handle dragging of the window
// 	draggable.OnMouseDown = func(e *fyne.PointEvent) {
// 		isDragging = true
// 		lastMouseX, lastMouseY = int(e.Position.X), int(e.Position.Y)
// 	}

// 	draggable.OnMouseUp = func(e *fyne.PointEvent) {
// 		isDragging = false
// 	}

// 	draggable.OnMouseMoved = func(e *fyne.PointEvent) {
// 		if isDragging {
// 			dx := int(e.Position.X) - lastMouseX
// 			dy := int(e.Position.Y) - lastMouseY
// 			newPos := fyne.NewPos(w.Canvas().Size().Width+dx, w.Canvas().Size().Height+dy)
// 			w.Canvas().SetPosition(newPos)
// 			lastMouseX, lastMouseY = int(e.Position.X), int(e.Position.Y)
// 		}
// 	}

// 	// Set content
// 	w.SetContent(draggable)

// 	// Fetch and update system stats periodically
// 	go func() {
// 		for {
// 			// CPU usage
// 			cpuPercent, _ := cpu.Percent(0, false)
// 			cpuLabel.SetText(fmt.Sprintf("CPU Usage: %.2f%%", cpuPercent[0]))

// 			// Memory usage
// 			memStats, _ := mem.VirtualMemory()
// 			memLabel.SetText(fmt.Sprintf("Memory Usage: %.2f%% (%.2f GB used out of %.2f GB)",
// 				memStats.UsedPercent, float64(memStats.Used)/1e9, float64(memStats.Total)/1e9))

// 			// Sleep before updating stats again
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	// Enable click-through functionality (OS-specific)
// 	go enableClickThrough(w)

// 	// Show the window
// 	w.ShowAndRun()
// }

// // Platform-specific code to enable click-through
// func enableClickThrough(w fyne.Window) {
// 	switch runtime.GOOS {
// 	case "windows":
// 		enableClickThroughWindows(w)
// 	case "linux":
// 		enableClickThroughLinux(w)
// 	default:
// 		fmt.Println("Click-through functionality is not supported on this OS")
// 	}
// }

// // Windows-specific click-through implementation using Windows API
// func enableClickThroughWindows(w fyne.Window) {
// 	hwnd := getHWND(w) // Fetch the window handle
// 	if hwnd == 0 {
// 		fmt.Println("Failed to get window handle")
// 		return
// 	}
// 	// Set the window to allow click-through
// 	style, _, _ := procGetWindowLong.Call(hwnd, uintptr(GWL_EXSTYLE))
// 	procSetWindowLong.Call(hwnd, uintptr(GWL_EXSTYLE), style|WS_EX_LAYERED|WS_EX_TRANSPARENT)
// 	procSetLayeredWindowAttributes.Call(hwnd, 0, 0xFF, LWA_ALPHA) // Optional: Set transparency level (0xFF for full opacity)
// }

// // Linux-specific click-through using xprop and Xlib
// func enableClickThroughLinux(w fyne.Window) {
// 	winID := getLinuxWindowID(w)
// 	if winID == "" {
// 		fmt.Println("Failed to get window ID")
// 		return
// 	}
// 	// Use xprop to set the window to allow click-through
// 	exec.Command("xprop", "-id", winID, "-f", "_NET_WM_WINDOW_TYPE", "32a", "-set", "_NET_WM_WINDOW_TYPE", "_NET_WM_WINDOW_TYPE_DOCK").Run()
// }

// // Fetch the HWND for Windows
// func getHWND(w fyne.Window) uintptr {
// 	// Windows-specific code to get the HWND
// 	// Fyne does not expose the HWND directly; you may need to use other methods or libraries.
// 	return 0
// }

// // Fetch the window ID for Linux
// func getLinuxWindowID(w fyne.Window) string {
// 	// Use xwininfo or xprop to retrieve the window ID (on Linux X11).
// 	winID, err := exec.Command("xwininfo", "-name", "System Stats Monitor").Output()
// 	if err != nil {
// 		fmt.Println("Failed to get window ID:", err)
// 		return ""
// 	}
// 	return string(winID)
// }
