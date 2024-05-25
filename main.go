package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	llm "github.com/go-skynet/go-llama.cpp"
)

var (
	modelPath string
	prompt    string
	timeout   int
)

func init() {
	flag.StringVar(&modelPath, "model", "", "Path to a gguf model")
	flag.StringVar(&prompt, "prompt", "", "prompt")
	flag.IntVar(&timeout, "timeout", 3, "Timeout for inference")
	flag.Parse()
}

type Response struct {
	s   string
	err error
}

func Predict(l *llm.LLama, prompt string, timeout time.Duration) (string, error) {
	response := make(chan Response)

	// Run the function in a separate goroutine
	go func() {
		r, err := l.Predict(prompt)
		response <- Response{s: r, err: err}
		close(response)
	}()

	select {
	case res := <-response:
		return res.s, res.err
	case <-time.After(timeout):
		return "", fmt.Errorf("function timed out")
	}
}

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	if modelPath == "" {
		fmt.Println("Error: model path is required.")
		flag.Usage()
		os.Exit(1)
	}

	l, err := llm.New(modelPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("\nQuestion: %s\n", prompt)
	response, err := Predict(l, prompt, time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", response)
}
