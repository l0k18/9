package config

func NewApp(name string, g ...AppGenerator) (out *App) {
	gen := AppGenerators(g)
	out = &App{
		Name: name,
		Cats: make(Cats),
	}
	gen.RunAll(out)
	return
}

// which is made from

func Version(ver string) AppGenerator {
	return func(ctx *App) {
		ctx.Version = func() string {
			return ver
		}
	}
}

func Group(name string, g ...CatGenerator) AppGenerator {
	G := CatGenerators(g)
	return func(ctx *App) {
		ctx.Cats[name] = make(Cat)
		G.RunAll(ctx.Cats[name])
	}
}

// which is made from

// TODO: these need to attach validator/accessors

func File(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.File
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Dir(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.Dir
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Port(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.Port
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func boolRow(name string, enabled bool, g RowGenerators) CatGenerator {
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.Bool
			cc.Value = &enabled
			g.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Enable(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return boolRow(name, false, G)
}

func Enabled(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return boolRow(name, true, G)
}

func Int(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.Int
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Tag(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.Tag
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Tags(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.Tags
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Addr(name string, defPort int, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = GenAddr(name, defPort)
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Addrs(name string, defPort int, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = GenAddrs(name, defPort)
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Level(g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	const lvl = "level"
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = lvl
			cc.Validate = Valid.Level
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[lvl] = *c
	}
}

func Algo(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.Algo
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Float(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.Float
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Duration(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.Duration
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

func Net(name string, g ...RowGenerator) CatGenerator {
	G := RowGenerators(g)
	return func(ctx *Cat) {
		c := &Row{}
		c.Init = func(cc *Row) {
			cc.Name = name
			cc.Validate = Valid.Net
			G.RunAll(cc)
		}
		c.Init(c)
		(*ctx)[name] = *c
	}
}

// which is populated by

// Usage populates the usage field for information about a config item
func Usage(usage string) RowGenerator {
	return func(ctx *Row) {
		ctx.Usage = usage
	}
}

// Default sets the default value for a config item
func Default(in interface{}) RowGenerator {
	return func(ctx *Row) {
		ctx.Validate(ctx, in)
	}
}

// Min attaches to the validator a test that enforces a minimum
func Min(min int) RowGenerator {
	return func(ctx *Row) {
		v := ctx.Validate
		ctx.Validate = func(r *Row, in interface{}) bool {
			n := min
			switch I := in.(type) {
			case int:
				n = I
			case *int:
				n = *I
			}
			if n < min {
				in = min
			}
			// none of the above will affect if this wasn't an int
			return v(r, in)
		}
	}
}

// Max attaches to the validator a test that enforces a maximum
func Max(max int) RowGenerator {
	return func(ctx *Row) {
		v := ctx.Validate
		ctx.Validate = func(r *Row, in interface{}) bool {
			n := max
			switch I := in.(type) {
			case int:
				n = I
			case *int:
				n = *I
			}
			if n > max {
				in = max
			}
			// none of the above will affect if this wasn't an int
			return v(r, in)
		}
	}
}

// RandomsString generates a random number and converts to base32 for
// a default random password of some number of characters
func RandomString(in int) RowGenerator {
	return func(ctx *Row) {

	}
}
