package main

import (
	"context"
)

type Foo struct{}

func (m *Foo) Desc(ctx context.Context) (string, error) {
	descGo, _ := dag.Git("https://github.com/jpadams/testgo.git").Branch("main").Tree().AsModule().Initialize().Description(ctx) 
	descTs, _ := dag.Git("https://github.com/jpadams/testts.git").Branch("main").Tree().AsModule().Initialize().Description(ctx) 
	descPy, _ := dag.Git("https://github.com/jpadams/testpy.git").Branch("main").Tree().AsModule().Initialize().Description(ctx) 
	return dag.Container().
		From("alpine:latest").
		WithWorkdir("src").
		WithNewFile("go", ContainerWithNewFileOpts{Contents: descGo}).
		WithNewFile("ts", ContainerWithNewFileOpts{Contents: descTs}).
		WithNewFile("py", ContainerWithNewFileOpts{Contents: descPy}).
		WithNewFile("xx", ContainerWithNewFileOpts{Contents: " xx "}). // sub xx in to see fail
		WithExec([]string{"sh", "-c", "diff go ts && diff ts py"}).
		Stdout(ctx)		
}
