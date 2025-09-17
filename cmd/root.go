
package cmd

import (
	"fmt"
	//"io"
	"log"
	"os"
	"github.com/spf13/cobra"
	"bufio"
	//"strings"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cccat",
	Short: "Unix Command Line Tool CAT",
	Long: `The cat utility reads files sequentially, writing them to the standard output.
The file operands are processed in command-line order.  If file is a single
dash (‘-’) or absent, cat reads from the standard input`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "-":
			reader:=bufio.NewReader(os.Stdin)
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(line)
		case "n":
			scanner := bufio.NewScanner(os.Stdin)
			i:=1
			for scanner.Scan() {
				fmt.Println(i, " ", scanner.Text()) // Println will add back the final '\n'
				i++
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		case "b":
			scanner := bufio.NewScanner(os.Stdin)
			i:=1
			for scanner.Scan() {
				if scanner.Text() != "" {
				fmt.Println(i, scanner.Text()) // Println will add back the final '\n'
				i++
				}
				if scanner.Text() == "" { 
				fmt.Println(scanner.Text()) // Println will add back the final '\n'
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		default: 
			for _, file := range args{
				dat, err := os.ReadFile(file)
    			check(err)
    			fmt.Print(string(dat))
			}
		}
	 },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cccat.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


