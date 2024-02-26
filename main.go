package main

import (
	"context"
)

type Foo struct{}

func (m *Foo) Desc(ctx context.Context) string {
	descGo, _ := dag.Git("https://github.com/jpadams/testgo.git").Branch("main").Tree().AsModule().Initialize().Description(ctx) 
	descTs, _ := dag.Git("https://github.com/jpadams/testts.git").Branch("main").Tree().AsModule().Initialize().Description(ctx) 
	descPy, _ := dag.Git("https://github.com/jpadams/testpy.git").Branch("main").Tree().AsModule().Initialize().Description(ctx) 
	return descGo+"\n============\n"+descTs+"\n============\n"+descPy
}
