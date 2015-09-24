This code was adapted from chatper 2 Go Programming Blueprints Book by Mat Ryer. The code allows you to type in the base URL (which right now is on localhost:8080') + /chat/ + any string and it will open a new chat room for that topic if one doesn't already exist' 

# Go Programming Blueprints

![Go Blueprints by Mat Ryer book cover](https://raw.githubusercontent.com/matryer/goblueprints/master/artwork/bookcover.jpg)

### Chapter 2

  * Browse the [Source code](https://github.com/matryer/goblueprints/tree/master/chapter2)

Notes:

  * Page 53: `w.Header.Set` should be `w.Header().Set` since `Header` is a function on `http.ResponseWriter`.


