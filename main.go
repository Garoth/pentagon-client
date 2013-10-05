package main

import (
    "log"
    "flag"
    "encoding/json"

    "code.google.com/p/go.net/websocket"
    "github.com/Garoth/go-signalhandlers"

    "github.com/Garoth/pentagon-model"
)

var (
    SERVER_ADDR = flag.String("address",
        "ws://localhost:9217/websocket",
        "server address")
    ORIGIN_ADDR = flag.String("origin-address",
        "http://localhost",
        "my address")
)

func main() {
    log.SetFlags(log.Ltime)
    flag.Parse()

    go signalhandlers.Interrupt()
    go signalhandlers.Quit()

    SendMessage()
}

func SendMessage() {
    ws, err := websocket.Dial(*SERVER_ADDR, "", *ORIGIN_ADDR)
    if err != nil {
        log.Fatalln("Couldn't dial server:", err)
    }

    componentInfo := &pentagonmodel.ClientHeader{}
    componentInfo.Component = pentagonmodel.COMPONENT_EMAIL
    componentInfo.Subcomponent = pentagonmodel.SUBCOMPONENT_EMAIL_MAIN
    bytes, err := json.Marshal(componentInfo)
    if err != nil {
        log.Fatalln("Error encoding component info:", err)
    }

    websocket.Message.Send(ws, string(bytes))

    emailMessage := &pentagonmodel.MailComponentMessage{}
    emailMessage.To = "garoth@gmail.com"
    emailMessage.From = "garoth@gmail.com"
    emailMessage.Subject = "Pentagon Test"
    emailMessage.Message = "Hello!\nTest"

    bytes, err = json.Marshal(emailMessage)
    if err != nil {
        log.Fatalln("Error encoding email message:", err)
    }

    websocket.Message.Send(ws, string(bytes))
}
