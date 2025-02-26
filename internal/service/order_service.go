package service

import(
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "yalyceum_hw_grpc/pkg/api/test/api" //maybe test
	"yalyceum_hw_grpc/internal/repo"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
	repo *repo.OrderRepo
}

func NewOrderServiceServer(r *repo.OrderRepo) *OrderServiceServer {
	return &OrderServiceServer{repo: r}
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	orderID := s.repo.CreateOrder(req.Item, req.Quantity)

	return &pb.CreateOrderResponse{Id: orderID}, nil
}

func (s *OrderServiceServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	res, exist := s.repo.GetOrder(req.Id) // how the fuck can I find this fuck

	if !exist {
		return nil, status.Errorf(codes.NotFound, "order not found while getting")
	}
	return &pb.GetOrderResponse{
		Order: &pb.Order{Id: res.ID,
		Item: res.Item,
		Quantity: res.Quantity,},
	}, nil
}

func (s *OrderServiceServer) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	res, exist:= s.repo.UpdateOrder(req.Id, req.Item, req.Quantity)
	if !exist {
		return nil, status.Errorf(codes.NotFound, "order not found while updating")
	}
	return &pb.UpdateOrderResponse{Order: &pb.Order{Id: res.ID, Item: res.Item, Quantity: res.Quantity}}, nil
}

func (s *OrderServiceServer) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	nice := s.repo.DeleteOrder(req.Id)
	if !nice {
		return &pb.DeleteOrderResponse{Success: nice}, status.Errorf(codes.NotFound, "order not found while deleting")
	}
	return &pb.DeleteOrderResponse{Success: nice}, nil
}

func (s *OrderServiceServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	kek := s.repo.ListOrders()
	var dodik []*pb.Order
	for _, v := range kek {
		var lol *pb.Order = &pb.Order{Id: v.ID, Item: v.Item, Quantity: v.Quantity}
		dodik = append(dodik, lol)
	}
	return &pb.ListOrdersResponse{Orders: dodik}, nil
}