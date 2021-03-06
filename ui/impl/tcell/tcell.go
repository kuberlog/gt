package tcell

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/kuberlog/gt/ui"
)

type TCellUI struct {
	screen tcell.Screen
}

func Init() *TCellUI {
	screen, errors := tcell.NewScreen()
	if errors != nil {
		fmt.Fprintf(os.Stderr, "Failed to create screen")
		os.Exit(1)
	}
	if e := screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize screen")
		os.Exit(1)
	}

	screen.SetStyle(tcell.StyleDefault)
	screen.Clear()
	return &TCellUI{screen: screen}
}

func (ui *TCellUI) SetContent(x int, y int, runeVal rune) {
	ui.screen.SetContent(x, y, runeVal, []rune{}, tcell.StyleDefault)
}

func (ui *TCellUI) ScreenSize() (int, int) {
	return ui.screen.Size()
}

func (ui *TCellUI) Show() {
	ui.screen.Show()
}

func (ui *TCellUI) Fini() {
	ui.screen.Fini()
}

func (ui *TCellUI) PollEvent() ui.InputEvent {
	return ui.screen.PollEvent()
}

func (ui *TCellUI) ShowCursor(col int, row int) {
	ui.screen.ShowCursor(col, row)
}

func (ui *TCellUI) HideCursor() {
	ui.screen.HideCursor()
}
