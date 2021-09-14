package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func getIoGit() (string, string, string) {
	log.SetFlags(log.Flags() | log.Lshortfile)
	args := strings.Fields("git remote -v")
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(string(out))
	retArr := strings.Fields(string(out))
	// log.Println(retArr)
	// log.Println(len(retArr))
	// 获得git项目地址
	// 制取到对应的io发布的地址 比如:git@github.com:hfbhfb/abc.git -> git@github.com:hfbhfb/abc-io.git
	for _, v := range retArr {
		// log.Println(v)
		if strings.Contains(v, "git@github.com") {
			// log.Println(v)
			arr1 := strings.Split(v, "/")
			// log.Println(arr1[1])
			arr2 := strings.Split(arr1[1], ".")
			// log.Println(arr2[0])
			currentRemote := arr1[0] + "/" + arr2[0] + "." + arr2[1]
			ioGitRemote := arr1[0] + "/" + arr2[0] + "-io." + arr2[1]
			log.Println(ioGitRemote)
			return currentRemote, ioGitRemote, arr2[0]
		}
	}
	return "", "", ""
}

func prePareIOdirAll(currtGit, ioGitInfo, baseName string) error {
	log.Println(currtGit)
	log.Println(ioGitInfo)
	if ioGitInfo == "" {
		return errors.New("no git info")
	}
	args := strings.Fields("mkdir -pv  docs/.vuepress/dist-io/")
	_, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir("docs/.vuepress/dist-io/")
	{

		pwdGit, _, _ := getIoGit()
		if pwdGit == currtGit {
			args := strings.Fields("git init")
			_, err := exec.Command(args[0], args[1:]...).Output()
			if err != nil {
				log.Fatal(err)
			}

			{
				args := strings.Fields("git remote add origin " + ioGitInfo)
				_, err := exec.Command(args[0], args[1:]...).Output()
				if err != nil {
					log.Fatal(err)
				}
			}
			{
				args := strings.Fields("git pull origin master")
				_, err := exec.Command(args[0], args[1:]...).Output()
				if err != nil {
					// log.Fatal(err)
				}
			}
		}

		pwdGitSecond, _, _ := getIoGit()
		if pwdGitSecond == currtGit {
			log.Fatal("发布的节点错误")
		} else {
			// {
			// 	args := strings.Fields("pwd")
			// 	log.Println(args)
			// 	out, err := exec.Command(args[0], args[1:]...).Output()
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// 	log.Println(string(out))
			// }

			{
				// args := strings.Fields(`/bin/sh -c 'cp -rf ../dist/* .'`)
				// log.Println(args)
				args := []string{"/bin/sh", "-c", "cp -rf ../dist/* ."}
				_, err := exec.Command(args[0], args[1:]...).Output()
				if err != nil {
					log.Fatal(err)
				}

			}
			{
				args := strings.Fields("git add .")
				_, err := exec.Command(args[0], args[1:]...).Output()
				if err != nil {
					log.Fatal(err)
				}
			}
			{
				// args := strings.Fields(`git commit -a -m "auto commit"`)
				args := []string{"/bin/sh", "-c", `git commit -a -m "auto commit"`}
				_, err := exec.Command(args[0], args[1:]...).Output()
				if err != nil {
					// 没内容时,退出变量为1
					// log.Fatal(err)
				}
			}

			if gPush {
				// args := strings.Fields(`git commit -a -m "auto commit"`)
				args := []string{"/bin/bash", "-c", `git push origin master`}
				// cmd := exec.Command("ls")
				_, err := exec.Command(args[0], args[1:]...).Output()
				if err != nil {
					log.Fatal(err)
				}
			}

		}

	}

	return nil
}

func prePareBase(currtGit, ioGitInfo, baseName string) error {

	fileName := "docs/.vuepress/config.js"
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "base:") {
			if gCname == "" {
				lines[i] = `  base: '/` + baseName + `-io/',`
			} else {
				lines[i] = `  base: '/` + baseName + `-io/',`
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	// log.Println(baseName)
	// log.Println(lines)
	// log.Println(baseName)
	return nil
}

func buildProject() {
	if gBuild {
		// args := strings.Fields(`/bin/sh -c 'cp -rf ../dist/* .'`)
		// log.Println(args)
		args := []string{"/bin/sh", "-c", "yarn run build"}
		_, err := exec.Command(args[0], args[1:]...).Output()
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

var gCname = ""
var gBuild = true
var gPush = true

func main() {
	// 默认没有cName值,直接使用github提供的pages服务
	var cName = flag.String("cname", "", "Input Your Name")
	var cBuild = flag.Bool("build", true, "Input Your Name")
	var cPush = flag.Bool("push", true, "Input Your Name")
	flag.Parse()
	gCname = *cName
	gBuild = *cBuild
	gPush = *cPush

	// log.Println(gCname)
	// log.Println(gBuild)
	// log.Println(gPush)
	// return

	// log.Println("hello world")
	prePareBase(getIoGit())
	buildProject()
	prePareIOdirAll(getIoGit())

}
