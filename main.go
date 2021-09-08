package main

import "bufio"
import "os"
import "fmt"
import "unicode"
import "strconv"
import "time"
import "path/filepath"
import "example.com/enigma"
import "github.com/manifoldco/promptui"

type settings_file struct {
    path string
    data string
}

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

    printSettings := func() {
        fmt.Print("\nCurrent settings: ")
        fmt.Print("P[" + enigma.GetPlugboard() + "] ")
        fmt.Print("R[" + strconv.Itoa(enigma.GetRotorOrder(3)) + "(" + string(enigma.GetRotorPosition(3)) + "), " + strconv.Itoa(enigma.GetRotorOrder(2)) + "(" + string(enigma.GetRotorPosition(2)) + "), " + strconv.Itoa(enigma.GetRotorOrder(1)) + "(" + string(enigma.GetRotorPosition(1)) + ")] ")
        fmt.Print("R[" + enigma.GetReflector() + "]\n")
    }

    getSettingsString := func() string {
        return enigma.GetPlugboard() + "," + strconv.Itoa(enigma.GetRotorOrder(3)) + "," + string(enigma.GetRotorPosition(3)) + "," + strconv.Itoa(enigma.GetRotorOrder(2)) + "," + string(enigma.GetRotorPosition(2)) + "," + strconv.Itoa(enigma.GetRotorOrder(1)) + "," + string(enigma.GetRotorPosition(1)) + "," + enigma.GetReflector()
    }

    for {
        printSettings()

        prompt := promptui.Select{
    		Label: "Select option",
    		Items: []string{"Encrypt", "Change Plugboard", "Change Rotors", "Change Reflector", "Saved Settings", "Quit"},
    	}

    	_, action, _ := prompt.Run()

        if(action == "Encrypt") {
            var encryption_string = getText("Text to encrypt")

            fmt.Print("Encrypted text:  ")

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
                printSettings()

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
                        fmt.Println("Successfully changed rotor order to [" + strconv.Itoa(rotor_3) + "," + strconv.Itoa(rotor_2) + "," + strconv.Itoa(rotor_1) + "]")
                    } else {
                        fmt.Println("Rotar order [" + strconv.Itoa(rotor_3) + "," + strconv.Itoa(rotor_2) + "," + strconv.Itoa(rotor_1) + "] is not valid.")
                    }
                } else if(rotor_action == "Change Rotor Positions") {
                    var rotor_1 = unicode.ToUpper(rune(getText("Rotor 1 position")[0]))
                    var rotor_2 = unicode.ToUpper(rune(getText("Rotor 2 position")[0]))
                    var rotor_3 = unicode.ToUpper(rune(getText("Rotor 3 position")[0]))

                    if(enigma.SetRotorPosition([]rune{rotor_3, rotor_2, rotor_1}) == true) {
                        fmt.Println("Successfully changed rotor positions to [" + string(rotor_3) + "," + string(rotor_2) + "," + string(rotor_1) + "]")
                    } else {
                        fmt.Println("Rotar position [" + string(rotor_3) + "," + string(rotor_2) + "," + string(rotor_1) + "] is not valid.")
                    }
                } else {
                    break
                }
            }
        } else if(action == "Change Reflector") {
            var reflector = unicode.ToUpper(rune(getText("Reflector selection")[0]))

            if(enigma.SetReflector(reflector) == true) {
                fmt.Println("Successfully changed to reflector", string(reflector))
            } else {
                fmt.Println("Reflector", string(reflector), "is not valid.")
            }
        } else if(action == "Saved Settings") {
            for {
                saved_settings := []settings_file{}
                already_saved := false
                options_array := []string{}

                filepath.Walk("io", func(path string, info os.FileInfo, err error) error {
                    data, _ := os.ReadFile(path)

                    if(path != "io") {
                        saved_settings = append(saved_settings, settings_file {
                            path: path,
                            data: string(data),
                        })

                        if(string(data) == getSettingsString()) {
                            already_saved = true
                        }
                    }

                    return nil
                })

                if(len(saved_settings) == 0) {
                    options_array = []string{"Save Current Settings", "Back"}
                } else if(already_saved) {
                    options_array = []string{"Switch to a Saved Setting", "Delete Settings", "Back"}
                } else {
                    options_array = []string{"Save Current Settings", "Switch to a Saved Setting", "Delete Settings", "Back"}
                }

                printSettings()

                saved_settings_prompt := promptui.Select{
                    Label: "Select option",
                    Items: options_array,
                }

                _, settings_action, _ := saved_settings_prompt.Run()

                if(settings_action == "Save Current Settings") {
                    now := time.Now()
                    sec := now.UnixNano()

                    f, _ := os.Create("io/saved_settings_" + strconv.FormatInt(sec,10))
                    f.WriteString(getSettingsString())
                } else if(settings_action == "Switch to a Saved Setting") {
                    switch_options_array := []string{"Cancel"}
                    for _, saved_setting := range saved_settings {
                        switch_options_array = append(switch_options_array, saved_setting.data)
                    }

                    switch_prompt := promptui.Select{
                        Label: "Which saved setting would you like to switch to?",
                        Items: switch_options_array,
                    }

                    _, switch_action, _ := switch_prompt.Run()

                    for _, saved_setting := range saved_settings {
                        if(switch_action == saved_setting.data) {
                            enigma.SetPlugboard(rune(saved_setting.data[0]))
                            enigma.SetRotorOrder([]int{int(saved_setting.data[2]), int(saved_setting.data[6]), int(saved_setting.data[10])})
                            enigma.SetRotorPosition([]rune{rune(saved_setting.data[4]), rune(saved_setting.data[8]), rune(saved_setting.data[12])})
                            enigma.SetReflector(rune(saved_setting.data[14]))
                        }
                    }
                } else if(settings_action == "Delete Settings") {
                    delete_options_array := []string{"Cancel"}
                    for _, saved_setting := range saved_settings {
                        delete_options_array = append(delete_options_array, saved_setting.data)
                    }

                    delete_prompt := promptui.Select{
                        Label: "Which saved setting would you like to delete?",
                        Items: delete_options_array,
                    }

                    _, delete_action, _ := delete_prompt.Run()

                    for _, saved_setting := range saved_settings {
                        if(delete_action == saved_setting.data) {
                            os.Remove(saved_setting.path)
                        }
                    }
                } else {
                    break
                }
            }
        } else {
            break
        }
    }
}
