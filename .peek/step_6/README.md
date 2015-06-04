#Solutiuons for exercises

**All make it in time**
```
mgrol@local:~/src/go/src/github.com/MarcGrol/go-training/.peek/step_6 (master)$ ./client  -concurrent=10 -delay=100
2015/06/04 12:04:13 Using baseURL: http://localhost:3000, delay: 100, deadline: 1000, concurrent: 10
2015/06/04 12:04:13 Client completed successfully in 5
2015/06/04 12:04:13 Got finished client: 1
2015/06/04 12:04:13 Client completed successfully in 33
2015/06/04 12:04:13 Got finished client: 2
2015/06/04 12:04:13 Client completed successfully in 34
2015/06/04 12:04:13 Got finished client: 3
2015/06/04 12:04:13 Client completed successfully in 36
2015/06/04 12:04:13 Got finished client: 4
2015/06/04 12:04:13 Client completed successfully in 62
2015/06/04 12:04:13 Got finished client: 5
2015/06/04 12:04:13 Client completed successfully in 76
2015/06/04 12:04:13 Got finished client: 6
2015/06/04 12:04:13 Client completed successfully in 85
2015/06/04 12:04:13 Got finished client: 7
2015/06/04 12:04:13 Client completed successfully in 92
2015/06/04 12:04:13 Got finished client: 8
2015/06/04 12:04:13 Client completed successfully in 99
2015/06/04 12:04:13 Got finished client: 9
2015/06/04 12:04:13 Client completed successfully in 119
2015/06/04 12:04:13 Got finished client: 10
2015/06/04 12:04:13 Done
10 of 10 clients returned within deadline 1000
average response-time: 64
```

**Half makes it in time**

```
mgrol@local:~/src/go/src/github.com/MarcGrol/go-training/.peek/step_6 (master)$ ./client  -concurrent=10 -delay=1000
2015/06/04 12:04:21 Using baseURL: http://localhost:3000, delay: 1000, deadline: 1000, concurrent: 10
2015/06/04 12:04:21 Client completed successfully in 284
2015/06/04 12:04:21 Got finished client: 1
2015/06/04 12:04:22 Client completed successfully in 626
2015/06/04 12:04:22 Got finished client: 2
2015/06/04 12:04:22 Client completed successfully in 707
2015/06/04 12:04:22 Got finished client: 3
2015/06/04 12:04:22 Client completed successfully in 845
2015/06/04 12:04:22 Got finished client: 4
2015/06/04 12:04:22 Timeout
2015/06/04 12:04:22 Client timout fetching with url http://localhost:3000/step6?delay=1000&loop=1
2015/06/04 12:04:22 Client timout fetching with url http://localhost:3000/step6?delay=1000&loop=2
2015/06/04 12:04:22 Client timout fetching with url http://localhost:3000/step6?delay=1000&loop=4
2015/06/04 12:04:22 Client timout fetching with url http://localhost:3000/step6?delay=1000&loop=6
2015/06/04 12:04:22 Client timout fetching with url http://localhost:3000/step6?delay=1000&loop=8
2015/06/04 12:04:22 Client timout fetching with url http://localhost:3000/step6?delay=1000&loop=9
4 of 10 clients returned within deadline 1000
average response-time: 615
```

**Almost no one makes it in time**
```
mgrol@local:~/src/go/src/github.com/MarcGrol/go-training/.peek/step_6 (master)$ ./client  -concurrent=10
-delay=2000
2015/06/04 12:04:32 Using baseURL: http://localhost:3000, delay: 2000, deadline: 1000, concurrent: 10
2015/06/04 12:04:32 Client completed successfully in 64
2015/06/04 12:04:32 Got finished client: 1
2015/06/04 12:04:33 Timeout
2015/06/04 12:04:33 Client timout fetching with url http://localhost:3000/step6?delay=2000&loop=2
2015/06/04 12:04:33 Client timout fetching with url http://localhost:3000/step6?delay=2000&loop=4
2015/06/04 12:04:33 Client timout fetching with url http://localhost:3000/step6?delay=2000&loop=5
2015/06/04 12:04:33 Client timout fetching with url http://localhost:3000/step6?delay=2000&loop=6
2015/06/04 12:04:33 Client timout fetching with url http://localhost:3000/step6?delay=2000&loop=7
2015/06/04 12:04:33 Client timout fetching with url http://localhost:3000/step6?delay=2000&loop=8
2015/06/04 12:04:33 Client timout fetching with url http://localhost:3000/step6?delay=2000&loop=9
2015/06/04 12:04:33 Client timout fetching with url http://localhost:3000/step6?delay=2000&loop=0
2015/06/04 12:04:33 Client timout fetching with url http://localhost:3000/step6?delay=2000&loop=1
1 of 10 clients returned within deadline 1000
average response-time: 64
```

