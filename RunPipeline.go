package main

import(
  "strings"
  "fmt"
  "os"
  "os/exec"
  "bufio"
  "strconv"

)


func CreateCommand(input string) *exec.Cmd{
  items := strings.Fields(input)
  command := items[0]
  args := items[1:]
  fmt.Printf("creating %s %s command \n", command, args[0])

  cmd := exec.Command(command, args...)
  return cmd
}

func RunCommand(cmd *exec.Cmd){
  cmd.Run()
  cmd.Wait()
}

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

func CheckError(err error) {
  if err!=nil{
    fmt.Println(err)
    os.Exit(1)
  }
}

//Takes in a sam filename and a genome length
//returns true if @SN is found and the value of LN is equal to LN on the first line of the sam file
func CheckSamFile(memOutputFile string , LN int) bool {
  file,err := os.Open(memOutputFile)
  CheckError(err)
  scanner := bufio.NewScanner(file)
  scanner.Scan()
  words := strings.Split(scanner.Text(),"	")
  genomeLength, err := strconv.Atoi((strings.Split(words[2],":")[1]))
  if genomeLength==LN{
    return true
  }
  return false
}

//takes a genomeFile and returns the amount of nucleotides in it.
func GetGenomeLength(genomeFile string) int {
  return 4411532
}

func main() {
  reference := "Mycobacterium_tuberculosis_h37rv.ASM19595v2.dna.chromosome.Chromosome.fa"
  LN:=GetGenomeLength(reference)
  //index genome
  bwaIndex := CreateCommand("bwa index "+reference)
  RunCommand(bwaIndex)
  //get Filename reads
  forwardReads := "A70376.fastq"
  //reverseReads := "A70376_2.fastq"
  memOutputFile := strings.Split(forwardReads,".")[0]+".sam"
  bwaMem := CreateCommand("bwa mem Mycobacterium_tuberculosis_h37rv.ASM19595v2.dna.chromosome.Chromosome.fa rpob.fa")
  OutputCommandToFile(bwaMem, memOutputFile)
  if !CheckSamFile(memOutputFile, LN){
    fmt.Println("BWA mem Failed")
    os.Exit(1)
  }

}
