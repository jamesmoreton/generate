package generator

import (
  "fmt"
  "bufio"
  "math/rand"
  "os"
  "path/filepath"
  "time"
)

// Fetches data point in file in O(n) time and constant memory,
// with each line have 1/N chance of being selected.
func getData(filename string) string {
  file := getFile(filename)
  scanner := bufio.NewScanner(file)

  randsource := rand.NewSource(time.Now().UnixNano())
  randgenerator := rand.New(randsource)

  lineNum := 1
  var pick string
  for scanner.Scan() {
    line := scanner.Text()
    roll := randgenerator.Intn(lineNum)
    if roll == 0 {
      pick = line
    }
    lineNum += 1
  }

  return pick
}

func getFile(filename string) *os.File {
  ex, err := os.Executable()
  if err != nil {
    panic(err)
  }
  path := filepath.Join(filepath.Dir(ex), fmt.Sprintf("generator/data/%s.txt", filename))
  file, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  return file
}

func oneInNChance(n int) bool {
  rand.Seed(time.Now().UnixNano())
  if rand.Intn(n) == 1 {
    return true
  }
  return false
}

func getRandomItem(list []string) string {
  rand.Seed(time.Now().UnixNano())
  return list[rand.Intn(len(list))]
}

func rangeIn(low, hi int) int {
  rand.Seed(time.Now().UnixNano())
  return low + rand.Intn(hi-low)
}

func randStringRunes(letters []rune, n int) string {
  rand.Seed(time.Now().UnixNano())
  b := make([]rune, n)
  for i := range b {
    b[i] = letters[rand.Intn(len(letters))]
  }
  return string(b)
}

func stringInSlice(a string, list[]string) bool {
  for _, b := range list {
    if b == a {
      return true
    }
  }
  return false
}

func validateRange(n int, low int, high int){
  if n < low || n > high {
    fmt.Printf("Please choose a valid count between %d and %d\n", low, high)
    os.Exit(1)
  }
}
