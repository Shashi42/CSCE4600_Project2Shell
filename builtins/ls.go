package builtins

import (
	"fmt"
	"os"
	"log"
)

func ListDirectory(args ...string) error {
	entries, err := os.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
 
    for _, e := range entries {
            fmt.Println(e.Name())
    }
	// switch len(args) {
	// case 0: // change to home directory if available
	// 	if HomeDir == "" {
	// 		return fmt.Errorf("%w: no homedir found, expected one argument (directory)", ErrInvalidArgCount)
	// 	}
	// 	return os.Chdir(HomeDir)
	// case 1:
	// 	return os.Chdir(args[0])
	// default:
	// 	return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidArgCount)
	// }
	return err
}
