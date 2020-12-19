//+build !windows

package applist

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func Files() []string {
	return _files("/Applications/Google Chrome.app/")
} //end Files

func IsExecutable(fileInfo os.FileInfo) bool {
	return fileInfo.Mode().Perm()&0111 > 0
} //end IsExecutable

func IsVerified(fileName string) bool {
	//filename := "/Applications/Microsoft Word.app/Contents/MacOS/Microsoft Word"
	//filename := "/usr/bin/grep"

	out, err := exec.Command("codesign", "--verify", fileName).Output()
	if err != nil {
		output := string(out[:])
		fmt.Print(output)
	} else {
		output := string(out[:])
		fmt.Print(output)
	}
	return err == nil
} //end IsVerified

func JoinPath(dir string, name string) string{
	return path.Join(dir, name)
} //end JoinPath