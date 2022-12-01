package wholeness

import (
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

func (d *driver) RunTUI() {
	var app = tview.NewApplication()

	table := tview.NewTable()
	table.SetBorder(true).SetBorderColor(tcell.ColorWhite)

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

	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
