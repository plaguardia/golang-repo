Go Routines are not "functions" but they are the way that functions are called.

To call a function as a go routine simply add "go" before the name of the function

ex)
    someFunction() //normal call of function
    go someFunction() //calling function as go routine

Calling something as a go routine allows that process to run in the background and not hold up the cpu.
go routines running in the background will not keep your app alive
Go routines a tied to thier parent functions, so they will stop as soon as thier parent process stops

so below code would result in one print line of "dont wait up", none of the infinit "hi"s showing up and then the program ending

```
func timer(){
    //infinit loop
    for {
        time.sleep(time.seconds * 1)
        fmt.println("hi")
    }
}

func main() {

    go timer()
    fmt.println("Dont wait up!")
}
```

Mutex
this are used when you are trying to prevent multiple functions/processes from accessing a shared element at the same time
Makes the tasks "mutually exclusive" from one another
https://www.youtube.com/watch?v=cjMdUmfzQWs&ab_channel=TutorialEdge
