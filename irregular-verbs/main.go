package main

import (
  "fmt"
  "net/http"
  "os"
  "path/filepath"
  "io/ioutil"
  "encoding/xml"
)

type Data struct {
  URL string // a url
  FileName string // a file name
  XML []byte // raw XML
  Verbs Verbs // unmarshal XML
}

type Verb struct {
  XMLName xml.Name `xml:"verb"`
  Infinitive string //`xml:"infinitive"`
  PastSimple string //`xml:"pastSimple"`
  PastParticiple string //`xml:"pastParticiple"`
  Translation string //`xml:"translation"`
}

type Verbs struct {
  XMLName xml.Name `xml:"verbs"`
  Verbs []Verb `xml:"verb"`
}

func printHelp() {
  fmt.Printf("Use: ./%s <run|fetch>\n", filepath.Base(os.Args[0]))
}

func runAppFetch(d Data) error {
  fmt.Printf("Fetch XML from %s\n", d.URL)
  err := d.FetchHTTP()
  if err != nil {
    return err
  }
  err = d.SaveBase()
  if err != nil {
    return err
  }
  fmt.Println("Fetch XML and save - done!")
  return nil
}

func runApp(d Data) error {
  if b := d.ExistFile(); b == false {
    err := runAppFetch(d)
    if err != nil {
      return err
    }
  }
  err := d.FetchFile()
  if err != nil {
    fmt.Printf("Fetch %v\n", err)
    return err
  }
  fmt.Printf("%s\n", d.Verbs)
  return nil
}

func (d *Data) ExistFile() bool {
   if _, err := os.Stat(d.FileName); os.IsNotExist(err) {
     return false
   }
  return true
}

func (d *Data) FetchHTTP() error {
  res, err := http.Get(d.URL)
  if err != nil {
    return err
  }
  d.XML, err = ioutil.ReadAll(res.Body)
  defer res.Body.Close()
  if err != nil {
    return err
  }
  return d.fetch()
}

func (d *Data) FetchFile() error {
  var err error
  d.XML, err = ioutil.ReadFile(d.FileName)
  if err != nil {
    return err
  }
  return d.fetch()
}

func (d *Data) fetch() error {
  err := xml.Unmarshal(d.XML, &d.Verbs)
  if err != nil {
    return err
  }
  return nil
}

func (d *Data) SaveBase() error {
  saveXML, err := xml.Marshal(&d.Verbs)
  if err != nil {
    return err
  }
  err = ioutil.WriteFile(d.FileName, saveXML, 0644)
  if err != nil {
    return err
  }
  return nil
}

func main() {
  if len(os.Args) < 2 {
    printHelp()
    os.Exit(1)
  }
  data := Data{
    "http://spacepilot.ru/irregular-verbs/verbs.xml",
    "./verbs.xml",
    nil,
    Verbs{},
  }
  switch os.Args[1] {
  case "run":
    err := runApp(data)
    if err != nil {
      fmt.Printf("Error: %s\n", err)
      os.Exit(1)
    }
  case "fetch":
    err := runAppFetch(data)
    if err != nil {
      fmt.Printf("Error: %s\n", err)
      os.Exit(1)
    }
  default:
    printHelp()
  }
  os.Exit(0)
}
