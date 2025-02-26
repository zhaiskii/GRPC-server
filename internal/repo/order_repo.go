package repo 

import(
	"fmt"
	//"time"
	"sync"
	"github.com/google/uuid"
)

type Order struct{
	ID string
	Item string
	Quantity int32
}

type OrderRepo struct{
	mu sync.Mutex
	orders map[string]Order
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{orders: make(map[string]Order)}
}

func (r *OrderRepo) CreateOrder(item string, quantity int32) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	//currentTimeStamp := time.Now().UnixNano() / int64(time.Microsecond)
	uuid := uuid.New().ID()
	var id string = ""
	for ; uuid>0; uuid/=10 {
		id += string((uuid%10)+'0')
	}
	fmt.Println(uuid)
	fmt.Println(id)
	r.orders[id] = Order{ID: id, Item: item, Quantity: quantity}
	fmt.Println(r.orders[id].Quantity)
	return id
}

func (r *OrderRepo) GetOrder(id string) (*Order, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lolkek, err := r.orders[id]
	//fmt.Println(r.orders[id].Quantity)
	return &lolkek, err
}

func (r *OrderRepo) UpdateOrder(id string, item string, quantity int32) (*Order, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, err := r.orders[id]
	if !err {
		return nil, err
	}
	r.orders[id] = Order{ID: id, Item: item, Quantity: quantity}
	lolkek := r.orders[id]
	return &lolkek, err
}

func (r *OrderRepo) DeleteOrder(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, err := r.orders[id]
	if !err {
		return false
	}
	delete(r.orders, id)
	return true
}

func (r *OrderRepo) ListOrders() []*Order {
	r.mu.Lock()
	defer r.mu.Unlock()
	var res []*Order
	for _, v := range r.orders {
		lol := v
		res = append(res, &lol)
	}
	return res
}