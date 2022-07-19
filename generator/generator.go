package generator

import (
  "fmt"
  "strings"
)

func getName(title *bool) string {
  first := getData("firstnames")
  last := getData("lastnames")

  if *title {
    title := getTitle()
    first = fmt.Sprintf("%s %s", title, first)
  }

  if oneInNChance(2) {
    middle := getData("firstnames")
    first = fmt.Sprintf("%s %s", first, middle)
  }

  return fmt.Sprintf("%s %s", first, last)
}

func getTitle() string {
  titles := []string{
    "Mr",
    "Miss",
    "Mrs",
    "Ms",
    "Mx",
    "Dr",
  }
  return getRandomItem(titles)
}

func getEmail() string {
  first := getData("firstnames")
  last := getData("lastnames")

  if oneInNChance(10) {
    email := fmt.Sprintf("%s@%s.com", first, last)
    return strings.ToLower(email)
  }

  joiner := getEmailJoiner()
  domain := getData("domains")
  email := fmt.Sprintf("%s%s%s@%s", first, joiner, last, domain)
  return strings.ToLower(email)
}

func getEmailJoiner() string {
  joiners := []string{
    ".",
    "_",
    "-",
    "",
  }
  return getRandomItem(joiners)
}

func getMobileNumber() string {
  prefixes := []string{
    "0",
    "+44",
  }
  number := rangeIn(100000000, 999999999)
  return fmt.Sprintf("%s7%d", getRandomItem(prefixes), number)
}

func getPostcode() string {
  return getData("postcodes")
}

func getNationalInsuranceNumber() string {
  prefix := getNationalInsurancePrefix()
  d1 := rangeIn(0, 99)
  d2 := rangeIn(0, 99)
  d3 := rangeIn(0, 99)
  suffix := randStringRunes([]rune("ABCD"), 1)
  return fmt.Sprintf("%s %02d %02d %02d %s", prefix, d1, d2, d3, suffix)
}

func getNationalInsurancePrefix() string {
  disallowedPrefixes := []string{
    "BG",
    "GB",
    "NK",
    "KN",
    "TN",
    "NT",
    "ZZ",
  }

  for {
    first := randStringRunes([]rune("ABCEGHJKLMNOPRSTWXYZ"), 1)
    second := randStringRunes([]rune("ABCEGHJKLMNPRSTWXYZ"), 1)
    prefix := fmt.Sprintf("%s%s", first, second)
    if !stringInSlice(prefix, disallowedPrefixes) {
      return prefix
    }
  }
}
