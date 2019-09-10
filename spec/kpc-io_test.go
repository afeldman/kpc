package spec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/afeldman/kpc"
)

const text = `{"kpc_version":"0.1.0","name":"Test-Project","version":"","requirements":[{"name":"test","version":"0.1.1"}],"conflicts":[{"name":"test","version":["1.12.31"]}],"authors":[],"source":{"url":"github","tag":"aaa123"},"prefix":"","main":""}`

var _ = Describe("KPC", func() {
	var (
		tkpc KPC
	)

	Context("check the KPC io", func() {
		It("get json", func() {
			tkpc = *InitKPC("Test-Project")
			con := InitConflict("test", "1.12.31")
			requirement := InitRequirement("test")
			requirement.SetVersion("0.1.1")

			tkpc.GetRepo().SetTag("aaa123")
			tkpc.GetRepo().SetURL("github")

			tkpc.AddConflict(*con)
			tkpc.AddRequirement(*requirement)
			_, json_s := tkpc.ToJSON()
			Expect(string(json_s)).Should(Equal(text))
		})
		It("from json", func() {
			_, tkpc := FromJSON([]byte(text))
			Expect(*(tkpc.GetName())).Should(Equal("Test-Project"))
			Expect(*(tkpc.GetVersion())).Should(Equal(""))
			Expect(*(tkpc.GetRepo().GetTag())).Should(Equal("aaa123"))
		})
	})
})
