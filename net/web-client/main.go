package main

func main() {

	cli := Setup(true, false)

	if err := DefaultGetGolang(); err != nil {
		panic(err)
	}

	if err := DoOps(cli); err != nil {
		panic(err)
	}

	c := Controller{Client: cli}
	if err := c.DoOps(); err != nil {
		panic(err)
	}

	Setup(true, true)
	if err := DefaultGetGolang(); err != nil {
		panic(err)
	}

}
