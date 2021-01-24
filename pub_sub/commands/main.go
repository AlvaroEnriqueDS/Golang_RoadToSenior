package main

import "os/exec"

func main() {
	app := "free"
	cmd := exec.Command(app)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	data := string(stdout)
	print(data)
}