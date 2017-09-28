# Go workers thread pool

Implementation of a Concurrent Pattern of processing jobs with a fixed amount of workers, it is based on [Marcio Castilho
 blog article "Handling 1M requests per minute"](https://medium.com/smsjunk/handling-1-million-requests-per-minute-with-golang-f70ac505fcaa). The differences from the article code & this repo:
 
 * extracted the process function (for testing and clarity)
 * removed the JOB_QUEUE global variable
 * added unit test & web server for a better understanding on how it works
 * split the algorithm to different files

I wanted to understand it better so I made a running example, with a few alternations. I also added a visual representation using a HTML basic client.

### Usage

```bash
git clone git@github.com:bgadrian/go-worker-thread-pool.git
cd go-worker-thread-pool/

#make sure you have port 8080 free
#for windows 
go build -x -o server.exe && ./server.exe
#for others
go build -o server && chmod +x ./server && ./server

#open http://localhost:8080 in your browser & keep this process open.
```



#### Copyright
B.G.Adrian 2017

