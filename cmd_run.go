package main

//	docker run \
//	--volume /APP/bamboo_agents/agent-1/xml-data/build-dir/PHSRVS-JSPLNR-JOB2:/APP/bamboo_agents/agent-1/xml-data/build-dir/PHSRVS-JSPLNR-JOB2 \
//	--volume /APP/bamboo_agents/agent-1/temp:/APP/bamboo_agents/agent-1/temp \
//	--detach \
//	--name 5fddebd3-e90f-4f9d-a7e9-8511730fbff271794852 \
//	--net=host \
//	docker-hub.binary.alfabank.ru/alpine:latest \
//		tail -f /dev/null

type RunCommand struct {
	Volume bool `short:"v" long:"volume" description:"Bind mount a volume"`
	Detach bool `short:"d" long:"detach" description:"Run container in background and print container ID"`
	Name bool `short:"n" long:"name" description:"Assign a name to the container"`
	Network bool `long:"network" alias:"net" description:"Connect a container to a network"`
	Args struct {
		ContainerImage   string `positional-arg-name:"IMAGE"`
		Rest  []string `positional-arg-name:"COMMAND [ARG...]" required:"1"`
	} `positional-args:"yes" required:"yes"`
}

func (x *RunCommand) Execute(args []string) error {
	//fmt.Printf("Adding (all=%v): %#v\n", x.All, args)
	return nil
}
