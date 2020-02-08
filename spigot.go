package minecraft_server_installer

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func InstallSpigot(path string, rev string){
	fmt.Println("Downloading installer (BuildTools) to ", path, "...")
	downloadBuildTools(path, rev)
	fmt.Println("Installing server via BuildTools...")
	time.Sleep(time.Millisecond * 100)
	installBuildTools(path, rev)
	fmt.Println("Running Server to prepare eula...")
	RunServer(path + "spigot-" + rev + ".jar")
	fmt.Println("Modifying eula...")
	ModifyEula(path)


}
func ModifyEula(path string){
	eula, err := ioutil.ReadFile(path + "eula.txt")
	if err != nil{
		log.Fatal("Error: Failed to read eula \n \n", err)
	}
	eulaString := string(eula)
	updatedEula := strings.Replace(eulaString, "false", strconv.FormatBool(getEula()), -1)
	err = ioutil.WriteFile(path+"eula.txt", []byte(updatedEula), 0400)
	if err != nil{
		log.Fatal("Error: Failed to write to eula \n \n", err)
	}
}
func RunServer(filepath string){
	//TODO: add error handling
	cmd := exec.Command("java", "-jar", getJavaArgs(), filepath)
	outPipe, _ := cmd.StdoutPipe()
	go fmt.Print(outPipe)
	inPipe, _ := cmd.StdinPipe()
	_ = cmd.Start()
	time.Sleep(time.Second * 3)
	_, _ = inPipe.Write([]byte("stop"))
	_ = cmd.Wait()

}
func installBuildTools(path string, rev string){
	cmd := exec.Command("java", "-jar", getJavaArgs(), path + "BuildTools.jar", "--rev", rev)
	cmd.Dir = path
	err := cmd.Start()
	if err != nil{
		log.Fatal("Error: Failed to run spigot installer (BuildTools) \n \n", err)
	}
	outPipe, errPipe := cmd.StdoutPipe()
	/*if err != nil{
		log.Fatal("Error: Failed to show spigot installer (BuildTools) output \n \n", err)
	}*/
	go fmt.Print(outPipe)
	go fmt.Print(errPipe)
	err = cmd.Wait()
	if err != nil{
		log.Fatal("Error: Something weird happened when waiting for installer (BuildTools) to complete \n \n", err)
	}
}
func downloadBuildTools(path string, rev string){
	jarUrl := "https://hub.spigotmc.org/jenkins/job/BuildTools/lastSuccessfulBuild/artifact/target/BuildTools.jar"
	err := DownloadFile(path + "spigot-" + rev + ".jar", jarUrl)
	if err != nil{
		log.Fatal("Error: Failed to download spigot (BuildTools) \n \n", err)
	}
}
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

