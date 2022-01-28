## How to run this app?
- From the root of the project directory run `go run .` and the app will work on `localhost:8004`
- If you want to adjust the `PORT` and the `DOMAIN`, you can do that [from here](config/config.go)

## API
### Encode url
- endpoint: http://localhost:8000/encode
- Method: POST
- Payload:
```
{
	"url": "https://www.google.com/search?q=anything"
}
```
---
### decode url
- endpoint: http://localhost:8000/decode/B
- Method: GET
- Payload:
```
{
  "id": 1,
  "shortUrl": "localhost::8000/B",
  "original": "https://www.google.com/search?q=anything"
}
```
