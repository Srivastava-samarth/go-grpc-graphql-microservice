package order

import (
	"context"
	"errors"
	"fmt"
	account "go-grpc-graphql-microservice/account"
	catalog "go-grpc-graphql-microservice/catalog"
	"go-grpc-graphql-microservice/order/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service       Service
	accountClient *account.Client
	catalogClient *catalog.Client
	pb.UnimplementedOrderServiceServer
}

func ListenGRPC(s Service, accountURL, catalogURL string, port int) error {
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterOrderServiceServer(serv, &grpcServer{
		UnimplementedOrderServiceServer: pb.UnimplementedOrderServiceServer{},
		service: s,
		accountClient: accountClient,
		catalogClient:catalogClient,
	})
	reflection.Register(serv)

	return serv.Serve(lis)
}

func (s *grpcServer) PostOrder(
	ctx context.Context,
	r *pb.PostOrderRequest,
) (*pb.PostOrderResponse, error) {
	_,err := s.accountClient.GetAccount(ctx,r.AccountId)
	if err!=nil{
		log.Println("Error getting account: ",err)
		return nil,errors.New("account not found")
	}
	//GET Ordered Products
	productIDs := []string{}
	for _,p := range r.Products{
		productIDs = append(productIDs, p.ProductId)
	}
	orderedProducts,err := s.catalogClient.GetProducts(ctx,0,0,productIDs,"")
	if err!=nil{
		log.Println("Error getting products: ",err)
		return nil,errors.New("products not found")
	}
	//construct products
	products := []OrderedProduct{}
	for _,p := range orderedProducts{
		product := OrderedProduct{
			ID:          p.ID,
			Quantity:    0,
			Price:       p.Price,
			Name:        p.Name,
			Description: p.Description,
		}
		for _,rp := range r.Products{
			if rp.ProductId == p.ID{
				product.Quantity = rp.Quantity
				break;
			}
		}
		if product.Quantity != 0{
			products = append(products,product)
		}
	}
	//call service implementation
	order,err := s.service.PostOrder(ctx,r.AccountId,products)
	if err!=nil{
		log.Println("Error posting order: ",err)
		return nil,errors.New("could not pots order")
	}
	//make response order
	orderProto := &pb.Order{
		Id:         order.ID,
		AccountId:  order.AccountID,
		TotalPrice: order.TotalPrice,
		Products:   []*pb.Order_OrderProduct{},
	}
	orderProto.CreatedAt,_ = order.CreatedAt.MarshalBinary()
	for _,p := range order.Products{
		orderProto.Products = append(orderProto.Products, &pb.Order_OrderProduct{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    p.Quantity,
		})
	}
	return &pb.PostOrderResponse{
		Order:orderProto,
	},nil
}

func (s *grpcServer) GetOrdersForAccount(
	ctx context.Context,
	r *pb.GetOrdersForAccountRequest,
) (*pb.GetOrdersForAccountResponse, error){
	//get orders for account
	accountOrders,err := s.service.GetOrdersForAccount(ctx,r.AccountId)
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	//get all ordered products
	productIDMap := map[string]bool{}
	for _,o := range accountOrders{
		for _,p := range o.Products{
			productIDMap[p.ID] = true
		}
	}
	productIDs := []string{}
	for id := range productIDMap{
		productIDs = append(productIDs, id)
	}
	products,err := s.catalogClient.GetProducts(ctx,0,0,productIDs,"")
	if err!=nil{
		log.Println("Error getting account products: ",err)
		return nil,err
	}
	//construct orders
	orders := []*pb.Order{}
	for _,o := range accountOrders{
		//encode order
		op := &pb.Order{
			AccountId:  o.AccountID,
			Id:         o.ID,
			TotalPrice: o.TotalPrice,
			Products:   []*pb.Order_OrderProduct{},
		}
		op.CreatedAt,_ = o.CreatedAt.MarshalBinary()
		//decorate orders with products
		for _,product := range o.Products{
			//populate product fields
			for _,p := range products{
				if p.ID == product.ID{
					product.Name = p.Name
					product.Description = p.Description
					product.Price = p.Price
					break
				}
			}
			op.Products = append(op.Products, &pb.Order_OrderProduct{
				Id:          product.ID,
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Quantity:    product.Quantity,
			})
		}
		orders = append(orders, op)
	}
	return &pb.GetOrdersForAccountResponse{Orders:orders},nil
}