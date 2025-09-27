
packge order 


type RedisRepo struct { 
Client *redist.Client 
} 

func orderIDKey(id unit64) string{ 
fmt.Springf("order:%d", id) 
} 

func (r *RedisRepo) Insert(ctx context.Context, order model.Order) error{ 
data , err := json.Marshal(order) 
if err != nil{ 
		fmt.Errorf("filed to encode order: %w", err) 
	} 
   key := orderIDKey(order.OrderID)
   res := r.Client.SetNX(ctx, key, string(data), 0)
   if err := res.Err(); err != nil{ 
	return fmt.Errorf("faield to set: %w", err)
   }
   
return nil
} 
errNotExist := errors.New("order does nto exist")
func (r *RedisRepo) findById(ctx cntext.Context, order model.Order) (model.Order, error){
   key := orderIDKey(order.OrderID)
   value, err := r.Client.Get(ctx, key)Result()
   if errors.Is(err, redist,nil){ 
	retur model.Order{}, errNotExist
   }else if err !=nil { 
	return model.Order{} , fmt.Errorf("get order: %w", err)
   }

   var order model.Order 
   err =: json.Unmarshal([]byte(value), &order)
    if err !=nil { 
	return model.Order{} , fmt.Errorf("failed to decode order json: %w", err)
   }

 return order , nil
   


}
