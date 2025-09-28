package databases

import (
	"bufio"
	"docker-db-management/types"
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
)

type MySQL struct {
	LatestImage  bool
	Password     string
	DatabaseName string
}

var _ DBHandler = &MySQL{}

const (
	mysqlImage = "mysql:8.0" // You can make this configurable later
)

func (m *MySQL) SetConfig(config types.Config) error {
	m.LatestImage = config.LatestImage
	m.Password = config.Password
	m.DatabaseName = config.DatabaseName
	return nil
}

func (m MySQL) Create() error {
	if !isDockerAvailable() {
		return fmt.Errorf("docker is not installed or not in PATH")
	}

	fmt.Println("Found docker")

	cmd := exec.Command("docker", "images", "--format", "{{.Repository}}:{{.Tag}}")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to list docker images: %w", err)
	}

	var mySqlImage string
	hasMySQL := false

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		image := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(image, ":", 2)
		if len(parts) != 2 {
			continue
		}
		repo, tag := parts[0], parts[1]

		if repo == "mysql" && tag != "<none>" {
			hasMySQL = true
			if tag == "latest" {
				mySqlImage = "mysql:latest"
				break
			}
			if mySqlImage == "" {
				mySqlImage = image
			}
		}
	}

	if !hasMySQL || m.LatestImage {
		fmt.Println("MySQL image not found locally. Pulling mysql:latest...")
		cmd = exec.Command("docker", "pull", "mysql:latest")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to pull mysql image: %w", err)
		}
	} else {
		fmt.Println("MySQL image found locally.")
	}

	containerName := "mysql-" + strconv.Itoa(rand.Intn(10000))
	fmt.Printf("Starting MySQL container '%s' with database '%s'...\n", containerName, m.DatabaseName)

	args := []string{
		"run", "-d",
		"--name", containerName,
		"-e", "MYSQL_ROOT_PASSWORD=" + m.Password,
	}

	if m.DatabaseName != "" {
		args = append(args, "-e", "MYSQL_DATABASE="+m.DatabaseName)
	}

	mySqlPort := 3306
	for {
		if isPortInUse(mySqlPort) {
			mySqlPort++
		} else {
			break
		}
	}

	args = append(args, "-p", fmt.Sprintf("127.0.0.1:%d:3306", mySqlPort))
	args = append(args, mySqlImage)

	fmt.Println(args)
	cmd = exec.Command("docker", args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start MySQL container: %w", err)
	}

	fmt.Println("MySQL container started and database created successfully!")
	return nil
}

func (m MySQL) Remove(name string) error {
	fmt.Printf("Removing MySQL database: %s\n", name)
	return nil
}
