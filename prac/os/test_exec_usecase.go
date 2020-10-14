package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	// find 'go' executable path
	// means 'which go'?
	goExecPath, err := exec.LookPath("go")
	if err == nil {
		fmt.Printf("exec.LookPath => %v\n", goExecPath)
	}

	// Construct 'go version'
	goVerCmd := &exec.Cmd{
		Path:   goExecPath,
		Args:   []string{goExecPath, "version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout, // easy to see err
	}

	fmt.Printf("exec.Cmd => %s\n", goVerCmd.String())

	if err = goVerCmd.Run(); err != nil {
		fmt.Println(err)
	}

	// Run Cmd with start and wait
	sleep3Cmd := &exec.Cmd{
		Path:   "./sleep_3.sh",
		Args:   []string{"./sleep_3.sh", "3"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	sleep3Cmd.Start()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Waitting for sleep 3 command to finish
	sleep3Cmd.Wait()

	// Construct command
	vCmd := exec.Command("go", "version")
	vCmd.Stderr = os.Stdout
	vCmd.Stdout = os.Stdout

	if err = vCmd.Run(); err != nil {
		fmt.Println(err)
	}

	// Command Output
	oCmd := exec.Command("go", "version")

	if output, err := oCmd.Output(); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("exec.Output => %s\n", output)
	}

}
