package enigma
import "unicode"

/////////////////////////////////////////////////////////////////
// STRUCTS
/////////////////////////////////////////////////////////////////
type rotor struct {
    pins []rune
    knockpoint rune
    position int
}

type reflector struct {
    reflector []rune
}

type plugboard struct {
    plugboard []rune
}

type enigma struct {
    available_rotors []int
    current_rotors []int
    rotors []rotor

    available_reflectors []rune
    current_reflector rune
    reflectors []reflector

    available_plugboards []rune
    current_plugboard rune
    plugboards []plugboard
}

/////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS
/////////////////////////////////////////////////////////////////
//convertLetter - converts rune character to its numerical representation (A = 1, B = 2...)
func convertLetter(letter rune) int {
    return int(letter) - 65
}

//convertNumber - converts integer to its alphabetical representation (1 = A, 2 = B...)
func convertNumber(number int) rune {
    return rune('A' + number)
}

//hasRune - returns whether or not an array of runes contains a rune
func hasRune(needle rune, haystack []rune) bool {
    for _, item := range haystack {
        if(item == needle) {
            return true
        }
    }

    return false
}

//hasRune - returns whether or not an array of ints contains an int
func hasInt(needle int, haystack []int) bool {
    for _, item := range haystack {
        if(item == needle) {
            return true
        }
    }

    return false
}

/////////////////////////////////////////////////////////////////
// ENIGMA FUNCTIONS
/////////////////////////////////////////////////////////////////
//GetRotorOrder - "Public" getter to obtain current rotor order with a given rotor number
func (enigma *enigma) GetRotorOrder(rotor_position int) int {
    //If rotor order setting is not valid
    if(rotor_position < 1 || rotor_position > 3) {
        return '0'
    }

    return enigma.current_rotors[rotor_position - 1]
}

//GetRotorPosition - "Public" getter to obtain current rotor position with a given rotor number
func (enigma *enigma) GetRotorPosition(rotor_position int) rune {
    //If rotor position setting is not valid
    if(rotor_position < 1 || rotor_position > 3) {
        return '0'
    }

    return convertNumber(enigma.getRotor(rotor_position).position)
}

//SetRotorOrder - "Public" setter to change all active rotor orders
func (enigma *enigma) SetRotorOrder(new_rotor_order []int) bool {
    //For each new rotor order provided
    for _, rotor_position := range new_rotor_order {
        //If rotor order setting is not valid
        if(!hasInt(rotor_position, enigma.available_rotors)) {
            return false
        }
    }

    enigma.current_rotors = new_rotor_order

    return true
}

//SetRotorPosition - "Public" setter to change all active rotor positions
func (enigma *enigma) SetRotorPosition(new_rotor_position []rune) bool {
    //For each new rotor position provided
    for index, rotor_position := range new_rotor_position {
        //Return false if not a letter
        if(!unicode.IsLetter(rotor_position)) {
            return false
        }

        //While the given rotor is not in the correct position
        for {
            if(convertNumber(enigma.getRotor(len(enigma.current_rotors) - index).position) != rotor_position) {
                //Increment the rotor
                enigma.incrementRotor(len(enigma.current_rotors) - index, false)
            } else {
                break
            }
        }
    }

    return true
}

//GetReflector - "Public" getter to obtain the current reflector in use
func (enigma *enigma) GetReflector() string {
    return string(enigma.current_reflector)
}

//SetReflector - "Public" setter to change the current reflector in use
func (enigma *enigma) SetReflector(new_reflector_setting rune) bool {
    //If reflector setting provided is valid
    if(hasRune(new_reflector_setting, enigma.available_reflectors)) {
        enigma.current_reflector = new_reflector_setting
        return true
    }

    return false
}

//GetPlugboard - "Public" getter to obtain the current plugboard in use
func (enigma *enigma) GetPlugboard() string {
    return string(enigma.current_plugboard)
}

//SetPlugboard - "Public" setter to change the current plugboard in use
func (enigma *enigma) SetPlugboard(new_plugboard_setting rune) bool {
    //If plugboard setting provided is valid
    if(hasRune(new_plugboard_setting, enigma.available_plugboards)) {
        enigma.current_plugboard = new_plugboard_setting
        return true
    }

    return false
}

//passThroughPlugboard - "Private" function which passes a letter through the current plugboard during encryption
func (enigma enigma) passThroughPlugboard(letter rune) rune {
    return enigma.plugboards[convertLetter(enigma.current_plugboard)].plugboard[convertLetter(letter)]
}

//passThroughRotor - "Private" function which passes a letter through the given rotor during encryption
func (enigma enigma) passThroughRotor(letter rune, rotor_position int) rune {
    return enigma.getRotor(rotor_position).pins[convertLetter(letter)]
}

//passThroughReflector - "Private" function which passes a letter through the current reflector during encryption
func (enigma enigma) passThroughReflector(letter rune) rune {
    return enigma.reflectors[convertLetter(enigma.current_reflector)].reflector[convertLetter(letter)]
}

//passThroughRotorReverse - "Private" function which passes a letter (in reverse) through the given rotor during encryption
func (enigma enigma) passThroughRotorReverse(letter rune, rotor_position int) rune {
    //For each pin in our current rotor
    for index, pin := range enigma.getRotor(rotor_position).pins {
        //If the pin is equivalent to our current letter
        if(pin == letter) {
            //Return the alphabetical value of our current index
            return convertNumber(index)
        }
    }

    return '0'
}

//getRotor - "Private" function which returns the active rotor in use for the given position
func (enigma *enigma) getRotor(rotor_position int) rotor {
    return enigma.rotors[enigma.current_rotors[rotor_position - 1] - 1] //subtract by one because indexes start at 0
}

//rotateRotorValues - "Private" function which rotates and converts a given rotor array, simulating the rotation of a single rotor
func (enigma *enigma) rotateRotorValues(rotor_position int) {
    //Save our first letter
    var saved_first_letter = enigma.getRotor(rotor_position).pins[0]

    //Move each pin (letter) to the previous index
    for i := 0; i < 25; i++ {
        enigma.getRotor(rotor_position).pins[i] = enigma.getRotor(rotor_position).pins[i + 1]
    }

    //Add our saved letter to the end of the pins array
    enigma.getRotor(rotor_position).pins[25] = saved_first_letter

    //Simulate rotation of a wheel by decrementing letter values by 1 (A => Z, B => A, C => B)
    for i := 0; i < 26; i++ {
        if(enigma.getRotor(rotor_position).pins[i] != 'A') {
            enigma.getRotor(rotor_position).pins[i] -= 1
        } else {
            enigma.getRotor(rotor_position).pins[i] = 'Z'
        }
    }
}

//incrementRotor - "Private" function which increments rotors during encryption and recursively when a knockpoint has been hit
func (enigma *enigma) incrementRotor(rotor_position int, check_knockpoint bool) {
    //Increment the current rotor position value
    enigma.rotors[enigma.current_rotors[rotor_position - 1] - 1].position += 1

    //Reset the current rotor position back to 1 when it reaches 26
    if(enigma.getRotor(rotor_position).position >= 26) {
        enigma.rotors[enigma.current_rotors[rotor_position - 1] - 1].position %= 26
    }

    //Simulate the actual movement of a rotor wheel
    enigma.rotateRotorValues(rotor_position)

    //Recursively increment the rotor in the next position if a knockpoint has been reached
    if(enigma.getRotor(rotor_position).position == convertLetter(enigma.getRotor(rotor_position).knockpoint) && len(enigma.current_rotors) > rotor_position && check_knockpoint) {
        enigma.incrementRotor(rotor_position + 1, true)
    }
}

//Encrypt - "Public" function which allows for encryption of a single character
func (enigma enigma) Encrypt(letter rune) rune {
    enigma.incrementRotor(1, true)

    letter = unicode.ToUpper(letter)
    letter = enigma.passThroughPlugboard(letter)
    letter = enigma.passThroughRotor(letter, 1)
    letter = enigma.passThroughRotor(letter, 2)
    letter = enigma.passThroughRotor(letter, 3)
    letter = enigma.passThroughReflector(letter)
    letter = enigma.passThroughRotorReverse(letter, 3)
    letter = enigma.passThroughRotorReverse(letter, 2)
    letter = enigma.passThroughRotorReverse(letter, 1)
    letter = enigma.passThroughPlugboard(letter)

    return letter
}

func Enigma() enigma {
    return enigma {
        //default settings
        current_rotors: []int{1,2,3},
        current_reflector: 'A',
        current_plugboard: 'A',
        //rotors, reflectors, plugboards
        available_rotors: []int{1,2,3},
        rotors: []rotor{
            rotor{
                []rune{'E','K','M','F','L','G','D','Q','V','Z','N','T','O','W','Y','H','X','U','S','P','A','I','B','R','C','J'},
                'Q',
                0,
            },
            rotor{
                []rune{'A','J','D','K','S','I','R','U','X','B','L','H','W','T','M','C','Q','G','Z','N','P','Y','F','V','O','E'},
                'E',
                0,
            },
            rotor{
                []rune{'B','D','F','H','J','L','C','P','R','T','X','V','Z','N','Y','E','I','W','G','A','K','M','U','S','Q','O'},
                'V',
                0,
            },
        },
        available_reflectors: []rune{'A','B','C'},
        reflectors: []reflector{
            reflector{
                []rune{'E','J','M','Z','A','L','Y','X','V','B','W','F','C','R','Q','U','O','N','T','S','P','I','K','H','G','D'},
            },
            reflector{
                []rune{'Y','R','U','H','Q','S','L','D','P','X','N','G','O','K','M','I','E','B','F','Z','C','W','V','J','A','T'},
            },
            reflector{
                []rune{'F','V','P','J','I','A','O','Y','E','D','R','Z','X','W','G','C','T','K','U','Q','S','B','N','M','H','L'},
            },
        },
        available_plugboards: []rune{'A','B','C'},
        plugboards: []plugboard{
            plugboard{
                []rune{'L','B','C','K','E','F','I','H','G','Z','D','A','O','X','M','S','W','Y','P','T','V','U','Q','N','R','J'},
            },
            plugboard{
                []rune{'A','M','C','D','E','F','K','H','I','J','G','L','B','S','O','P','Q','R','N','T','Z','V','W','X','Y','U'},
            },
            plugboard{
                []rune{'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'},
            },
        },
    }
}
