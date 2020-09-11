# Welcome to Go-Cors!

Go-Cors is middleware to apply CORS headers to CORS [requests](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS).

When I split one of my projects from a Golang SSR application to a ReactJS + Golang API SPA pattern, I noticed that when I was testing my APIs locally, that CORS started popping out in the console as the React frontend and Golang backend were hosted on different ports.

I ended up using the existing [rs/cors](https://github.com/rs/cors) package which solved my program. But, my curiosity got the best of me, and I wanted to figure out how to apply CORS headers myself, creating my own middleware.

## Table of Contents
- [Use Case](#use-case)
  * [Sub-heading](#sub-heading)
    + [Sub-sub-heading](#sub-sub-heading)
- [Installation and Use](#installation-and-use)
  * [Sub-heading](#sub-heading-1)
    + [Sub-sub-heading](#sub-sub-heading-1)

## [Use Case](#use-case)

As mentioned earlier, this use case comes up when we're trying to communicate off of different sites. A different site as defined by CORS includes the protocol, domain, and port.

For my use case, I am testing apis using the `fetch` api from the ReactJS frontend to Golang backend. Here's the JavaScript for my use case:

![Fetch from APIs](img/fetch-commands.png)

The ReactJS application is hosted on port `3000` and the Golang server on `8080`, so by CORS definition they're on different sites. Without whitelisting, I got this:

![Fetch blocked by CORS](img/cors-console-err.png)

Now, whitelisting it...

![Fetch successful](img/fetch-successful.png)

```go
handler := mux.Router()
```

## [Installation and Use](#installation-and-use)