package enigma
import "fmt"
import "unicode"

type rotor struct {
    pins [26]rune
    knockpoint rune
    position int
}

type reflector struct {
    reflector [26]rune
}

type plugboard struct {
    plugboard [26]rune
}

type enigma struct {
    current_rotors [3]int
    rotor_1 rotor
    rotor_2 rotor
    rotor_3 rotor
    trace_letters bool

    current_reflector rune
    reflector_A reflector
    reflector_B reflector
    reflector_C reflector

    current_plugboard rune
    plugboard_A plugboard
    plugboard_B plugboard
    plugboard_C plugboard
}

func (enigma *enigma) TraceLetters(trace_letters bool) {
    enigma.trace_letters = trace_letters;
}

func (enigma enigma) trace(letter rune) {
    if(enigma.trace_letters) {
        fmt.Println(string(letter));
    }
}

func (enigma *enigma) GetRotorPositions() string {
    //TODO complete this once rotor position has been implemented
    return "A, A, A";
}

func (enigma *enigma) SetRotorPositions(new_rotor_settings [3]int) bool {
    for _, rotor_position := range new_rotor_settings {
        if(rotor_position != 1 && rotor_position != 2 && rotor_position != 3) {
            return false;
        }
    }

    enigma.current_rotors = new_rotor_settings;

    return true;
}

func (enigma *enigma) GetReflector() string {
    return string(enigma.current_reflector);
}

func (enigma *enigma) SetReflector(new_reflector_setting rune) bool {
    if(new_reflector_setting == 'A' || new_reflector_setting == 'B' || new_reflector_setting == 'C') {
        enigma.current_reflector = new_reflector_setting;
        return true;
    }

    return false;
}

func (enigma *enigma) GetPlugboard() string {
    return string(enigma.current_plugboard);
}

func (enigma *enigma) SetPlugboard(new_plugboard_setting rune) bool {
    if(new_plugboard_setting == 'A' || new_plugboard_setting == 'B' || new_plugboard_setting == 'C') {
        enigma.current_plugboard = new_plugboard_setting;
        return true;
    }

    return false;
}

func (enigma enigma) convertLetter(letter rune) int {
    return int(letter) - 65;
}

func (enigma enigma) convertNumber(number int) rune {
    return rune('A' - 1 + number);
}

func (enigma enigma) passThroughPlugboard(letter rune) rune {
    if(enigma.current_plugboard == 'A') {
        return enigma.plugboard_A.plugboard[enigma.convertLetter(letter)];
    } else if(enigma.current_plugboard == 'B') {
        return enigma.plugboard_B.plugboard[enigma.convertLetter(letter)];
    } else if(enigma.current_plugboard == 'C') {
        return enigma.plugboard_C.plugboard[enigma.convertLetter(letter)];
    }

    return '0';
}

func (enigma enigma) passThroughRotor(letter rune, rotor int) rune {
    if(rotor == 1) {
        return enigma.rotor_1.pins[enigma.convertLetter(letter)];
    } else if(rotor == 2) {
        return enigma.rotor_2.pins[enigma.convertLetter(letter)];
    } else if(rotor == 3) {
        return enigma.rotor_3.pins[enigma.convertLetter(letter)];
    }

    return '0';
}

func (enigma enigma) passThroughReflector(letter rune) rune {
    if(enigma.current_reflector == 'A') {
        return enigma.reflector_A.reflector[enigma.convertLetter(letter)];
    } else if(enigma.current_reflector == 'B') {
        return enigma.reflector_B.reflector[enigma.convertLetter(letter)];
    } else if(enigma.current_reflector == 'C') {
        return enigma.reflector_C.reflector[enigma.convertLetter(letter)];
    }

    return '0';
}

func (enigma enigma) passThroughRotorReverse(letter rune, rotor int) rune {
    if(rotor == 1) {
        for i := 0; i < 26; i++ {
    		if(enigma.rotor_1.pins[i] == letter) {
                return enigma.convertNumber(i + 1);
            }
    	}
    } else if(rotor == 2) {
        for i := 0; i < 26; i++ {
            if(enigma.rotor_2.pins[i] == letter) {
                return enigma.convertNumber(i + 1);
            }
        }
    } else if(rotor == 3) {
        for i := 0; i < 26; i++ {
            if(enigma.rotor_3.pins[i] == letter) {
                return enigma.convertNumber(i + 1);
            }
        }
    }

    return '0';
}

func (enigma *enigma) incrementRotor(step int) {
    //TODO implement this
}

func (enigma enigma) Encrypt(letter rune) rune {
    enigma.incrementRotor(1);

    letter = unicode.ToUpper(letter);
    enigma.trace(letter);
    letter = enigma.passThroughPlugboard(letter);
    enigma.trace(letter);
    letter = enigma.passThroughRotor(letter, enigma.current_rotors[0]);
    enigma.trace(letter);
    letter = enigma.passThroughRotor(letter, enigma.current_rotors[1]);
    enigma.trace(letter);
    letter = enigma.passThroughRotor(letter, enigma.current_rotors[2]);
    enigma.trace(letter);
    letter = enigma.passThroughReflector(letter);
    enigma.trace(letter);
    letter = enigma.passThroughRotorReverse(letter, enigma.current_rotors[2]);
    enigma.trace(letter);
    letter = enigma.passThroughRotorReverse(letter, enigma.current_rotors[1]);
    enigma.trace(letter);
    letter = enigma.passThroughRotorReverse(letter, enigma.current_rotors[0]);
    enigma.trace(letter);
    letter = enigma.passThroughPlugboard(letter);
    enigma.trace(letter);

    return letter;
}

func Enigma() enigma {
    return enigma {
        //default settings
        current_rotors: [3]int{1,2,3},
        current_reflector: 'A',
        current_plugboard: 'A',
        trace_letters: false,
        //rotors
        rotor_1: rotor{
            [26]rune{'E','K','M','F','L','G','D','Q','V','Z','N','T','O','W','Y','H','X','U','S','P','A','I','B','R','C','J'},
            'Q',
            0,
        },
        rotor_2: rotor{
            [26]rune{'A','J','D','K','S','I','R','U','X','B','L','H','W','T','M','C','Q','G','Z','N','P','Y','F','V','O','E'},
            'E',
            0,
        },
        rotor_3: rotor{
            [26]rune{'B','D','F','H','J','L','C','P','R','T','X','V','Z','N','Y','E','I','W','G','A','K','M','U','S','Q','O'},
            'V',
            0,
        },
        //reflectors
        reflector_A: reflector{
            [26]rune{'E','J','M','Z','A','L','Y','X','V','B','W','F','C','R','Q','U','O','N','T','S','P','I','K','H','G','D'},
        },
        reflector_B: reflector{
            [26]rune{'Y','R','U','H','Q','S','L','D','P','X','N','G','O','K','M','I','E','B','F','Z','C','W','V','J','A','T'},
        },
        reflector_C: reflector{
            [26]rune{'F','V','P','J','I','A','O','Y','E','D','R','Z','X','W','G','C','T','K','U','Q','S','B','N','M','H','L'},
        },
        //plugboards
        plugboard_A: plugboard{
            [26]rune{'L','B','C','K','E','F','I','H','G','Z','D','A','O','X','M','S','W','Y','P','T','V','U','Q','N','R','J'},
        },
        plugboard_B: plugboard{
            [26]rune{'A','M','C','D','E','F','K','H','I','J','G','L','B','S','O','P','Q','R','N','T','Z','V','W','X','Y','U'},
        },
        plugboard_C: plugboard{
            [26]rune{'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'},
        },
    }
}
