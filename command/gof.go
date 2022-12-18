package main

type Command interface {
	run()
}

type Logger struct {
	Message string
}

func (l *Logger) run() {
	println("Logging: " + l.Message)
}

type FileSaver struct {
	Message string
}

func (f *FileSaver) run() {
	println("Logging: " + f.Message)
}

type Mailer struct {
	Message string
}

func (m *Mailer) run() {
	println("Sending: " + m.Message)
}

func executeGOF(task []Command) {
	for _, cmd := range task {
		cmd.run()
	}
}

//func main() {
//	var tasks []Command
//	logger := &Logger{Message: "Hi"}
//	fileSaver := &FileSaver{Message: "Cheers"}
//	mailer := &Mailer{Message: "Bye"}
//
//	tasks = append(tasks, logger, fileSaver, mailer)
//
//	execute(tasks)
//}
