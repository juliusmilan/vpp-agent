// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package ip

import "reflect"

var Types = map[string]reflect.Type{
	"FibMplsLabel":                                  reflect.TypeOf((*FibMplsLabel)(nil)).Elem(),
	"FibPath":                                       reflect.TypeOf((*FibPath)(nil)).Elem(),
	"IP4ArpEvent":                                   reflect.TypeOf((*IP4ArpEvent)(nil)).Elem(),
	"IP6FibDetails":                                 reflect.TypeOf((*IP6FibDetails)(nil)).Elem(),
	"IP6FibDump":                                    reflect.TypeOf((*IP6FibDump)(nil)).Elem(),
	"IP6MfibDetails":                                reflect.TypeOf((*IP6MfibDetails)(nil)).Elem(),
	"IP6MfibDump":                                   reflect.TypeOf((*IP6MfibDump)(nil)).Elem(),
	"IP6NdEvent":                                    reflect.TypeOf((*IP6NdEvent)(nil)).Elem(),
	"IP6RaEvent":                                    reflect.TypeOf((*IP6RaEvent)(nil)).Elem(),
	"IP6RaPrefixInfo":                               reflect.TypeOf((*IP6RaPrefixInfo)(nil)).Elem(),
	"IP6ndProxyAddDel":                              reflect.TypeOf((*IP6ndProxyAddDel)(nil)).Elem(),
	"IP6ndProxyAddDelReply":                         reflect.TypeOf((*IP6ndProxyAddDelReply)(nil)).Elem(),
	"IP6ndProxyDetails":                             reflect.TypeOf((*IP6ndProxyDetails)(nil)).Elem(),
	"IP6ndProxyDump":                                reflect.TypeOf((*IP6ndProxyDump)(nil)).Elem(),
	"IP6ndSendRouterSolicitation":                   reflect.TypeOf((*IP6ndSendRouterSolicitation)(nil)).Elem(),
	"IP6ndSendRouterSolicitationReply":              reflect.TypeOf((*IP6ndSendRouterSolicitationReply)(nil)).Elem(),
	"IPAddDelRoute":                                 reflect.TypeOf((*IPAddDelRoute)(nil)).Elem(),
	"IPAddDelRouteReply":                            reflect.TypeOf((*IPAddDelRouteReply)(nil)).Elem(),
	"IPAddressDetails":                              reflect.TypeOf((*IPAddressDetails)(nil)).Elem(),
	"IPAddressDump":                                 reflect.TypeOf((*IPAddressDump)(nil)).Elem(),
	"IPContainerProxyAddDel":                        reflect.TypeOf((*IPContainerProxyAddDel)(nil)).Elem(),
	"IPContainerProxyAddDelReply":                   reflect.TypeOf((*IPContainerProxyAddDelReply)(nil)).Elem(),
	"IPDetails":                                     reflect.TypeOf((*IPDetails)(nil)).Elem(),
	"IPDump":                                        reflect.TypeOf((*IPDump)(nil)).Elem(),
	"IPFibDetails":                                  reflect.TypeOf((*IPFibDetails)(nil)).Elem(),
	"IPFibDump":                                     reflect.TypeOf((*IPFibDump)(nil)).Elem(),
	"IPMfibDetails":                                 reflect.TypeOf((*IPMfibDetails)(nil)).Elem(),
	"IPMfibDump":                                    reflect.TypeOf((*IPMfibDump)(nil)).Elem(),
	"IPMrouteAddDel":                                reflect.TypeOf((*IPMrouteAddDel)(nil)).Elem(),
	"IPMrouteAddDelReply":                           reflect.TypeOf((*IPMrouteAddDelReply)(nil)).Elem(),
	"IPNeighborAddDel":                              reflect.TypeOf((*IPNeighborAddDel)(nil)).Elem(),
	"IPNeighborAddDelReply":                         reflect.TypeOf((*IPNeighborAddDelReply)(nil)).Elem(),
	"IPNeighborDetails":                             reflect.TypeOf((*IPNeighborDetails)(nil)).Elem(),
	"IPNeighborDump":                                reflect.TypeOf((*IPNeighborDump)(nil)).Elem(),
	"IPPuntPolice":                                  reflect.TypeOf((*IPPuntPolice)(nil)).Elem(),
	"IPPuntPoliceReply":                             reflect.TypeOf((*IPPuntPoliceReply)(nil)).Elem(),
	"IPPuntRedirect":                                reflect.TypeOf((*IPPuntRedirect)(nil)).Elem(),
	"IPPuntRedirectReply":                           reflect.TypeOf((*IPPuntRedirectReply)(nil)).Elem(),
	"IPReassemblyEnableDisable":                     reflect.TypeOf((*IPReassemblyEnableDisable)(nil)).Elem(),
	"IPReassemblyEnableDisableReply":                reflect.TypeOf((*IPReassemblyEnableDisableReply)(nil)).Elem(),
	"IPReassemblyGet":                               reflect.TypeOf((*IPReassemblyGet)(nil)).Elem(),
	"IPReassemblyGetReply":                          reflect.TypeOf((*IPReassemblyGetReply)(nil)).Elem(),
	"IPReassemblySet":                               reflect.TypeOf((*IPReassemblySet)(nil)).Elem(),
	"IPReassemblySetReply":                          reflect.TypeOf((*IPReassemblySetReply)(nil)).Elem(),
	"IPSourceAndPortRangeCheckAddDel":               reflect.TypeOf((*IPSourceAndPortRangeCheckAddDel)(nil)).Elem(),
	"IPSourceAndPortRangeCheckAddDelReply":          reflect.TypeOf((*IPSourceAndPortRangeCheckAddDelReply)(nil)).Elem(),
	"IPSourceAndPortRangeCheckInterfaceAddDel":      reflect.TypeOf((*IPSourceAndPortRangeCheckInterfaceAddDel)(nil)).Elem(),
	"IPSourceAndPortRangeCheckInterfaceAddDelReply": reflect.TypeOf((*IPSourceAndPortRangeCheckInterfaceAddDelReply)(nil)).Elem(),
	"IPTableAddDel":                                 reflect.TypeOf((*IPTableAddDel)(nil)).Elem(),
	"IPTableAddDelReply":                            reflect.TypeOf((*IPTableAddDelReply)(nil)).Elem(),
	"IoamDisable":                                   reflect.TypeOf((*IoamDisable)(nil)).Elem(),
	"IoamDisableReply":                              reflect.TypeOf((*IoamDisableReply)(nil)).Elem(),
	"IoamEnable":                                    reflect.TypeOf((*IoamEnable)(nil)).Elem(),
	"IoamEnableReply":                               reflect.TypeOf((*IoamEnableReply)(nil)).Elem(),
	"MfibSignalDetails":                             reflect.TypeOf((*MfibSignalDetails)(nil)).Elem(),
	"MfibSignalDump":                                reflect.TypeOf((*MfibSignalDump)(nil)).Elem(),
	"ProxyArpAddDel":                                reflect.TypeOf((*ProxyArpAddDel)(nil)).Elem(),
	"ProxyArpAddDelReply":                           reflect.TypeOf((*ProxyArpAddDelReply)(nil)).Elem(),
	"ProxyArpIntfcEnableDisable":                    reflect.TypeOf((*ProxyArpIntfcEnableDisable)(nil)).Elem(),
	"ProxyArpIntfcEnableDisableReply":               reflect.TypeOf((*ProxyArpIntfcEnableDisableReply)(nil)).Elem(),
	"ResetFib":                                      reflect.TypeOf((*ResetFib)(nil)).Elem(),
	"ResetFibReply":                                 reflect.TypeOf((*ResetFibReply)(nil)).Elem(),
	"SetArpNeighborLimit":                           reflect.TypeOf((*SetArpNeighborLimit)(nil)).Elem(),
	"SetArpNeighborLimitReply":                      reflect.TypeOf((*SetArpNeighborLimitReply)(nil)).Elem(),
	"SetIPFlowHash":                                 reflect.TypeOf((*SetIPFlowHash)(nil)).Elem(),
	"SetIPFlowHashReply":                            reflect.TypeOf((*SetIPFlowHashReply)(nil)).Elem(),
	"SwInterfaceIP6EnableDisable":                   reflect.TypeOf((*SwInterfaceIP6EnableDisable)(nil)).Elem(),
	"SwInterfaceIP6EnableDisableReply":              reflect.TypeOf((*SwInterfaceIP6EnableDisableReply)(nil)).Elem(),
	"SwInterfaceIP6SetLinkLocalAddress":             reflect.TypeOf((*SwInterfaceIP6SetLinkLocalAddress)(nil)).Elem(),
	"SwInterfaceIP6SetLinkLocalAddressReply":        reflect.TypeOf((*SwInterfaceIP6SetLinkLocalAddressReply)(nil)).Elem(),
	"SwInterfaceIP6ndRaConfig":                      reflect.TypeOf((*SwInterfaceIP6ndRaConfig)(nil)).Elem(),
	"SwInterfaceIP6ndRaConfigReply":                 reflect.TypeOf((*SwInterfaceIP6ndRaConfigReply)(nil)).Elem(),
	"SwInterfaceIP6ndRaPrefix":                      reflect.TypeOf((*SwInterfaceIP6ndRaPrefix)(nil)).Elem(),
	"SwInterfaceIP6ndRaPrefixReply":                 reflect.TypeOf((*SwInterfaceIP6ndRaPrefixReply)(nil)).Elem(),
	"WantIP4ArpEvents":                              reflect.TypeOf((*WantIP4ArpEvents)(nil)).Elem(),
	"WantIP4ArpEventsReply":                         reflect.TypeOf((*WantIP4ArpEventsReply)(nil)).Elem(),
	"WantIP6NdEvents":                               reflect.TypeOf((*WantIP6NdEvents)(nil)).Elem(),
	"WantIP6NdEventsReply":                          reflect.TypeOf((*WantIP6NdEventsReply)(nil)).Elem(),
	"WantIP6RaEvents":                               reflect.TypeOf((*WantIP6RaEvents)(nil)).Elem(),
	"WantIP6RaEventsReply":                          reflect.TypeOf((*WantIP6RaEventsReply)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewIP4ArpEvent":                                   reflect.ValueOf(NewIP4ArpEvent),
	"NewIP6FibDetails":                                 reflect.ValueOf(NewIP6FibDetails),
	"NewIP6FibDump":                                    reflect.ValueOf(NewIP6FibDump),
	"NewIP6MfibDetails":                                reflect.ValueOf(NewIP6MfibDetails),
	"NewIP6MfibDump":                                   reflect.ValueOf(NewIP6MfibDump),
	"NewIP6NdEvent":                                    reflect.ValueOf(NewIP6NdEvent),
	"NewIP6RaEvent":                                    reflect.ValueOf(NewIP6RaEvent),
	"NewIP6ndProxyAddDel":                              reflect.ValueOf(NewIP6ndProxyAddDel),
	"NewIP6ndProxyAddDelReply":                         reflect.ValueOf(NewIP6ndProxyAddDelReply),
	"NewIP6ndProxyDetails":                             reflect.ValueOf(NewIP6ndProxyDetails),
	"NewIP6ndProxyDump":                                reflect.ValueOf(NewIP6ndProxyDump),
	"NewIP6ndSendRouterSolicitation":                   reflect.ValueOf(NewIP6ndSendRouterSolicitation),
	"NewIP6ndSendRouterSolicitationReply":              reflect.ValueOf(NewIP6ndSendRouterSolicitationReply),
	"NewIPAddDelRoute":                                 reflect.ValueOf(NewIPAddDelRoute),
	"NewIPAddDelRouteReply":                            reflect.ValueOf(NewIPAddDelRouteReply),
	"NewIPAddressDetails":                              reflect.ValueOf(NewIPAddressDetails),
	"NewIPAddressDump":                                 reflect.ValueOf(NewIPAddressDump),
	"NewIPContainerProxyAddDel":                        reflect.ValueOf(NewIPContainerProxyAddDel),
	"NewIPContainerProxyAddDelReply":                   reflect.ValueOf(NewIPContainerProxyAddDelReply),
	"NewIPDetails":                                     reflect.ValueOf(NewIPDetails),
	"NewIPDump":                                        reflect.ValueOf(NewIPDump),
	"NewIPFibDetails":                                  reflect.ValueOf(NewIPFibDetails),
	"NewIPFibDump":                                     reflect.ValueOf(NewIPFibDump),
	"NewIPMfibDetails":                                 reflect.ValueOf(NewIPMfibDetails),
	"NewIPMfibDump":                                    reflect.ValueOf(NewIPMfibDump),
	"NewIPMrouteAddDel":                                reflect.ValueOf(NewIPMrouteAddDel),
	"NewIPMrouteAddDelReply":                           reflect.ValueOf(NewIPMrouteAddDelReply),
	"NewIPNeighborAddDel":                              reflect.ValueOf(NewIPNeighborAddDel),
	"NewIPNeighborAddDelReply":                         reflect.ValueOf(NewIPNeighborAddDelReply),
	"NewIPNeighborDetails":                             reflect.ValueOf(NewIPNeighborDetails),
	"NewIPNeighborDump":                                reflect.ValueOf(NewIPNeighborDump),
	"NewIPPuntPolice":                                  reflect.ValueOf(NewIPPuntPolice),
	"NewIPPuntPoliceReply":                             reflect.ValueOf(NewIPPuntPoliceReply),
	"NewIPPuntRedirect":                                reflect.ValueOf(NewIPPuntRedirect),
	"NewIPPuntRedirectReply":                           reflect.ValueOf(NewIPPuntRedirectReply),
	"NewIPReassemblyEnableDisable":                     reflect.ValueOf(NewIPReassemblyEnableDisable),
	"NewIPReassemblyEnableDisableReply":                reflect.ValueOf(NewIPReassemblyEnableDisableReply),
	"NewIPReassemblyGet":                               reflect.ValueOf(NewIPReassemblyGet),
	"NewIPReassemblyGetReply":                          reflect.ValueOf(NewIPReassemblyGetReply),
	"NewIPReassemblySet":                               reflect.ValueOf(NewIPReassemblySet),
	"NewIPReassemblySetReply":                          reflect.ValueOf(NewIPReassemblySetReply),
	"NewIPSourceAndPortRangeCheckAddDel":               reflect.ValueOf(NewIPSourceAndPortRangeCheckAddDel),
	"NewIPSourceAndPortRangeCheckAddDelReply":          reflect.ValueOf(NewIPSourceAndPortRangeCheckAddDelReply),
	"NewIPSourceAndPortRangeCheckInterfaceAddDel":      reflect.ValueOf(NewIPSourceAndPortRangeCheckInterfaceAddDel),
	"NewIPSourceAndPortRangeCheckInterfaceAddDelReply": reflect.ValueOf(NewIPSourceAndPortRangeCheckInterfaceAddDelReply),
	"NewIPTableAddDel":                                 reflect.ValueOf(NewIPTableAddDel),
	"NewIPTableAddDelReply":                            reflect.ValueOf(NewIPTableAddDelReply),
	"NewIoamDisable":                                   reflect.ValueOf(NewIoamDisable),
	"NewIoamDisableReply":                              reflect.ValueOf(NewIoamDisableReply),
	"NewIoamEnable":                                    reflect.ValueOf(NewIoamEnable),
	"NewIoamEnableReply":                               reflect.ValueOf(NewIoamEnableReply),
	"NewMfibSignalDetails":                             reflect.ValueOf(NewMfibSignalDetails),
	"NewMfibSignalDump":                                reflect.ValueOf(NewMfibSignalDump),
	"NewProxyArpAddDel":                                reflect.ValueOf(NewProxyArpAddDel),
	"NewProxyArpAddDelReply":                           reflect.ValueOf(NewProxyArpAddDelReply),
	"NewProxyArpIntfcEnableDisable":                    reflect.ValueOf(NewProxyArpIntfcEnableDisable),
	"NewProxyArpIntfcEnableDisableReply":               reflect.ValueOf(NewProxyArpIntfcEnableDisableReply),
	"NewResetFib":                                      reflect.ValueOf(NewResetFib),
	"NewResetFibReply":                                 reflect.ValueOf(NewResetFibReply),
	"NewSetArpNeighborLimit":                           reflect.ValueOf(NewSetArpNeighborLimit),
	"NewSetArpNeighborLimitReply":                      reflect.ValueOf(NewSetArpNeighborLimitReply),
	"NewSetIPFlowHash":                                 reflect.ValueOf(NewSetIPFlowHash),
	"NewSetIPFlowHashReply":                            reflect.ValueOf(NewSetIPFlowHashReply),
	"NewSwInterfaceIP6EnableDisable":                   reflect.ValueOf(NewSwInterfaceIP6EnableDisable),
	"NewSwInterfaceIP6EnableDisableReply":              reflect.ValueOf(NewSwInterfaceIP6EnableDisableReply),
	"NewSwInterfaceIP6SetLinkLocalAddress":             reflect.ValueOf(NewSwInterfaceIP6SetLinkLocalAddress),
	"NewSwInterfaceIP6SetLinkLocalAddressReply":        reflect.ValueOf(NewSwInterfaceIP6SetLinkLocalAddressReply),
	"NewSwInterfaceIP6ndRaConfig":                      reflect.ValueOf(NewSwInterfaceIP6ndRaConfig),
	"NewSwInterfaceIP6ndRaConfigReply":                 reflect.ValueOf(NewSwInterfaceIP6ndRaConfigReply),
	"NewSwInterfaceIP6ndRaPrefix":                      reflect.ValueOf(NewSwInterfaceIP6ndRaPrefix),
	"NewSwInterfaceIP6ndRaPrefixReply":                 reflect.ValueOf(NewSwInterfaceIP6ndRaPrefixReply),
	"NewWantIP4ArpEvents":                              reflect.ValueOf(NewWantIP4ArpEvents),
	"NewWantIP4ArpEventsReply":                         reflect.ValueOf(NewWantIP4ArpEventsReply),
	"NewWantIP6NdEvents":                               reflect.ValueOf(NewWantIP6NdEvents),
	"NewWantIP6NdEventsReply":                          reflect.ValueOf(NewWantIP6NdEventsReply),
	"NewWantIP6RaEvents":                               reflect.ValueOf(NewWantIP6RaEvents),
	"NewWantIP6RaEventsReply":                          reflect.ValueOf(NewWantIP6RaEventsReply),
}

var Variables = map[string]reflect.Value{}

var Consts = map[string]reflect.Value{}
