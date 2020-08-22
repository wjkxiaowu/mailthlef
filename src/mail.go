package main

import (
   "os/exec"
   "net/smtp"
   "strings"
   "fmt"
   "flag"
)

var cmdStr string

func init() {
    flag.StringVar(&cmdStr,"cmdStr","ls#-lah","command")
}

func SendToMail(user, pwd, host, to, subject, body, mailtype string) error{
   hp := strings.Split(host, ":")
   auth := smtp.PlainAuth("", user, pwd, hp[0])
   var content_type string
   if mailtype == "html" {
      content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
   } else {
      content_type = "Content-Type: text/plain" + "; charset=UTF-8"
   }

   msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " +subject+ "\r\n" + content_type + "\r\n\r\n" + body)
   send_to := strings.Split(to, ";")
   err := smtp.SendMail(host, auth, user, send_to, msg)
   return err
}


func main(){
    flag.Parse()

    user    := "*********@qq.com"
    pwd     := "****************" //Note: this is your qq mailbox Authorization code, not the password of your qq mailbox !!!
    host    := "smtp.qq.com:25"
    to      := "********@163.com"
    subject := "Present for you !"

    cmdStrs := strings.Split(cmdStr, "#")
    cmd := exec.Command(cmdStrs[0], cmdStrs[1:]...)
    out, _ := cmd.CombinedOutput()
    body := `<html><body><h3>"%s"</h3></body></html>`
    body = fmt.Sprintf(body, string(out))
    fmt.Println(body)

    fmt.Println("SEND EMAIL ...")
    err := SendToMail(user, pwd, host, to, subject, body, "html")
    if err != nil{
       fmt.Printf("SEND MAIL ERROR !\nreason: %s", err)
    }
    fmt.Println("SEND MAIL SUCCESS !")
}
