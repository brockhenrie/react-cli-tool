/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
    "log"
    "os"
)

// randomCmd represents the random command
var componentCmd = &cobra.Command{
	Use:   "component",
    Short: "A react component",
	Long: `Generates a react component`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(args)
        generateComponent(args[0], cmd)
	},
}

func init() {

 	rootCmd.AddCommand(componentCmd)
   componentCmd.Flags().BoolP("tsx","t", false,"Use flag to change output type to typescript")
   componentCmd.Flags().BoolP("dry","d", false,"Use flag to output component to console and not file")

}

type Component struct {
    Name      string
    Function  bool
    Ts        bool
}


func generateComponent(name string, cmd *cobra.Command){
    var fileName string
    var component []string
    if cmd.Flag("dry").Changed && cmd.Flag("tsx").Changed{
        component = generateTsComp(&name)
        printComp(&component)
        return
    }

    if cmd.Flag("tsx").Changed {
        fileName = name + ".tsx"
        component = generateTsComp(&name)
        writeComp(&fileName, &component)
        return
    }

    fileName = name + ".jsx"
    component = generateJsComp(&name)

    if cmd.Flag("dry").Changed {
        printComp(&component)
        return
    }

    writeComp(&fileName, &component)
    return

}

func writeComp(fileName *string, component *[]string){
  f, err := os.Create(*fileName)
    if err != nil{
        log.Printf("Could not create file - %v", err)
    }

    defer f.Close()

    for _, line := range *component {
        f.WriteString(line + "\n")
    }

    fmt.Println("File Created!")
    return
}


func printComp(comp *[]string){
    for _,line := range *comp {
    fmt.Println(line)
    }
}


func generateTsComp(name *string) []string{

    return []string{
    `import 'react' from 'react'`,
    ` `,
    fmt.Sprintf("export const %s: react.RC = (props) => { ",*name),
    ` `,
    `   return ( `,
    `<div>`,
    fmt.Sprintf("Component %s works!!!",*name),
    `   </div>`,
    `) `,
    `} `,
    ` `,
}

}


func generateJsComp(name *string) []string{
    return []string{
    `import 'react' from 'react'`,
    ` `,
    fmt.Sprintf("export const %s = (props) => { ",*name),
    ` `,
    `   return ( `,
    `<div>`,
    fmt.Sprintf("Component %s works!!!",*name),
    `   </div>`,
    `) `,
    `} `,
    ` `,
}
}















