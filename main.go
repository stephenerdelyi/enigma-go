package main

import "bufio"
import "os"
import "fmt"
import "unicode"
import "strconv"
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

    for {
        fmt.Println()
        fmt.Println("Current settings: P[" + enigma.GetPlugboard() + "] R[" + enigma.GetRotorPositions() + "] R[" + enigma.GetReflector() + "]")

        prompt := promptui.Select{
    		Label: "Select option",
    		Items: []string{"Encrypt", "Change Plugboard", "Change Rotors", "Change Reflector", "Quit"},
    	}

    	_, action, _ := prompt.Run()

        if(action == "Encrypt") {
            var encryption_string = getText("Text to encrypt")

            fmt.Print("Encrypted text: ");

            for _, letter := range encryption_string {
                if(unicode.IsLetter(letter)) {
                    fmt.Print(string(enigma.Encrypt(rune(letter))))
                } else {
                    fmt.Print(string(letter))
                }
            }
            fmt.Println()
        } else if(action == "Change Plugboard") {
            var plugboard = unicode.ToUpper(rune(getText("Plugboard selection")[0]))

            if(enigma.SetPlugboard(plugboard) == true) {
                fmt.Println("Successfully changed to plugboard", string(plugboard))
            } else {
                fmt.Println("Plugboard", string(plugboard), "is not valid.")
            }
        } else if(action == "Change Rotors") {
            for {
                rotor_prompt := promptui.Select{
                    Label: "Select option",
                    Items: []string{"Change Rotor Order", "Change Rotor Positions", "Back"},
                }

                _, rotor_action, _ := rotor_prompt.Run()

                if(rotor_action == "Change Rotor Order") {
                    var rotor_1 = int(getText("Rotor 1 selection")[0]) - 48
                    var rotor_2 = int(getText("Rotor 2 selection")[0]) - 48
                    var rotor_3 = int(getText("Rotor 3 selection")[0]) - 48

                    if(enigma.SetRotorOrder([]int{rotor_1, rotor_2, rotor_3}) == true) {
                        fmt.Println("Successfully changed rotor order to [" + strconv.Itoa(rotor_1) + "," + strconv.Itoa(rotor_2) + "," + strconv.Itoa(rotor_3) + "]")
                    } else {
                        fmt.Println("Rotar order [" + strconv.Itoa(rotor_1) + "," + strconv.Itoa(rotor_2) + "," + strconv.Itoa(rotor_3) + "] is not valid.")
                    }
                } else if(rotor_action == "Change Rotor Positions") {
                    var rotor_1 = rune(getText("Rotor 1 position")[0])
                    var rotor_2 = rune(getText("Rotor 2 position")[0])
                    var rotor_3 = rune(getText("Rotor 3 position")[0])

                    if(enigma.SetRotorPosition([]rune{rotor_1, rotor_2, rotor_3}) == true) {
                        fmt.Println("Successfully changed rotor positions to [" + string(rotor_1) + "," + string(rotor_2) + "," + string(rotor_3) + "]")
                    } else {
                        fmt.Println("Rotar position [" + string(rotor_1) + "," + string(rotor_2) + "," + string(rotor_3) + "] is not valid.")
                    }
                } else {
                    break
                }
            }
        }  else if(action == "Change Reflector") {
            var reflector = unicode.ToUpper(rune(getText("Reflector selection")[0]))

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
