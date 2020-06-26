package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
  cmd := exec.Command("/bin/sh")

  /* */
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  cmd.Env = []string{"# "}
  /* */
  cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER,
    UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
  }

  if err := cmd.Run() ; err != nil {
    fmt.Println("Err ",err)
    os.Exit(1)
  }


}
