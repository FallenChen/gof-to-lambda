package main

type Action func(message string)

//	type Task interface {
//		execute(Action, string)
//	}
func log(message string) {
	println("Logging: " + message)
}

func save(message string) {
	println("Saving: " + message)
}

func send(message string) {
	println("Sending: " + message)
}

func executeLambda(tasks []func()) {
	for _, task := range tasks {
		task()
	}
}

func main() {
	tasks := []func(){func() { log("Hi") }, func() { save("Cheers") }, func() { send("Bye") }}
	executeLambda(tasks)
}
