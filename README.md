# gologrotate
logrotate for golang

# Example

## Code

```
package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "os/signal"
        "syscall"
        "time"
        "github.com/tkuchiki/gologrotate"
)

func main() {
        var writer *bufio.Writer
        var err error
        var f *os.File

        logpath := "/tmp/log"
        f, err = os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE, 0644)
        writer = bufio.NewWriter(f)

        sigChan := make(chan os.Signal, 1)
        signal.Notify(sigChan,
                syscall.SIGUSR1)
        go func() {
                for {
                        s := <-sigChan
                        switch s {
                        case syscall.SIGUSR1:
                                log.Println("Received SIGUSR1")
                                writer, err = gologrotate.Rotate(writer, f, logpath, "", os.O_WRONLY|os.O_CREATE, 0644)
                        }
                }
        }()

        i := 0
        for {
                _, _ = writer.Write([]byte(fmt.Sprintf("%d\n", i)))
                writer.Flush()
                i++
                time.Sleep(10 * time.Millisecond)
        }
}
```

## Result

```
# Current date 2015-10-02
$ go build -o sample.go
$ ./sample &
$ kill -USR1 PID
$ ls /tmp/log*
/tmp/log  /tmp/log-20151002

$ tail -n 5 /tmp/log-20151002
713
714
715
716
717

$ head -n 5 /tmp/log
718
719
720
721
722
```