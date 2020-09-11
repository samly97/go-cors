# Welcome to Go-Cors!

Go-Cors is middleware to apply CORS headers to CORS [requests](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS).

When I split one of my projects from a Golang SSR application to a ReactJS + Golang API SPA pattern, I noticed that when I was testing my APIs locally, that CORS started popping out in the console as the React frontend and Golang backend were hosted on different ports.

I ended up using the existing [rs/cors](https://github.com/rs/cors) package which solved my program. But, my curiosity got the best of me, and I wanted to figure out how to apply CORS headers myself, creating my own middleware.