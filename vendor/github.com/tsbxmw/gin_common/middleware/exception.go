package middleware

import (
    "bytes"
    "fmt"
    "github.com/gin-gonic/gin"
    common "github.com/tsbxmw/gin_common"
    "io/ioutil"
    "net"
    "net/http/httputil"
    "os"
    "runtime"
    "strings"
    "time"
)

var (
    dunno     = []byte("???")
    centerDot = []byte("·")
    dot       = []byte(".")
    slash     = []byte("/")
)

func ExceptionInit(e *gin.Engine) {
    exceptionMiddleware := ExceptionMiddleware()
    e.Use(exceptionMiddleware)
}

func ExceptionMiddleware() (gin.HandlerFunc) {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                var brokenPipe bool
                if ne, ok := err.(*net.OpError); ok {
                    if se, ok := ne.Err.(*os.SyscallError); ok {
                        if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
                            brokenPipe = true
                        }
                    }
                }
                stack := stack(3)
                httpRequest, _ := httputil.DumpRequest(c.Request, false)
                headers := strings.Split(string(httpRequest), "\r\n")
                for idx, header := range headers {
                    current := strings.Split(header, ":")
                    if current[0] == "Authorization" {
                        headers[idx] = current[0] + ": *"
                    }
                }
                if brokenPipe {
                    common.LogrusLogger.Errorf("%s\n%s", err, string(httpRequest))
                } else if gin.IsDebugging() {
                    common.LogrusLogger.Errorf("[Recovery] %s panic recovered:\n%s\n%s\n%s",
                        timeFormat(time.Now()), strings.Join(headers, "\r\n"), err, stack)
                    fmt.Printf("[Recovery] %s panic recovered:\n%s\n%s\n%s",
                        timeFormat(time.Now()), strings.Join(headers, "\r\n"), err, stack)
                } else {
                    common.LogrusLogger.Error(err)
                }

                extension := make(map[string]interface{}, 0)
                extension["code"] = c.Keys["code"]
                extension["message"] = err
                TracerHandler("exception", "server", c, false, extension)
                
                c.AbortWithStatusJSON(common.HTTP_STATUS_OK, gin.H{
                    "message": err.(error).Error(),
                    "data":    []string{},
                    "code":    c.Keys["code"],
                })
            }
        }()
        c.Next()
    }
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func stack(skip int) []byte {
    buf := new(bytes.Buffer) // the returned data
    // As we loop, we open files and read them. These variables record the currently
    // loaded file.
    var lines [][]byte
    var lastFile string
    for i := skip; ; i++ { // Skip the expected number of frames
        pc, file, line, ok := runtime.Caller(i)
        if !ok {
            break
        }
        // Print this much at least.  If we can't find the source, it won't show.
        fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
        if file != lastFile {
            data, err := ioutil.ReadFile(file)
            if err != nil {
                continue
            }
            lines = bytes.Split(data, []byte{'\n'})
            lastFile = file
        }
        fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
    }
    return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
    n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
    if n < 0 || n >= len(lines) {
        return dunno
    }
    return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
    fn := runtime.FuncForPC(pc)
    if fn == nil {
        return dunno
    }
    name := []byte(fn.Name())
    // The name includes the path name to the package, which is unnecessary
    // since the file name is already included.  Plus, it has center dots.
    // That is, we see
    //	runtime/debug.*T·ptrmethod
    // and want
    //	*T.ptrmethod
    // Also the package path might contains dot (e.g. code.google.com/...),
    // so first eliminate the path prefix
    if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
        name = name[lastSlash+1:]
    }
    if period := bytes.Index(name, dot); period >= 0 {
        name = name[period+1:]
    }
    name = bytes.Replace(name, centerDot, dot, -1)
    return name
}

func timeFormat(t time.Time) string {
    var timeString = t.Format("2006/01/02 - 15:04:05")
    return timeString
}