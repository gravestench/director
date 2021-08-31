package pkg

func (d *Director) initDirectorSystems() {
	d.AddSystem(&screenRenderingSystem{})
	d.AddSystem(&luaSystem{
		director: d,
	})
}