package main

//import "fmt"

// docker exec \
//    -u root \
//    5fddebd3-e90f-4f9d-a7e9-8511730fbff271794852 \
//        chown root:root /tmp/initialiseContainer.sh

type ExecCommand struct {
	Interactive bool `short:"i" long:"interactive" description:"Keep STDIN open even if not attached"`
	Tty bool `short:"t" long:"tty" description:"Allocate a pseudo-TTY"`
	User string `short:"u" long:"user" description:"Username or UID (format: <name|uid>[:<group|gid>])"`
	Args struct {
		CID   string `positional-arg-name:"CONTAINER"`
		Rest  []string `positional-arg-name:"COMMAND [ARG...]" required:"1"`
	} `positional-args:"yes" required:"yes"`

}

func (x *ExecCommand) Execute(args []string) error {
	//fmt.Printf("Adding (all=%v): %#v\n", x.All, args)
	return nil
}