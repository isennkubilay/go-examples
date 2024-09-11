package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func main() {
	for {
		fmt.Println("\nDevOps Daily CLI")
		fmt.Println("1. Docker Commands")
		fmt.Println("2. Kubernetes Commands")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			dockerMenu()
		case 2:
			kubernetesMenu()
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func dockerMenu() {
	for {
		fmt.Println("\nDocker Commands")
		fmt.Println("1. Docker Login")
		fmt.Println("2. Docker Logs")
		fmt.Println("3. Docker Build")
		fmt.Println("4. Docker Push")
		fmt.Println("5. Docker Pull")
		fmt.Println("6. Docker Save (gzip)")
		fmt.Println("7. Docker Load")
		fmt.Println("8. Back to Main Menu")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			dockerLogin()
		case 2:
			dockerLogs()
		case 3:
			dockerBuild()
		case 4:
			dockerPush()
		case 5:
			dockerPull()
		case 6:
			dockerSave()
		case 7:
			dockerLoad()
		case 8:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func kubernetesMenu() {
	for {
		fmt.Println("\nKubernetes Commands")
		fmt.Println("1. Get Pods")
		fmt.Println("2. Get Logs from Pod")
		fmt.Println("3. Delete Pod")
		fmt.Println("4. Get Services")
		fmt.Println("5. Apply Manifest")
		fmt.Println("6. Back to Main Menu")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			getPods()
		case 2:
			getLogsFromPod()
		case 3:
			deletePod()
		case 4:
			getServices()
		case 5:
			applyManifest()
		case 6:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func dockerLogin() {
	var username, server string

	fmt.Print("Enter Docker registry server (press Enter for Docker Hub): ")
	fmt.Scanln(&server)

	fmt.Print("Enter username: ")
	fmt.Scan(&username)

	fmt.Print("Enter password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("Error reading password: %v\n", err)
		return
	}
	password := string(bytePassword)
	fmt.Println() // Print a newline after password input

	var cmd *exec.Cmd
	if server == "" {
		cmd = exec.Command("docker", "login", "-u", username, "--password-stdin")
	} else {
		cmd = exec.Command("docker", "login", server, "-u", username, "--password-stdin")
	}

	cmd.Stdin = strings.NewReader(password)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error logging in to Docker registry: %v\n", err)
	} else {
		fmt.Println("Successfully logged in to Docker registry")
	}
}

func dockerLogs() {
	fmt.Print("Enter container name or ID: ")
	var container string
	fmt.Scan(&container)

	cmd := exec.Command("docker", "logs", container)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func dockerBuild() {
	var dockerfilePath, imageTag, contextPath string

	fmt.Print("Enter Dockerfile name (e.g., Dockerfile): ")
	fmt.Scan(&dockerfilePath)

	fmt.Print("Enter image name and tag (e.g., myimage:latest): ")
	fmt.Scan(&imageTag)

	fmt.Print("Enter build context path (press Enter for current directory): ")
	fmt.Scanln(&contextPath)

	if contextPath == "" {
		contextPath = "."
	}

	// Ensure Dockerfile exists
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		fmt.Printf("Error: Dockerfile '%s' not found in the current directory.\n", dockerfilePath)
		return
	}

	cmd := exec.Command("docker", "build", "-t", imageTag, "-f", dockerfilePath, contextPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Building Docker image with command: docker build -t %s -f %s %s\n", imageTag, dockerfilePath, contextPath)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error building Docker image: %v\n", err)
		fmt.Println("Make sure Docker is running and you have the necessary permissions.")
	} else {
		fmt.Printf("Successfully built Docker image: %s\n", imageTag)
	}
}

func dockerPush() {
	fmt.Print("Enter image name and tag to push: ")
	var imageTag string
	fmt.Scan(&imageTag)

	cmd := exec.Command("docker", "push", imageTag)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func dockerPull() {
	fmt.Print("Enter image name and tag to pull: ")
	var imageTag string
	fmt.Scan(&imageTag)

	cmd := exec.Command("docker", "pull", imageTag)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func dockerSave() {
	var imageName, outputFile string

	fmt.Print("Enter the image name to save: ")
	fmt.Scan(&imageName)

	fmt.Print("Enter the output file name (including .tar.gz extension): ")
	fmt.Scan(&outputFile)

	if !strings.HasSuffix(outputFile, ".tar.gz") {
		outputFile += ".tar.gz"
	}

	fmt.Printf("Saving Docker image '%s' to '%s'...\n", imageName, outputFile)

	cmd := exec.Command("sh", "-c", fmt.Sprintf("docker save %s | gzip > %s", imageName, outputFile))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error saving Docker image: %v\n", err)
	} else {
		fmt.Printf("Successfully saved Docker image '%s' to '%s'\n", imageName, outputFile)
	}
}

func dockerLoad() {
	var inputFile string

	fmt.Print("Enter the path to the Docker image file (.tar or .tar.gz): ")
	fmt.Scan(&inputFile)

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("Error: File '%s' not found.\n", inputFile)
		return
	}

	var cmd *exec.Cmd
	if strings.HasSuffix(inputFile, ".gz") {
		fmt.Printf("Loading gzipped Docker image from '%s'...\n", inputFile)
		cmd = exec.Command("sh", "-c", fmt.Sprintf("gunzip -c %s | docker load", inputFile))
	} else {
		fmt.Printf("Loading Docker image from '%s'...\n", inputFile)
		cmd = exec.Command("docker", "load", "-i", inputFile)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error loading Docker image: %v\n", err)
	} else {
		fmt.Println("Successfully loaded Docker image")
	}
}

func getPods() {
	fmt.Println("Getting Kubernetes pods...")

	var namespace string
	fmt.Print("Enter namespace (press Enter for default): ")
	fmt.Scanln(&namespace)

	var cmd *exec.Cmd
	if namespace == "" {
		cmd = exec.Command("kubectl", "get", "pods")
	} else {
		cmd = exec.Command("kubectl", "get", "pods", "-n", namespace)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error getting pods: %v\n", err)
	}
}

func getLogsFromPod() {
	var podName, namespace string

	fmt.Print("Enter pod name: ")
	fmt.Scan(&podName)

	fmt.Print("Enter namespace (press Enter for default): ")
	fmt.Scanln(&namespace)

	var cmd *exec.Cmd
	if namespace == "" {
		cmd = exec.Command("kubectl", "logs", podName)
	} else {
		cmd = exec.Command("kubectl", "logs", podName, "-n", namespace)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error getting logs from pod: %v\n", err)
	}
}

func deletePod() {
	var podName, namespace string

	fmt.Print("Enter pod name to delete: ")
	fmt.Scan(&podName)

	fmt.Print("Enter namespace (press Enter for default): ")
	fmt.Scanln(&namespace)

	var cmd *exec.Cmd
	if namespace == "" {
		cmd = exec.Command("kubectl", "delete", "pod", podName)
	} else {
		cmd = exec.Command("kubectl", "delete", "pod", podName, "-n", namespace)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error deleting pod: %v\n", err)
	} else {
		fmt.Printf("Pod '%s' deleted successfully\n", podName)
	}
}

func getServices() {
	// Implementation for getting Kubernetes services
	fmt.Println("Getting Kubernetes services...")
	// Add kubectl command execution here
}

func applyManifest() {
	// Implementation for applying Kubernetes manifest
	fmt.Println("Applying Kubernetes manifest...")
	// Add kubectl apply command execution here
}

func deleteResource() {
	// Implementation for deleting Kubernetes resource
	fmt.Println("Deleting Kubernetes resource...")
	// Add kubectl delete command execution here
}
