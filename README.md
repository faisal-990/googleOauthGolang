
## Installation
First get the google's credentials , and save it in a .env file in the root dir of your project 
Install my-project with npm

```bash
go mod tidy
go run .
```
    
## API Reference

#### in the Browser ,go to
http://localhost:8080/google_login

```http
  GET /google_login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Google Credentials|




## Appendix

Any additional information goes here


## Usage/Examples

```go
func main() {
	// load the google oauth2 configs
	router := gin.Default()
	GoogleConfig()

	router.GET("/google_login", HandleGoogleLogin)
	router.GET("/google_callback", HandleGoogleCallback)

	log.Fatal(router.Run(":8080"))
}
```


## Contributing

Contributions are always welcome!



## License

[MIT](https://choosealicense.com/licenses/mit/)

