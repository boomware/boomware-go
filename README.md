# Boomware
Go client for boomware API https://boomware.com/docs/en/

### Send SMS
```go
// Create a default client
client := boomware.New(<#your token#>)

// Make a sms request
request := new(boomware.SMSRequest) 
request.Number = "+18000000000"
request.Text = "hi!"

// Send sms
response := client.SMS(request)

if response.Error != nil {
    log.Println("sending sms error", response.Error)
} else {
    log.Println("successfully sent sms, id:", response.ID)
}
```