package main

type Cell struct {
	alive bool
}

func (c *Cell) IsAlive() bool {
	return c.alive
}

func (c *Cell) SetAlive(alive bool) {
	c.alive = alive
}
