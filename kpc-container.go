package kpc

import (
	"fmt"
	"strings"

	fs "github.com/afeldman/go-util/fs"
)

var KPCs map[string]KPC

func InitKPCs(kpc_root string) error {

	KPCs = make(map[string]KPC)

	files, err := fs.FilePathWalkDir(kpc_root)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.HasSuffix(strings.ToLower(file), "kpc") {
			err, kpc := ReadKPCFile(file)
			if err != nil {
				return err
			}
			KPCs[*(kpc.GetName())] = *kpc
		}
	}
	return nil
}

func KPCs_Size() int {
	return len(KPCs)
}

func All() string {
	var names string

	for kpc_key, _ := range KPCs {
		names += fmt.Sprintf("%s\n", kpc_key)
	}

	return names
}

func ContainsPackage(comp string) bool {
	if _, ok := KPCs[comp]; ok {
		return true
	}
	return false
}

func GetKPC(kpg_name string) *KPC {
	if ContainsPackage(kpg_name) {
		kpc := KPCs[kpg_name]
		return (&kpc)
	} else {
		return nil
	}
}
