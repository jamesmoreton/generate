package main

import (
  "fmt"
  "flag"
  "os"

  "github.com/jamesmoreton/generate/generator"
)

func main() {
  flag.NewFlagSet("help", flag.ExitOnError)

  cntUsage := "Number of randomly generated items to return"
  nameCmd := flag.NewFlagSet("name", flag.ExitOnError)
  nameCnt := nameCmd.Int("count", 1, cntUsage)

  emailCmd := flag.NewFlagSet("email", flag.ExitOnError)
  emailCnt := emailCmd.Int("count", 1, cntUsage)

  mobileCmd := flag.NewFlagSet("mobile", flag.ExitOnError)
  mobileCnt := mobileCmd.Int("count", 1, cntUsage)

  postcodeCmd := flag.NewFlagSet("postcode", flag.ExitOnError)
  postcodeCnt := postcodeCmd.Int("count", 1, cntUsage)

  nationalInsuranceCmd := flag.NewFlagSet("nationalinsurance", flag.ExitOnError)
  nationalInsuranceCnt := nationalInsuranceCmd.Int("count", 1, cntUsage)

  if len(os.Args) < 2 {
    generator.HandleHelp()
    os.Exit(1)
  }

  switch os.Args[1] {
    case "help":
      generator.HandleHelp()
    case "name":
      generator.HandleName(nameCmd, nameCnt)
    case "email":
      generator.HandleEmail(emailCmd, emailCnt)
    case "mobile":
      generator.HandleMobile(mobileCmd, mobileCnt)
    case "postcode":
      generator.HandlePostcode(postcodeCmd, postcodeCnt)
    case "nationalinsurance":
      generator.HandleNationalInsurance(nationalInsuranceCmd, nationalInsuranceCnt)
    default:
      fmt.Printf("generate: '%v' is not a generate command. See 'generate help'.\n", os.Args[1])
      os.Exit(1)
  }
}
