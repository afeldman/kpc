package kpc

import (
	"fmt"
)

var KPCs map[string]KPC

func InitKPCs(basic_path []string) {
	KPCs = make(map[string]KPC)
}

func KPC_Size() int {
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
