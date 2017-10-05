package main

import(
	"os/exec"
	"bytes"
	"log"
	"strings"
)

func getDoc(text string) string {
	mass := strings.Fields(text)
	if len(mass) == 3 && mass[0] == "godoc" {
	cmd := exec.Command("godoc", mass[1], mass[2])
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
	return out.String()
}
 return "Неверный запрос"
} 
