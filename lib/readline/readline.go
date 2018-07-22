package readline

import (
   "bufio"
   "os"
)

// (string -> IO) -> error
func On(callback func(string)) error {
    stdinScanner := bufio.NewScanner(os.Stdin)
    for stdinScanner.Scan() {
        text := stdinScanner.Text()
        callback(text)
    }

    return stdinScanner.Err()
}
