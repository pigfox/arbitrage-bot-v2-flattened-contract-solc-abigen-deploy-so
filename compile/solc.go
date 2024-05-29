package compile

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/util"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Run(filename string) {
	currentPath := util.GetCurrentPath()
	contractFileName := filename
	contractFile := currentPath + "/sol/" + contractFileName + ".sol" // Path to the Solidity contract
	outputDir := currentPath + "/solc-output"                         // Directory to store ABI and BIN files
	apiDir := currentPath + "/api"                                    // Directory to store generated Go bindings
	generatedFileName := filename
	generatedPackageName := "api"

	err := util.EmptyFolder(outputDir)
	if err != nil {
		log.Fatalf(" %v\n: %s can't empty folder", err, outputDir)
	}

	err = util.EmptyFolder(apiDir)
	if err != nil {
		log.Fatalf(" %v\n: %s can't empty folder", err, apiDir)
	}

	// Ensure output and API directories exist
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf(" %v\n: %s not writable", err, outputDir)
	}
	err = os.MkdirAll(apiDir, os.ModePerm)
	if err != nil {
		log.Fatalf(" %v\n: %s not writable", err, outputDir)
	}

	// Compile the contract using solc
	cmd := exec.Command("solc", "--optimize", "--optimize-runs", "200", "--abi", "--bin", "--overwrite", "--output-dir", outputDir, contractFile) //nolint:lll
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to compile contract: %v\nOutput: %s", err, string(output))
	}
	/*
		err = util.List(outputDir)
		if err != nil {
			log.Fatalf(" %v\n: %s can't list folder", err, outputDir)
		}
	*/
	fmt.Println("Contract compiled successfully")
	//mergedFiles(outputDir, apiDir, generatedFileName, generatedPackageName)
	baseFile(outputDir, apiDir, generatedFileName, generatedPackageName)
}

//both the baseFile and mergedFiles functions are failing
