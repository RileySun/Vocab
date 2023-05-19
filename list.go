package main

import(
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)


type List struct {
	Index int
	Items []*Quiz
	LaunchQuiz func(q *Quiz)
}

func NewList() *List {
	list := &List{
		Index:0,
	}
	list.Items = list.LoadQuizes()
	return list
}

func (l *List) Render() *fyne.Container {
	list := widget.NewList(
		func() int {
			return len(l.Items)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("Template", func() {})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {			
			//Set Name
			o.(*widget.Button).Text = l.Items[i].Name
			o.(*widget.Button).OnTapped = func() {
				l.LaunchQuiz(l.Items[i])
			}
			o.(*widget.Button).Refresh()
		},
	)
	
	return container.NewBorder(nil, nil, nil, nil, list)
}