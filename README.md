# Golang Enigma Machine

## What is Enigma?
The Enigma machine is a cipher device developed and used in the early- to mid-20th century to protect commercial, diplomatic, and military communication. It was employed extensively by Nazi Germany during World War II in all branches of the German military. The Germans believed that use of the Enigma machine enabled them to communicate securely and thus giving them a huge advantage in World War II. [Demonstration Link](https://www.youtube.com/watch?v=-mdSvGUd0_c)

## Why Enigma?
Creating a digital Enigma machine employs many challenges which I've found helpful when learning new languages. It requires use of multiple data structures, types, and presents unique obstacles which can be solved one step at a time (like a puzzle).

## Program Notes:

### Setup
- Download the latest version of Go from [here](https://golang.org/dl/)
- Clone the `enigma-go` repository from [Github](https://github.com/stephenerdelyi/enigma-go)
- Inside the downloaded folder, run the project with `go run .`

### Using Enigma
Enigma requires the exact same settings to be used at the start of encryption to ensure a correct cipher/decipher process. Settings are printed before each menu. The following settings can be configured:

 - Plugboard (Options: A, B, C) - The current plugboard used when encrypting.
 - Rotor Order (Options: 1, 2, 3) - The order in which rotors are passed through during encryption. The first rotor is the right-most position, while the last rotor is the left-most position.
 - Rotor Position (Options: Any letter of the english alphabet) - The current position of a given rotor. The right-most (first) rotor increments by one letter each time a letter is encrypted. When the rotor hits a knockpoint, the next rotor will increment and so on.
 - Reflector (Options: A, B, C) - The current reflector used when encrypting.

Settings are printed like such: `P[A] R[1(C), 2(B), 3(A)] R[B]`

The above example shows an enigma using "plugboard A", with the right-most (first) rotor using "rotor 3" set in "position A", the middle (second) rotor using "rotor 2" set in "position B", the left-most (last) rotor using "rotor 1" set in "position C" and finally, "reflector B".

Once you have memorized your chosen settings, you can use the "Encrypt" option to enter a string of characters [A-Z] which, upon pressing enter, will be encrypted using the machine. You will notice the only settings which change as a result of encryption are the rotor positions. If you copy the output to your clipboard, reset your rotor positions to the same settings prior to encrypting, and enter the encrypted text back into the "Encrypt" option, you will get your original message in plaintext. You've just encrypted your first message using enigma!

### Using Saved Settings
Saved settings can speed up the process of switching between enigma settings when encrypting. These settings are non-volatile, meaning you can quit the program, reopen it, and they will still be there for use.

To add a new saved setting, do the following:
 - Set your enigma settings (plugboard, rotors, reflector) to your desired configuration
 - Select the "Saved Settings" menu
 - Select "Save Current Settings"

To use a previously saved setting, do the following:
 - Select the "Saved Settings" menu
 - Select "Switch to a Saved Setting"
 - Select the saved setting you'd like to use

Volia! You can now easily encrypt text and switch back to your previous settings for decryption or encrypting a new string of text.

### Future Considerations
No program is ever perfect! Here's some future-phase considerations to take this from good to great:
 - Implement error checking at various stages
 - Switch from use of local enigma package to GitHub-hosted package
 - Format function definitions with proper documentation (including @param and @return, etc).
 - All logic should be dynamic based on the hardware available. ie: allows infinite scalability of new rotors, plugboards and reflectors without the need to adjust getter/setter functions or the encryption process.
 - Move save settings functionality to its own package for a cleaner main.go file.
 - Use `getRotor` function when incrementing position in the `incrementRotor` function.
 - Display available options for hardware configurations in various settings menus (plugboard, rotor, reflector).
 - Allow the user to create their own plugboard pairs like a physical enigma would.
 - Display saved settings in a format which matches that of the "current settings" display line.
 - Refactor "saved settings" to be dynamic and not based on string position.
