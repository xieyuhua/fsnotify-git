/*
 * @Descripttion:
 * @version:
 * @Author: seaslog
 * @Date: 2022-03-04 15:26:58
 * @LastEditors: 谢余华
 * @LastEditTime: 2022-03-04 17:04:02
 */
package cmd

import (
	"fmt"
	"fsnotifyGit/config"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var flags config.Flags

var rootCmd = &cobra.Command{
	Use:   "cronPush",
	Short: "cronPush will watch the filepath and auto run the git command ",
	Long:  "1.cronPush will watch the filepath\n2.if the filepath has some change, auto run the git command 'git add . && git commit && git push'\n\ninput the path you want to watch as the flag",
	//Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("the path you input: %s", flags.Path)

		if flags.Path == "" {
			log.Println("you should input the -p flag. or you can watch the help menu -h")
			return
		}
		thisFunc(flags)
	},
}

//var cycleCmd = &cobra.Command{
//	Use:   "pushcycle",
//	Short: "git push once time each %n seconds; the default is 5s",
//	Long:  "",
//	Args:  cobra.MinimumNArgs(1),
//	Run: func(cmd *cobra.Command, args []string) {
//		log.Printf("the cycle time you input is: %s s",args[0])
//	},
//}

var thisFunc func(config.Flags)

func Execute(watcher func(config.Flags)) {
	thisFunc = watcher
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.PersistentFlags().IntVarP(&flags.Cycle, "pushCycle", "c", 5, "git push once time each %n seconds; the default is 5s")
	rootCmd.PersistentFlags().StringVarP(&flags.Path, "path", "p", "./", "input the path you want to watch as the flag;  the default is ./")
	//rootCmd.AddCommand(cycleCmd)
}
