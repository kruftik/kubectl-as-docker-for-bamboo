package main

// docker cp \
//    /tmp/initialiseDockerContainer.sh8438742724428139147.tmp \
//    5fddebd3-e90f-4f9d-a7e9-8511730fbff271794852:/tmp/initialiseContainer.sh

type CopyCommand struct {
	Args struct {
		From   string `positional-arg-name:"SRC_PATH"`
		To  string `positional-arg-name:"CONTAINER:DEST_PATH"`
	} `positional-args:"yes" required:"yes"`
}

func (x *CopyCommand) Execute(args []string) error {
	//fmt.Printf("Adding (all=%v): %#v\n", x.All, args)
	return nil
}
