package script_runner

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"reflect"
)

// var (
// back  backend.Local = backend.Local{}
// shell powershell.Shell
// )

// Takes shell command, executes it and returns output as string
func RunLinuxCommand(command string) string {
	var (
		output string = ""
	)
	fmt.Println("Command:", command)
	cmd := exec.Command("/bin/sh", "-c", command)

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanOut := bufio.NewScanner(stdout)
	scanOut.Split(bufio.ScanWords)
	for scanOut.Scan() {
		m := scanOut.Text()
		output = output + m + " " // Spaces have to be added manuallyss
	}
	err := cmd.Wait()
	if err != nil {
		log.Fatalln("ERR command failed:", command, reflect.TypeOf(err), err)
	}

	fmt.Println("Output:", string(output))

	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalln("ERR Error while running command:", err, output)
	// }

	return string(output)
}

// Takes path to shell script, executes it and returns output as string
func RunLinuxScriptAt(scriptPath string) string {
	cmd := exec.Command("/bin/sh", scriptPath)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalln("ERR Error while running script:", err)
	}

	return string(output)
}

// Takes powershell command, executes it and returns output as string
func RunWindowsCommand(command string) string {
	// var (
	// 	err    error
	// 	stdout string
	// 	stderr string
	// )

	log.Fatalln("ERR RunWindowsCommand not implemented.")

	// // start a local powershell process
	// shell, err = powershell.New(&back)
	// if err != nil {
	// 	log.Println("ERR Couldn't create powershell process:", err)
	// }
	// defer shell.Exit()

	// // ... and interact with it
	// // stdout, stderr, err = shell.Execute("Get-WmiObject -Class Win32_Processor")
	// stdout, stderr, err = shell.Execute(command)
	// if err != nil {
	// 	log.Println("ERR Error while running powershell command:", err, stderr)
	// }

	// return stdout
	return ""
}

// Takes path to powershell script, executes it and returns output as string
func RunWindowsScriptAt(scriptPath string) string {
	// var (
	// 	err    error
	// 	stdout string
	// 	stderr string
	// )

	log.Fatalln("ERR RunWindowsScriptAt not implemented.")

	// // start a local powershell process
	// shell, err = powershell.New(&back)
	// if err != nil {
	// 	log.Println("ERR Couldn't create powershell process:", err)
	// }
	// defer shell.Exit()

	// // ... and interact with it
	// stdout, stderr, err = shell.Execute(scriptPath)
	// if err != nil {
	// 	log.Println("ERR Error while running powershell command:", err, stderr)
	// }

	// return stdout
	return ""
}
