package main

import(
  "strings"
  "fmt"
  "os/exec"
  "os"
)

/*
Creates a bash command using the input string. returns a pointer to the command
*/
func CreateCommand(input string) *exec.Cmd{
  items := strings.Fields(input)
  command := items[0]
  args := items[1:]
  if len(args)>0{
    fmt.Printf("creating %s command \n", input)
  } else {
    fmt.Printf("creating %s command \n", input)
  }
  cmd := exec.Command(command, args...)
  return cmd
}

/*
runs and waits for the command to finish
*/
func RunCommand(cmd *exec.Cmd){
  cmd.Run()
  cmd.Wait()
}

/*
OutputCommandToFile takes a command and a filename and writes the output of
that command to the filename
*/
func OutputCommandToFile(cmd *exec.Cmd, filename string){
  file, err := os.Create(filename)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  cmd.Stdout = file
  cmd.Run()
  cmd.Wait()
}

/*
WriteOutputToString takes a command and writes it toa string.
*/
func WriteOutputToString(cmd *exec.Cmd) string{
  out,err := cmd.Output()
  CheckError(err)
  return string(out)[:len(out)-1]
}

/*
UnzipFile unzips a file if it can be unzipped. otherwise it does nothing.
*/
func UnzipFile(file string) string {
  if strings.HasSuffix(file, ".gz") {
    gunzip := CreateCommand("gunzip " + file)
    RunCommand(gunzip)
  }
  return strings.TrimSuffix(file, ".gz")
}
