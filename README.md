# systeminfo

### get process id or process status by process name 
processinfo.go 

demo:
```
package main
import "fmt"
import "github.com/lixiuhu/systeminfo"
func main() {
	fmt.Println(systeminfo.GetProcessPid("notepad.exe"))
}
```
