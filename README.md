# views-counter 
[![views](https://views-counter.dimaglushkov.xyz/github.com/dimaglushkov/views-counter/repository%20views.svg)](#)

## Usage
1. Clone this repo to your server
2. Change urls for which you want to count views in `main.go` (for now it's hardcoded):
  ```go
  var urls = []string{
      "github.com/dimaglushkov/dimaglushkov",
      "github.com/dimaglushkov/views-counter",
      "dimaglushkov.xyz",
  }
  ```
3. build and run docker container with `docker-compose build && docker-compose up`
4. add badge to your website/readme page:
```markdown

[![alt text](https://{your server}/{url to count views}/{prefered label}.svg)(#)](#)
```
```html
<img src="https://{your server}/{url to count views}/{prefered label}.svg" alt="alt text"></img>
```
  
