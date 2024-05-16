package resolvers

type Root struct {
	*Chains
	*MessagesResolver
	*StatsResolver
}

type Provider interface {
	ChainsProvider
	MessagesProvider
	StatsProvider
}

func NewRoot(p Provider) *Root {
	return &Root{
		Chains:           NewChainsProvider(p),
		MessagesResolver: NewMessagesResolver(p),
		StatsResolver:    &StatsResolver{Provider: p},
	}
}
