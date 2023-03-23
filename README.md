# Golang API for sending emails

This is a simple API written in Golang that accepts a POST request and sends an email with the contents of the fields.
I'm planning to use it @ my website for a simple contact form.

## Usage

To use the API, you'll need to start the server by running the following command in the project directory:

```go
go run main.go
```

There's a dockerfile too, run it as needed, it's just an example.

POST request body

- `from`: The email address of the sender
- `to`: The email address of the recipient
- `subject`: The subject of the email
- `body`: The body of the email


Config file example, put it in a .go file inside `private` folder

```go
type Config struct {
	Server   string
	User     string
	Password string
}

func GetConfig() Config {
	return Config{
		Server:   "smtp.example.com",
		User:     "your-email@example.com",
		Password: "your-email-password",
	}
} 
```
