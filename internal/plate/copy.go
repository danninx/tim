package plate

import (
	"fmt"
	"os"
	"os/exec"
)

func GitClone(src string, dest string) error {
	fmt.Printf("using git to copy source \"%v\"\n", src)	
	cmd := exec.Command("git", "clone", src, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd)
	return cmd.Run()	
}
