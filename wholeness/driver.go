package wholeness

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type driver struct {
	m *Model
	w World
}

func NewDriver(m *Model, w World) *driver {
	d := &driver{m, w}
	d.init()
	return d
}

func (d *driver) init() {
	d.m.BigBang(d.w)
}

func (d *driver) RunConsole() {
	d.w.RenderDebugDump()

	for i := 0; i < 10; i++ {
		d.w.RenderDebugDump()
		d.w.Tick()
	}
}

var helpText = `Instructions:
- Press "Ctrl-R" to Run one step
- Press "Ctrl-T" to run many steps ("Turbo")
- Press "Ctrl-C" to exit
`

func (d *driver) RunTUI() {
	var app = tview.NewApplication()

	table := tview.NewTable()
	table.SetBorder(true).SetBorderColor(tcell.ColorWhite)
	title := fmt.Sprintf("Wholeness -- %s", d.m.name)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle(title), 0, 1, false).
		AddItem(table, 0, 12, true).
		AddItem(tview.NewTextView().SetText(helpText), 5, 1, false)

	step := func(num int) {
		for i := 0; i < num; i++ {
			d.w.RenderTable(table)
			d.w.Tick()
			app.Draw()
		}
	}

	// Initial setup
	go func() {
		d.w.RenderTable(table)
		app.Draw()
	}()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// "Run"
		if event.Key() == tcell.KeyCtrlR {
			go step(1)
			return nil
		}
		// "Turbo"
		if event.Key() == tcell.KeyCtrlT {
			go step(5)
			return nil
		}
		return event
	})

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
