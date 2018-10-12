package pingaline

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// TableHealth format the health data to print
func TableHealth(h []Health) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n%s\t%s\t%s\t%s\t", "# Name", "# TYPE", "# STATUS", "# URL")

	for _, v := range h {
		fmt.Fprintf(w, "\n%v\t%v\t%v\t\"%v\"\t", v.Name, v.Type, v.Status, v.URL)
	}
}
