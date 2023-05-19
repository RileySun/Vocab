package main

import(
	"log"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var fyneApp fyne.App
var window fyne.Window
var content *fyne.Container

var list *List
var quiz *Quiz

func setup() {
	fyneApp = app.NewWithID("com.sunshine.vocab")
	window = fyneApp.NewWindow("Vocab")
	
	list = NewList()
	list.LaunchQuiz = LaunchQuiz
}

func main() {
	log.Println("Loading Vocab")
	setup()
	log.Println("Loading Complete")
	
	content = container.NewBorder(nil, nil, nil, nil, list.Render())
	
	window.SetContent(content)
	
	window.CenterOnScreen()
	
	window.Resize(fyne.NewSize(720, 480))
	window.SetFixedSize(true)
	
	window.SetMaster()
	
	window.ShowAndRun()
}


func LaunchQuiz(quiz *Quiz) {
	quiz = RandomizeQuiz(quiz)
	quiz.QuitQuiz = QuitQuiz
	quiz.EndQuiz = EndQuiz
	content = quiz.Render()
	window.SetContent(content)
	window.Canvas().Focus(quiz.Entry)
}

func EndQuiz(quiz *Quiz) {
	content = quiz.Result()
	window.SetContent(content)
}

func QuitQuiz() {
	content = list.Render()
	window.SetContent(content)
}