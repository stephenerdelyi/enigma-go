package enigma
import "fmt"
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
    return rune('A' - 1 + number)
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
func (enigma *enigma) TraceLetters(trace_letters bool) {
    enigma.trace_letters = trace_letters
}

func (enigma enigma) trace(letter rune) {
    if(enigma.trace_letters) {
        fmt.Println(string(letter))
    }
}

func (enigma *enigma) GetRotorPositions() string {
    //TODO complete this once rotor position has been implemented
    var return_value = ""
    for _, rotor := range enigma.current_rotors {
        return_value += strconv.Itoa(rotor) + "(A), "
    }

    //trim last 2 characters
    return_value = return_value[0:len(return_value) - 2]

    return return_value
}

func (enigma *enigma) SetRotorPositions(new_rotor_settings []int) bool {
    for _, rotor_position := range new_rotor_settings {
        if(!hasInt(rotor_position, enigma.available_rotors)) {
            return false
        }
    }

    enigma.current_rotors = new_rotor_settings

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

func (enigma enigma) passThroughRotor(letter rune, rotor int) rune {
    return enigma.rotors[rotor - 1].pins[convertLetter(letter)]
}

func (enigma enigma) passThroughReflector(letter rune) rune {
    return enigma.reflectors[convertLetter(enigma.current_reflector)].reflector[convertLetter(letter)]
}

func (enigma enigma) passThroughRotorReverse(letter rune, rotor int) rune {
    for index, pin := range enigma.rotors[rotor - 1].pins {
        if(pin == letter) {
            return convertNumber(index + 1)
        }
    }

    return '0'
}

func (enigma *enigma) incrementRotor(step int) {
    //TODO implement this
}

func (enigma enigma) Encrypt(letter rune) rune {
    enigma.incrementRotor(1)

    letter = unicode.ToUpper(letter)
    enigma.trace(letter)
    letter = enigma.passThroughPlugboard(letter)
    enigma.trace(letter)
    letter = enigma.passThroughRotor(letter, enigma.current_rotors[0])
    enigma.trace(letter)
    letter = enigma.passThroughRotor(letter, enigma.current_rotors[1])
    enigma.trace(letter)
    letter = enigma.passThroughRotor(letter, enigma.current_rotors[2])
    enigma.trace(letter)
    letter = enigma.passThroughReflector(letter)
    enigma.trace(letter)
    letter = enigma.passThroughRotorReverse(letter, enigma.current_rotors[2])
    enigma.trace(letter)
    letter = enigma.passThroughRotorReverse(letter, enigma.current_rotors[1])
    enigma.trace(letter)
    letter = enigma.passThroughRotorReverse(letter, enigma.current_rotors[0])
    enigma.trace(letter)
    letter = enigma.passThroughPlugboard(letter)
    enigma.trace(letter)

    return letter
}

func Enigma() enigma {
    return enigma {
        //default settings
        current_rotors: []int{1,2,3},
        current_reflector: 'A',
        current_plugboard: 'A',
        trace_letters: false,
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
