package helloworldpb

import (
	"fmt"
	"strings"
)

func PrepareMessage(c *HelloWorldConfiguration) string {
	if h := c.GetHello(); h != nil {
		return prepareHello(h)
	}
	if b := c.GetBye(); b != nil {
		return prepareBye(b)
	}
	return "Not command specified: This should not happen due to validation that has to happen on consumer side."
}

func prepareHello(c *HelloCommand) string {
	switch c.Lang {
	case Lang_POLISH:
		msg := fmt.Sprintf("Witam %s świat dla %q w %v roku!", c.World, c.Name, c.Year)
		if c.PleaseAddReally || c.AddReally {
			return "(naprawdę!) " + msg
		}
		return msg
	case Lang_GERMAN:
		return "can't speak German really, contributions welcome ):"
	default:
		fallthrough
	case Lang_ENGLISH:
		msg := fmt.Sprintf("Hello %s world for %q in %v year!", c.World, c.Name, c.Year)
		if c.PleaseAddReally || c.AddReally {
			return "(really!) " + msg
		}
		return msg
	}
}

func prepareJustBye(l Lang) string {
	switch l {
	case Lang_POLISH:
		return "Pa pa świat!"
	case Lang_GERMAN:
		return "can't speak German really, contributions welcome ):"
	default:
		fallthrough
	case Lang_ENGLISH:
		return "See Ya world!"
	}
}

func prepareBye(c *ByeCommand) string {
	s := prepareJustBye(c.Lang)

	if b := c.GetConfigurable(); b != nil {
		for _, e := range b.Extra {
			s += fmt.Sprint(e)
		}
		// YOLO
		if b.Config.Configs[b.ConfigId].Capitalized {
			return strings.ToUpper(s)
		}
	}
	return s
}
