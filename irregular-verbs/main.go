package main

import (
  "fmt"
  "net/http"
  "os"
  "path/filepath"
  "io/ioutil"
  "encoding/xml"
  "strings"
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
  fmt.Printf("Use: ./%s <run|fetch|list>\n", filepath.Base(os.Args[0]))
}

func runAppFetch(d Data) error {
  fmt.Printf("Fetch XML from %s\n", d.URL)
  if err := d.FetchHTTP(); err != nil {
    return err
  }
  if err := d.SaveBase(); err != nil {
    return err
  }
  fmt.Println("Fetch XML and save - done!")
  return nil
}

func runApp(d Data) error {
  if b := d.ExistFile(); b == false {
    if err := runAppFetch(d); err != nil {
      return err
    }
  }
  if err := d.FetchFile(); err != nil {
    fmt.Printf("Fetch %v\n", err)
    return err
  }
  for {
    for i := 0; i < len(d.Verbs.Verbs); i++ {
      fmt.Printf("Translation: %s\n", d.Verbs.Verbs[i].Translation)
      checkVerbs("Infinitive: ", d.Verbs.Verbs[i].Infinitive)
      checkVerbs("Past Simple: ", d.Verbs.Verbs[i].PastSimple)
      checkVerbs("Past Participle: ", d.Verbs.Verbs[i].PastParticiple)
    }
  }
  return nil
}

func compareStrings(s1 string, s2 string) bool {
  for _, e := range strings.Split(s2, ",") {
    if strings.TrimSpace(e) == s1 {
      return true
    }
  }
  return false
}

func checkVerbs(s string, verb string) {
  resp := ""
  fmt.Printf("%s", s)
  fmt.Scanf("%s", &resp)
  if compareStrings(resp, verb) {
    fmt.Printf("\033[32mCorrect:\033[0m %s\n", verb)
  } else {
    fmt.Printf("\033[31mWRONG!\033[0m Use '%s'.\n", verb)
  }
}

func runAppList(d Data) error {
  if b := d.ExistFile(); b == false {
    if err := runAppFetch(d); err != nil {
      return err
    }
  }
  if err := d.FetchFile(); err != nil {
    fmt.Printf("Fetch %v\n", err)
    return err
  }

  fmt.Println("Infinitive\tPast Simple\tPast Participle\tTranslation")
  for i := 0; i < len(d.Verbs.Verbs); i++ {
    fmt.Printf(
      "%s\t%s\t%s\t%s\n",
      d.Verbs.Verbs[i].Infinitive,
      d.Verbs.Verbs[i].PastSimple,
      d.Verbs.Verbs[i].PastParticiple,
      d.Verbs.Verbs[i].Translation,
    )
  }
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
  return xml.Unmarshal(d.XML, &d.Verbs)
}

func (d *Data) SaveBase() error {
  saveXML, err := xml.Marshal(&d.Verbs)
  if err != nil {
    return err
  }
  if err = ioutil.WriteFile(d.FileName, saveXML, 0644); err != nil {
    return err
  }
  return nil
}

func main() {

  data := Data{
    "http://spacepilot.ru/irregular-verbs/verbs.xml",
    "./verbs.xml",
    nil,
    Verbs{},
  }

  if len(os.Args) < 2 {
    printHelp()
    os.Exit(1)
  }

  switch os.Args[1] {
  case "run":
    if err := runApp(data); err != nil {
      fmt.Printf("Error: %s\n", err)
      os.Exit(1)
    }
  case "fetch":
    if err := runAppFetch(data); err != nil {
      fmt.Printf("Error: %s\n", err)
      os.Exit(1)
    }
  case "list":
    if err := runAppList(data); err != nil {
      fmt.Printf("Error: %s\n", err)
      os.Exit(1)
    }
  default:
    printHelp()
  }
  os.Exit(0)
}
