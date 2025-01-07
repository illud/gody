package ftp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jlaffaye/ftp"
)

func Ftp(ftpServer string, username string, password string, projectPath string, remoteProjectPath string) error {
	// Change the working directory to the project path
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		return err
	}

	err = os.Chdir(absPath)
	if err != nil {
		return err
	}

	// Connect to FTP server
	conn, err := ftp.Dial(ftpServer)
	if err != nil {
		fmt.Println("Error connecting to FTP server:", err)
		return err
	}
	defer conn.Quit()

	// Log in to the FTP server
	err = conn.Login(username, password)
	if err != nil {
		fmt.Println("Error logging in:", err)
		return err
	}

	// Recursively upload files
	err = uploadFolder(conn, absPath, remoteProjectPath) // "/" is the remote root directory
	if err != nil {
		fmt.Println("Error uploading folder:", err)
		return err
	}

	fmt.Println("All files and folders uploaded successfully.")

	return nil
}

// uploadFolder recursively uploads files and folders
func uploadFolder(conn *ftp.ServerConn, localFolder, remoteFolder string) error {
	// Walk through the local folder
	return filepath.Walk(localFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		// Compute relative path to maintain folder structure
		relPath, err := filepath.Rel(localFolder, path)
		if err != nil {
			return fmt.Errorf("error computing relative path: %v", err)
		}

		// Join with the remote folder
		remotePath := filepath.Join(remoteFolder, relPath)

		// Ensure the remote path uses forward slashes
		remotePath = strings.Replace(remotePath, "\\", "/", -1)

		// If it's a directory, create it on the FTP server
		if info.IsDir() {
			if err := conn.MakeDir(remotePath); err != nil {
				// Ignore errors if the directory already exists
				if !ftpErrIsDirExists(err) {
					return fmt.Errorf("error creating directory %s: %v", remotePath, err)
				}
			}
			return nil
		}

		// If it's a file, upload it
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("error opening file %s: %v", path, err)
		}
		defer file.Close()

		if err := conn.Stor(remotePath, file); err != nil {
			return fmt.Errorf("error uploading file %s: %v", remotePath, err)
		}

		fmt.Printf("Uploaded: %s to %s\n", path, remotePath)
		return nil
	})
}

// ftpErrIsDirExists checks if an error is because the directory already exists
func ftpErrIsDirExists(err error) bool {
	return err != nil && (err.Error() == "550 Can't create directory: File exists" || err.Error() == "550 File exists")
}
