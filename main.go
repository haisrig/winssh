package main

import (
	"io"
	"log"
	"os/exec"
	"runtime"

	"github.com/gliderlabs/ssh"
)

func main() {
	handler := func(session ssh.Session) {
		log.Printf("Request recived from: %s", session.RemoteAddr())
		cmd := exec.Command("bash")
		if "windows" == runtime.GOOS {
			cmd = exec.Command("cmd")
		}
		configureTerminalIO(cmd, session)
		cmd.Run()
	}
	passwordAuth := ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
		return ("foo" == ctx.User() && "bar" == password)
	})

	log.Println("Listening on port 2200...")
	ssh.ListenAndServe(":2200", handler, passwordAuth)
}

func configureTerminalIO(cmd *exec.Cmd, s ssh.Session) {
	stdin, _ := cmd.StdinPipe()
	go func() {
		_, err := io.Copy(stdin, s)
		if err != nil {
			log.Println("Error while sending the data to stdin")
		}
	}()
	stdout, _ := cmd.StdoutPipe()
	go func() {
		_, err := io.Copy(s, stdout)
		if err != nil {
			log.Println("Error while sending the data to session")
		}
	}()
	stderr, _ := cmd.StderrPipe()
	go func() {
		_, err := io.Copy(s, stderr)
		if err != nil {
			log.Println("Error while sending the data to session")
		}
	}()
}
