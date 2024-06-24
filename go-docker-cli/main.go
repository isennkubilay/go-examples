package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interactive Docker CLI")
	fmt.Println("----------------------")

	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Build Docker image")
		fmt.Println("2. Push Docker image")
		fmt.Println("3. Save Docker image as gzip")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		choice, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			buildImage(reader)
		case "2":
			pushImage(reader)
		case "3":
			saveImageAsGzip(reader)
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func buildImage(reader *bufio.Reader) {
	fmt.Print("Enter the Docker build context (e.g., . for current directory): ")
	context, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	context = strings.TrimSpace(context)

	fmt.Print("Enter the Docker image name (e.g., my-image): ")
	imageName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	imageName = strings.TrimSpace(imageName)

	cmd := exec.Command("docker", "build", "-t", imageName, context)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Building Docker image...")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running docker build: %v\n", err)
	} else {
		fmt.Println("Docker image built successfully!")
	}
}

func pushImage(reader *bufio.Reader) {
	fmt.Print("Enter the Docker image name to push (e.g., my-image): ")
	imageName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	imageName = strings.TrimSpace(imageName)

	cmd := exec.Command("docker", "push", imageName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Pushing Docker image...")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running docker push: %v\n", err)
	} else {
		fmt.Println("Docker image pushed successfully!")
	}
}

func saveImageAsGzip(reader *bufio.Reader) {
	fmt.Print("Enter the Docker image name to save (e.g., my-image): ")
	imageName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	imageName = strings.TrimSpace(imageName)

	fmt.Print("Enter the output file name (e.g., my-image.tar.gz): ")
	outputFileName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	outputFileName = strings.TrimSpace(outputFileName)

	cmd := exec.Command("docker", "save", imageName)
	pipeReader, pipeWriter := io.Pipe()
	cmd.Stdout = pipeWriter

	gzipCmd := exec.Command("gzip", "-c")
	gzipCmd.Stdin = pipeReader
	gzipCmd.Stdout, err = os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}

	fmt.Println("Saving Docker image as gzip...")
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting docker save: %v\n", err)
		return
	}
	if err := gzipCmd.Start(); err != nil {
		fmt.Printf("Error starting gzip: %v\n", err)
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error waiting for docker save: %v\n", err)
	}
	pipeWriter.Close()
	if err := gzipCmd.Wait(); err != nil {
		fmt.Printf("Error waiting for gzip: %v\n", err)
	} else {
		fmt.Println("Docker image saved as gzip successfully!")
	}
}
