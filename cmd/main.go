package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	log.Println("Starting process...", os.Args)

	// We use command-line arguments to distinguish between the parent and child processes.
	switch os.Args[len(os.Args)-1] {
	case "child":
		// This block runs if the program is executed with "child" as the last argument.
		// It's the code for the **child process**.
		log.Printf("Child process (PID: %d) - Parent PID: %d\n", os.Getpid(), os.Getppid())
		// The child can perform its specific tasks here.
		// For demonstration, we'll have it sleep for a bit.
		time.Sleep(30 * time.Second)
		log.Printf("Child process (PID: %d) exiting.\n", os.Getpid())
	default:
		// This block runs if the program is executed without "child" as the last argument.
		// It's the code for the **parent process**.

		// Prepare the command to run itself again, but with "child" appended to the arguments.
		cmd := exec.Command(os.Args[0], append(os.Args[1:], "child")...)

		log.Println("Executing child process with command:", cmd)

		// **Forking happens here**: Start the new process.
		err := cmd.Start()
		if err != nil {
			log.Printf("Error starting child process: %v\n", err)
			return
		}

		log.Printf("Parent process (PID: %d) - Child PID: %d\n", os.Getpid(), cmd.Process.Pid)

		// Optionally, the parent can wait for the child to finish.
		// If you remove this, the parent will exit immediately and the child will continue running.
		_, err = cmd.Process.Wait()
		if err != nil {
			log.Printf("Error waiting for child process: %v\n", err)
		}
		log.Printf("Parent process (PID: %d) exiting.\n", os.Getpid())
	}
}

/*실행방법

go build main.go

$ ./main.exe
2025/08/02 14:04:44 Starting process... [C:\Users\samsung\Desktop\linux_study\cmd\main.exe]
2025/08/02 14:04:44 Executing child process with command: C:\Users\samsung\Desktop\linux_study\cmd\main.exe child
2025/08/02 14:04:45 Parent process (PID: 13400) - Child PID: 12884
2025/08/02 14:05:15 Parent process (PID: 13400) exiting. -- 지정한 30초 만큼 대기 후 종료가 되는 것으로 보아서 자식 프로세스가 실행이 되긴 된다는 의미다.


*/
