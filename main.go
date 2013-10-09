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
    WEBSOCKET *websocket.Conn
)

func main() {
    log.SetFlags(log.Ltime)
    flag.Parse()

    go signalhandlers.Interrupt()
    go signalhandlers.Quit()

    var err error
    WEBSOCKET, err = websocket.Dial(*SERVER_ADDR, "", *ORIGIN_ADDR)
    if err != nil {
        log.Fatalln("Couldn't dial server:", err)
    }

    TryKV()
}

func SendMessage() {
    componentInfo := &pentagonmodel.ClientHeader{}
    componentInfo.Component = pentagonmodel.COMPONENT_EMAIL
    componentInfo.Subcomponent = pentagonmodel.SUBCOMPONENT_EMAIL_MAIN
    bytes, err := json.Marshal(componentInfo)
    if err != nil {
        log.Fatalln("Error encoding component info:", err)
    }

    websocket.Message.Send(WEBSOCKET, string(bytes))

    emailMessage := &pentagonmodel.MailComponentMessage{}
    emailMessage.To = "garoth@gmail.com"
    emailMessage.From = "garoth@gmail.com"
    emailMessage.Subject = "Pentagon Test"
    emailMessage.Message = "Hello!\nTest"

    bytes, err = json.Marshal(emailMessage)
    if err != nil {
        log.Fatalln("Error encoding email message:", err)
    }

    websocket.Message.Send(WEBSOCKET, string(bytes))
}

func TryKV() {
    header := &pentagonmodel.ClientHeader{}
    header.Component = pentagonmodel.COMPONENT_KV
    header.Subcomponent = pentagonmodel.SUBCOMPONENT_KV_WRITE
    bytes, _ := json.Marshal(header)
    websocket.Message.Send(WEBSOCKET, string(bytes))

    command := &pentagonmodel.KeyValueWriteMessage{}
    command.Category = "client-test"
    command.Key = "herro"
    command.Value = "wurld"
    bytes, _ = json.Marshal(command)
    websocket.Message.Send(WEBSOCKET, string(bytes))

    header2 := &pentagonmodel.ClientHeader{}
    header2.Component = pentagonmodel.COMPONENT_KV
    header2.Subcomponent = pentagonmodel.SUBCOMPONENT_KV_READ
    bytes, _ = json.Marshal(header2)
    websocket.Message.Send(WEBSOCKET, string(bytes))

    command2 := &pentagonmodel.KeyValueReadMessage{}
    command2.Category = "client-test"
    command2.Key = "herro"
    bytes, _ = json.Marshal(command2)
    websocket.Message.Send(WEBSOCKET, string(bytes))
}
