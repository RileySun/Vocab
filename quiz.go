package main

import(
	"strings"
	"strconv"
	
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

type Quiz struct {
	Name string
	Current int
	Cards []*Card
	Score int
	Practice bool
	
	Question *canvas.Text
	Entry *widget.Entry
	Answer *canvas.Text
	Button *widget.Button
	
	EndQuiz func(q *Quiz)
	QuitQuiz func()
}

type Card struct {
	Question string 
	Answer string
}

func (l *List) LoadQuizes() []*Quiz {
	var quizzes []*Quiz
	
	quizList := LoadAllData()
	
	for _, quiz := range quizList {
		quiz := &Quiz {
			Name:strings.Split(quiz, ".")[0],
			Current:0,
			Score:0,
			Cards:LoadCardsFromJSON(quiz),
			Practice:false,
			
			
			Question:canvas.NewText("Question", WHITE),
			Entry:widget.NewEntry(),
			Answer:canvas.NewText("", WHITE),
		}
		quiz.Button = widget.NewButton("Submit", quiz.SubmitAnswer)
		quiz.Question.Alignment = 1
		quiz.Answer.Alignment = 1
		quiz.Question.TextSize = 38
		quiz.Answer.TextSize = 38
		quizzes = append(quizzes, quiz)
	}
	
	return quizzes
}

func (q *Quiz) Render() *fyne.Container { 
	q.Question.Text = q.Cards[q.Current].Question
	
	if q.Practice {
		q.Entry.Hidden = true
	} else {
		q.Entry.Hidden = false
	}
	
	buttons := container.NewBorder(nil, nil, nil, q.Button, q.Entry)
	vbox := container.New(layout.NewVBoxLayout(), q.Question, q.Answer)
	center := container.New(layout.NewCenterLayout(), vbox)
	final := container.NewBorder(nil, buttons, nil, nil, center)
	
	return final
}

func (q *Quiz) SubmitAnswer() {
	if !q.Practice {
		if q.Entry.Text == q.Cards[q.Current].Answer {
			q.Score++
		}
	}
	q.Answer.Text = q.Cards[q.Current].Answer
	q.Entry.Disable()
	q.Button.Text = "Next"
	q.Button.OnTapped = q.NextQuestion
	q.Button.Refresh()
}

func (q *Quiz) NextQuestion() {
	if q.Current < len(q.Cards) - 1 {
		q.Current++
		q.Question.Text = q.Cards[q.Current].Question
		q.Reset()
		window.Canvas().Focus(q.Entry)
	} else {
		q.Reset()
		q.EndQuiz(q)
		return
	}	
}

func (q *Quiz) Reset() {
	q.Answer.Text = ""
	q.Entry.Text = ""
	q.Entry.Enable()
	q.Button.Text = "Submit"
	q.Button.OnTapped = q.SubmitAnswer
	q.Button.Refresh()
}

func (q *Quiz) Result() *fyne.Container {
	label := canvas.NewText("You have completed the quiz with a score of:", WHITE)
	label.Alignment = 1
	label.TextSize = 30
	scoreFinal := int((float64(q.Score)/float64(len(q.Cards))) * 100)
	score := canvas.NewText(strconv.Itoa(scoreFinal) + "%", WHITE)
	score.Alignment = 1
	score.TextSize = 30
	button := widget.NewButton("End", func() {
		q.Score = 0
		q.Current = 0
		q.QuitQuiz()
	})
	vbox := container.New(layout.NewVBoxLayout(), label, score)
	center := container.New(layout.NewCenterLayout(), vbox)
	final := container.NewBorder(nil, button, nil, nil, center)
	return final
}