package minecraft_server_installer

import (
	"fmt"
	"log"
)

//path: Dir where server should be installed
//rev: Version of paper to install
func InstallPaper(path string, rev string) {
	fmt.Println("Downloading Paper jar...")
	downloadPaper(path, rev)
	fmt.Println("Running Server to prepare eula...")
	RunServer(path+"paper.jar", path)
	fmt.Println("Modifying eula...")
	ModifyEula(path)
}
func downloadPaper(path string, rev string) {
	paper188 := "https://papermc.io/ci/job/Paper/443/artifact/paperclip.jar"
	paper1122 := "https://papermc.io/ci/job/Paper/lastSuccessfulBuild/artifact/paperclip.jar"
	paper1144 := "https://papermc.io/ci/job/Paper-1.14/lastSuccessfulBuild/artifact/paperclip.jar"
	paper1152 := "https://papermc.io/ci/job/Paper-1.15/lastSuccessfulBuild/artifact/paperclip.jar"
	var revUrl string
	if rev == "1.8" || rev == "1.8.8" || rev == "1.8.9" {
		revUrl = paper188
	} else if rev == "1.12" || rev == "1.12.2" {
		revUrl = paper1122
	} else if rev == "1.14" || rev == "1.14.4" {
		revUrl = paper1144
	} else if rev == "1.15" || rev == "1.15.2" {
		revUrl = paper1152
	} else {
		log.Fatal("Unfortunately, we probably don't support the version of paper you want to install \n" +
			"The builds for versions outside of 1.15.2, 1.14.4, 1.12.2, and 1.8.8/1.8.9 are surprisingly hard to find\n" +
			"If you really need to install another version, consider manually installing or using spigot builds")
	}
	err := DownloadFile(path+"paper.jar", revUrl)
	if err != nil {
		log.Fatal("Error: Failed to download paper \n \n", err)
	}

}
