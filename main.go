package main

import (
	"fmt"
  "strings"
  "os"
)

const (
 // 应该放在 config 之类的包里
 DaoCloudCIName = "DaoCloud"
)

func isLowerAlnum(c rune) bool {
 return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z')
}

func isSpecialChar(c rune) bool {
 return c == '-' || c == '_' || c == '.'
}

// tag must match "[a-z0-9]+(?:[._-][a-z0-9]+)*"
func FixBuildTag(tag string) string {
 lastRune := ' '
 tmp := ""

 tag = strings.ToLower(tag)
 for _, c := range tag {
  if isLowerAlnum(c) {
   tmp += string(c)
  } else if isSpecialChar(c) && isLowerAlnum(lastRune) {
   tmp += string(c)
  } else if isLowerAlnum(lastRune) {
   tmp += "-"
  }
  lastRune = c
 }

 if len(tmp) < 1 {
  tmp = DaoCloudCIName
 }

 if isSpecialChar(rune(tmp[len(tmp)-1])) {
  tmp = tmp[:len(tmp)-1]
 }

 return tmp
}


func main() {
  var imageTag string
  var branch = os.Getenv("DAO_COMMIT_BRANCH")
  var tag = os.Getenv("DAO_COMMIT_TAG")
  var hash = os.Getenv("DAO_COMMIT_SHA")

  if len(tag)>0 {
    imageTag = tag
  } else {
    imageTag = branch + "-" + hash[:7]
  }
  imageTag = FixBuildTag(imageTag)
  fmt.Println(imageTag)
}
