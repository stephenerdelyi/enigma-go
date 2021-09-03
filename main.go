package main

import "bufio"
import "os"
import "fmt"
import "unicode"
import "example.com/enigma"
import "github.com/manifoldco/promptui"

func getText(label string) string {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print(label, ": ")
    if scanner.Scan() {
        input := scanner.Text()
        return input
    }

    return ""
}

func main() {
    enigma := enigma.Enigma()
    //enigma.TraceLetters(true)

    for {
        fmt.Println("Current settings: P[" + enigma.GetPlugboard() + "] R[" + enigma.GetRotorPositions() + "] R[" + enigma.GetReflector() + "]")

        prompt := promptui.Select{
    		Label: "Select option",
    		Items: []string{"Encrypt", "Change Plugboard", "Change Rotors", "Change Reflector", "Quit"},
    	}

    	_, action, err := prompt.Run()

        if err != nil {
    		fmt.Printf("Prompt failed %v\n", err)
    		return
    	} else if(action == "Encrypt") {
            var encryption_string = getText("Text to encrypt")

            for _, letter := range encryption_string {
                if(unicode.IsLetter(letter)) {
                    fmt.Print(string(enigma.Encrypt(rune(letter))))
                } else {
                    fmt.Print(string(letter))
                }
            }
            fmt.Println()
        } else if(action == "Change Plugboard") {
            var plugboard = rune(getText("Plugboard selection")[0])

            if(enigma.SetPlugboard(plugboard) == true) {
                fmt.Println("Successfully changed to plugboard", string(plugboard))
            } else {
                fmt.Println("Plugboard", string(plugboard), "is not valid.")
            }
        } else if(action == "Change Rotors") {
            var rotor_1 = int(getText("Rotor 1 selection")[0]) - 48
            var rotor_2 = int(getText("Rotor 2 selection")[0]) - 48
            var rotor_3 = int(getText("Rotor 3 selection")[0]) - 48

            fmt.Println(rotor_1)

            if(enigma.SetRotorPositions([]int{rotor_1, rotor_2, rotor_3}) == true) {
                fmt.Println("Successfully changed rotors to [", rotor_1, ",", rotor_2, ",", rotor_3, "]")
            } else {
                fmt.Println("Rotar selection [", rotor_1, ",", rotor_2, ",", rotor_3, "] is not valid.")
            }
        }  else if(action == "Change Reflector") {
            var reflector = rune(getText("Reflector selection")[0])

            if(enigma.SetReflector(reflector) == true) {
                fmt.Println("Successfully changed to reflector", string(reflector))
            } else {
                fmt.Println("Reflector", string(reflector), "is not valid.")
            }
        } else {
            break
        }
    }
}
