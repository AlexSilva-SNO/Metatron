package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var logFile *os.File

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || info.Size() == 0 {
		return false
	}
	return true
}

func waitForFile(path string, etapa string, retries int) {
	for i := 0; i < retries; i++ {
		if fileExists(path) {
			return
		}
		time.Sleep(1 * time.Second)
	}
	logMessage(fmt.Sprintf("[ERRO] Falha na etapa %s: arquivo %s não encontrado ou vazio.", etapa, path))
	os.Exit(1)
}

func countLines(path string) int {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("wc -l < %s", path))
	out, err := cmd.Output()
	if err != nil {
		return 0
	}
	line := strings.TrimSpace(string(out))
	var count int
	fmt.Sscanf(line, "%d", &count)
	return count
}

func runCommand(stage, command string) {
	msg := fmt.Sprintf("[%s] Executando: %s", stage, command)
	fmt.Println(msg)
	logMessage(msg)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		logMessage(fmt.Sprintf("[ERRO] Falha ao executar %s: %v", stage, err))
		log.Fatalf("[ERRO] Falha ao executar %s: %v", stage, err)
	}
}

func logMessage(msg string) {
	if logFile != nil {
		logFile.WriteString(fmt.Sprintf("%s\n", msg))
	}
}
func printBanner() {
	fmt.Println(" ███╗   ███╗███████╗████████╗ █████╗ ████████╗██████╗  ██████╗ ███╗   ██╗")
	fmt.Println(" ████╗ ████║██╔════╝╚══██╔══╝██╔══██╗╚══██╔══╝██╔══██╗██╔═══██╗████╗  ██║")
	fmt.Println(" ██╔████╔██║█████╗     ██║   ███████║   ██║   ██████╔╝██║   ██║██╔██╗ ██║")
	fmt.Println(" ██║╚██╔╝██║██╔══╝     ██║   ██╔══██║   ██║   ██╔══██╗██║   ██║██║╚██╗██║")
	fmt.Println(" ██║ ╚═╝ ██║███████╗   ██║   ██║  ██║   ██║   ██║  ██║╚██████╔╝██║ ╚████║")
	fmt.Println(" ╚═╝     ╚═╝╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝")
	fmt.Println("")
	fmt.Println("                         Metatron v1.0 - Desenvolvido por Alex Silva")
	fmt.Println("===============================================================================================")
}

func main() {
	if len(os.Args) < 3 || os.Args[1] != "--domain" {
		log.Fatalf("[ERRO] Uso: %s --domain <domínio>", os.Args[0])
	}
  printBanner()
	domain := os.Args[2]
	outputDir := filepath.Join("results", domain)
	os.MkdirAll(outputDir, 0755)

	logPath := filepath.Join(outputDir, "log.txt")
	var err error
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("[ERRO] Falha ao abrir log.txt: %v", err)
	}
	defer logFile.Close()

	subsPath := filepath.Join(outputDir, "subs.txt")
	runCommand("1", fmt.Sprintf("subfinder -d %s -silent | anew %s", domain, subsPath))
	waitForFile(subsPath, "1", 20)
	totalSubs := countLines(subsPath)
	if totalSubs == 0 {
		logMessage("[!] Nenhum subdomínio encontrado para rodar httpx.")
		log.Fatalf("[!] Nenhum subdomínio encontrado para rodar httpx.")
	}
	fmt.Printf("[!] Subdomínios encontrados: %d\n", totalSubs)
	logMessage(fmt.Sprintf("[!] Subdomínios encontrados: %d", totalSubs))

	httpPath := filepath.Join(outputDir, "http200.txt")
	runCommand("2", fmt.Sprintf("httpx -list %s -silent | anew %s || touch %s", subsPath, httpPath, httpPath))
	waitForFile(httpPath, "2", 20)
	if countLines(httpPath) == 0 {
		logMessage("[ERRO] httpx executado, mas nenhum host retornou status HTTP válido.")
		log.Fatalf("[ERRO] httpx executado, mas nenhum host retornou status HTTP válido.")
	}

	paramsPath := filepath.Join(outputDir, "params.txt")
	runCommand("3", fmt.Sprintf("katana -list %s -f qurl -silent | anew %s || touch %s", httpPath, paramsPath, paramsPath))
	waitForFile(paramsPath, "3", 10)

	urlsPath := filepath.Join(outputDir, "urls.txt")
	runCommand("4", fmt.Sprintf("cat %s | urlfinder -silent | anew %s || touch %s", subsPath, urlsPath, urlsPath))
	waitForFile(urlsPath, "4", 10)

	urls2Path := filepath.Join(outputDir, "urls2.txt")
	runCommand("5", fmt.Sprintf("cat %s %s | anew %s || touch %s", paramsPath, urlsPath, urls2Path, urls2Path))
	waitForFile(urls2Path, "5", 10)

	xssPath := filepath.Join(outputDir, "xss.txt")
	runCommand("6", fmt.Sprintf("cat %s | dalfox pipe --skip-bav --skip-grepping | anew %s || touch %s", urls2Path, xssPath, xssPath))
	waitForFile(xssPath, "6", 10)

fmt.Printf("✅ [Metatron] Resultados salvos em: %s\n", xssPath)
fmt.Printf("✅ [Metatron] Pipeline concluído com sucesso. Todos os resultados estão em: %s\n", outputDir)
logMessage(fmt.Sprintf("✅ [Metatron] Resultados salvos em: %s", xssPath))
logMessage(fmt.Sprintf("✅ [Metatron] Pipeline concluído com sucesso. Todos os resultados estão em: %s", outputDir))

}
