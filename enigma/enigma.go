package enigma
import "unicode"
import "strconv"

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
    trace_letters bool

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
func convertLetter(letter rune) int {
    return int(letter) - 65
}

func convertNumber(number int) rune {
    return rune('A' + number)
}

func hasRune(needle rune, haystack []rune) bool {
    for _, item := range haystack {
        if(item == needle) {
            return true
        }
    }

    return false
}

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
func (enigma *enigma) GetRotorPositions() string {
    var return_value = ""
    for i := len(enigma.current_rotors) - 1; i >= 0; i-- {
        return_value += strconv.Itoa(enigma.current_rotors[i]) + "(" + string(convertNumber(enigma.rotors[enigma.current_rotors[i] - 1].position)) + "), "
    }

    //trim last 2 characters
    return_value = return_value[0:len(return_value) - 2]

    return return_value
}

func (enigma *enigma) SetRotorOrder(new_rotor_order []int) bool {
    for _, rotor_position := range new_rotor_order {
        if(!hasInt(rotor_position, enigma.available_rotors)) {
            return false
        }
    }

    enigma.current_rotors = new_rotor_order

    return true
}

func (enigma *enigma) SetRotorPosition(new_rotor_position []rune) bool {
    for index, rotor_position := range new_rotor_position {
        if(!unicode.IsLetter(rotor_position)) {
            return false
        }
        
        for {
            if(convertNumber(enigma.getRotor(len(enigma.current_rotors) - index).position) != rotor_position) {
                enigma.incrementRotor(len(enigma.current_rotors) - index, false)
            } else {
                break
            }
        }
    }

    return true
}

func (enigma *enigma) GetReflector() string {
    return string(enigma.current_reflector)
}

func (enigma *enigma) SetReflector(new_reflector_setting rune) bool {
    if(hasRune(new_reflector_setting, enigma.available_reflectors)) {
        enigma.current_reflector = new_reflector_setting
        return true
    }

    return false
}

func (enigma *enigma) GetPlugboard() string {
    return string(enigma.current_plugboard)
}

func (enigma *enigma) SetPlugboard(new_plugboard_setting rune) bool {
    if(hasRune(new_plugboard_setting, enigma.available_plugboards)) {
        enigma.current_plugboard = new_plugboard_setting
        return true
    }

    return false
}

func (enigma enigma) passThroughPlugboard(letter rune) rune {
    return enigma.plugboards[convertLetter(enigma.current_plugboard)].plugboard[convertLetter(letter)]
}

func (enigma enigma) passThroughRotor(letter rune, rotor_position int) rune {
    return enigma.getRotor(rotor_position).pins[convertLetter(letter)]
}

func (enigma enigma) passThroughReflector(letter rune) rune {
    return enigma.reflectors[convertLetter(enigma.current_reflector)].reflector[convertLetter(letter)]
}

func (enigma enigma) passThroughRotorReverse(letter rune, rotor_position int) rune {
    for index, pin := range enigma.getRotor(rotor_position).pins {
        if(pin == letter) {
            return convertNumber(index)
        }
    }

    return '0'
}

func (enigma *enigma) getRotor(rotor_position int) rotor {
    return enigma.rotors[enigma.current_rotors[rotor_position - 1] - 1]
}

func (enigma *enigma) rotateRotorValues(rotor_position int) {
    var saved_first_letter = enigma.getRotor(rotor_position).pins[0]

    for i := 0; i < 25; i++ {
        enigma.getRotor(rotor_position).pins[i] = enigma.getRotor(rotor_position).pins[i + 1]
    }

    enigma.getRotor(rotor_position).pins[25] = saved_first_letter

    for i := 0; i < 26; i++ {
        if(enigma.getRotor(rotor_position).pins[i] != 'A') {
            enigma.getRotor(rotor_position).pins[i] -= 1
        } else {
            enigma.getRotor(rotor_position).pins[i] = 'Z'
        }
    }
}

func (enigma *enigma) incrementRotor(rotor_position int, check_knockpoint bool) {
    enigma.rotors[enigma.current_rotors[rotor_position - 1] - 1].position += 1

    if(enigma.getRotor(rotor_position).position >= 26) {
        enigma.rotors[enigma.current_rotors[rotor_position - 1] - 1].position %= 26
    }

    enigma.rotateRotorValues(rotor_position)

    if(enigma.getRotor(rotor_position).position == convertLetter(enigma.getRotor(rotor_position).knockpoint) && len(enigma.current_rotors) > rotor_position && check_knockpoint) {
        enigma.incrementRotor(rotor_position + 1, true)
    }
}

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
