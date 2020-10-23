// automatically generated by stateify.

package control

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (fs *RightsFiles) StateTypeName() string {
	return "pkg/sentry/socket/control.RightsFiles"
}

func (fs *RightsFiles) StateFields() []string {
	return nil
}

func (c *scmCredentials) StateTypeName() string {
	return "pkg/sentry/socket/control.scmCredentials"
}

func (c *scmCredentials) StateFields() []string {
	return []string{
		"t",
		"kuid",
		"kgid",
	}
}

func (c *scmCredentials) beforeSave() {}

func (c *scmCredentials) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.t)
	stateSinkObject.Save(1, &c.kuid)
	stateSinkObject.Save(2, &c.kgid)
}

func (c *scmCredentials) afterLoad() {}

func (c *scmCredentials) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.t)
	stateSourceObject.Load(1, &c.kuid)
	stateSourceObject.Load(2, &c.kgid)
}

func init() {
	state.Register((*RightsFiles)(nil))
	state.Register((*scmCredentials)(nil))
}