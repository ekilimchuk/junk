package main

import (
  "fmt"
  "net/http"
  "os"
  "path/filepath"
  "io/ioutil"
)

type Data struct {
  url string // url
  path string // filepath
}

func printHelp() {
  fmt.Printf("Use: ./%s <run|fetch>\n", filepath.Base(os.Args[0]))
}

func runAppFetch(d Data) {
  err := d.Fetch()
  fmt.Printf("Fetch %v\n", err)
}

func runApp() {
  fmt.Println("Run")
}

func (d Data) Exist() bool {
  fmt.Println(d.path)
  return true
}

func (d Data) Fetch() error {
  res, err := http.Get(d.url)
  if err != nil {
    return err
  }
  body, err := ioutil.ReadAll(res.Body)
  defer res.Body.Close()
  if err != nil {
    return err
  }
  fmt.Printf("%s\n", body)
  return err
}

func main()  {

  data := Data{
    "http://spacepilot.ru/irregular-verbs/verbs.xml",
    "./verbs.xml",
  }

  if len(os.Args) < 2 {
    printHelp()
    os.Exit(1)
  }

  switch os.Args[1] {
  case "run":
    runApp()
  case "fetch":
    runAppFetch(data)
  default:
    printHelp()
  }
  os.Exit(0)
}
