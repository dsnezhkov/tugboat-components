package travel2

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dsnezhkov/tugboat/defs"
)

func (component *ComponentTravel2) invoke(msg defs.Message, handoffTo []string) bool {

	if component.Options.Directive != "" {

		switch component.Options.Directive {
		case "duration":
		default:
			fmt.Printf("Component::Travel2: Unknown command directive: %s.\n", component.Options.Directive)
		}

		if component.Options.DirectiveOpts != nil {
			durationStr := component.Options.DirectiveOpts[0]
			durationInt, err := strconv.ParseInt(durationStr, 10, 64)
			if err != nil {
				log.Printf("Component::Travel2: Error in conversion of duration: %v", err)
			}
			log.Printf("Component::Travel2: \n\nSleeping .... ")
			time.Sleep(time.Duration(durationInt * int64(time.Millisecond)))
		}
	}
	log.Printf("Component::Travel2: Not handing off")

	return true
}
