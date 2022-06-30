package create

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

func CmdCreate(cmdName string, repoName string) *cobra.Command {
	cmdName = fmt.Sprintf("%sCmd", cmdName)

	cmd := &cobra.Command{
		Use: cmdName,
		Run: func(cmd *cobra.Command, args []string) {
			cmdd := exec.Command("go", "get", "-u", repoName)
			cmdd.Stdout = os.Stdout
			cmdd.Stderr = os.Stderr

			if err := cmdd.Run(); err != nil {
				log.Fatal(err)
			}
			log.Println("all downloads are done.")
		},
	}
	return cmd
}
