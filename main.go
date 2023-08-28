package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/en"
	"github.com/olebedev/when/rules/common"
	"github.com/gen2brain/beeep"
	"strings"
)

const (
	markName = "GOLANG-CLI-REMINDER"
	markVal = "1"
)

func main(){
	if len(os.Args) < 3{
		fmt.Printf("Usage: %s <hh:mm> <text-message> \n", os.Args[0])
		os.Exit(1)
	}

	now := time.Now()
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	r, err := w.Parse(os.Args[1], now)
	if err != nil {
		// an error has occurred
		fmt.Printf("An unexpected error has occured! \n")
		os.Exit(2)
	}
	if  r == nil {
		// no matches found
		fmt.Printf("Usage: %s <hh:mm> <text-message> \n", os.Args[0])
		os.Exit(2)
	}
	if now.After(r.Time) || now.Equal(r.Time){
		fmt.Printf("Please set remainder for a future time! \n")
		os.Exit(3)
	}
	diff := r.Time.Sub(now)

	if os.Getenv(markName) == markVal{
		time.Sleep(diff)
		err := beeep.Alert("Reminder", strings.Join(os.Args[2:],"\n"), "assets/information.png")
		if err != nil{
			fmt.Println(err)
			os.Exit(4)
		}
	}else{
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markVal))
		if err := cmd.Start(); err != nil{
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println("Remainder will be displayed after ", diff.Round(time.Second))
		os.Exit(0)
	}
}