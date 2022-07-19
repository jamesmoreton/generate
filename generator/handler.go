package generator

import (
  "fmt"
  "flag"
  "os"
)

func HandleHelp(){
  fmt.Println("usage: generate <command> [<args>]")
  fmt.Println()
  fmt.Println("Available Generate commands:")
  fmt.Println()
  fmt.Printf("   %s\t\t\tPerson first, last, and sometimes middle name, with optional title\n", colourize("name"))
  fmt.Printf("   %s\t\tEmail address with variety of domains\n", colourize("email"))
  fmt.Printf("   %s\t\tMobile phone number (UK format), sometimes with area code\n", colourize("mobile"))
  fmt.Printf("   %s\t\tPostcode (UK)\n", colourize("postcode"))
  fmt.Printf("   %s\tNational Insurance number (UK)\n", colourize("nationalinsurance"))
  fmt.Println()
  fmt.Println("See generate <command> -help to see the usage of a specific subcommand.")
}

func colourize(s string) string {
  colourGreen := "\u001b[32m"
  colourReset := "\u001b[0m"
  return colourGreen + s + colourReset
}

func HandleName(nameCmd *flag.FlagSet, title *bool, count *int){
  nameCmd.Parse(os.Args[2:])
  validateRange(*count, 1, 100)
  for i := 0; i < *count; i++ {
    name := getName(title)
    fmt.Println(name)
  }
}

func HandleEmail(emailCmd *flag.FlagSet, count *int){
  emailCmd.Parse(os.Args[2:])
  validateRange(*count, 1, 100)
  for i := 0; i < *count; i++ {
    email := getEmail()
    fmt.Println(email)
  }
}

func HandleMobile(mobileCmd *flag.FlagSet, count *int){
  mobileCmd.Parse(os.Args[2:])
  validateRange(*count, 1, 100)
  for i := 0; i < *count; i++ {
    mobile := getMobileNumber()
    fmt.Println(mobile)
  }
}

func HandlePostcode(postcodeCmd *flag.FlagSet, count *int){
  postcodeCmd.Parse(os.Args[2:])
  validateRange(*count, 1, 100)
  for i := 0; i < *count; i++ {
    postcode := getPostcode()
    fmt.Println(postcode)
  }
}

func HandleNationalInsurance(nationalInsuranceCmd *flag.FlagSet, count *int){
  nationalInsuranceCmd.Parse(os.Args[2:])
  validateRange(*count, 1, 100)
  for i := 0; i < *count; i++ {
    nationalinsurance := getNationalInsuranceNumber()
    fmt.Println(nationalinsurance)
  }
}
