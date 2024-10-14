package main



type Server struct {
	// 	accountClient *account.Client
	// 	catalogClient *catalog.Client
	// 	orderClient *order.Client
}

// func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
// 	accountClient, err := account.NewClient(accountUrl)
// 	if err != nil {
// 		return nil, err
// 	}

// 	catalogClient, err := catalog.NewClient(catalogUrl)
// 	accountClient.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	orderClient, err := order.NewClient(orderUrl)
// 	accountClient.Close()
// 	catalogClient.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Server{
// 		accountClient,
// 		catalogClient,
// 		orderClient,
// 	}, nil
// }

// func (S *Server) Mutation() MutationResolver {
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() QueryResolver {
// 	return &queryResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() AccountResolver {
// 	return &accountResolver{
// 		server: s,
// 	}
// }

// func (s *Server) ToExecutableSchema() graphql.ExecutableSchema{
// 	return NewExecutableSchema(Config{
// 		Resolvers: s,
// 	})
// }
