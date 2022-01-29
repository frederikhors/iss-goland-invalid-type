package main

import (
	"go/types"

	"golang.org/x/tools/go/packages"
)

func main() {
	playerModel := LoadPackage("./player.go")

	var playerStruct *types.Struct

	for _, entity := range playerModel.Types.Scope().Names() {
		if entity == "Player" {

			playerStruct = GetStruct(entity, playerModel)

			break
		}
	}

	for i := 0; i < playerStruct.NumFields(); i++ {
		field := playerStruct.Field(i)

		println("field.Name: " + field.Name() + " - field.Type: " + field.Type().String())
	}
}

func LoadPackage(path string) *packages.Package {
	const mode = packages.NeedTypes |
		packages.NeedName |
		packages.NeedSyntax |
		packages.NeedFiles |
		packages.NeedTypesInfo |
		packages.NeedTypesInfo |
		packages.NeedModule

	cfg := &packages.Config{Mode: mode}

	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		panic(err)
	}

	return pkgs[0]
}

func GetStruct(structName string, pkg *packages.Package) *types.Struct {
	foundStruct := pkg.Types.Scope().Lookup(structName)

	if foundStruct == nil {
		return nil
	}

	res, _ := foundStruct.Type().Underlying().(*types.Struct)

	return res
}
