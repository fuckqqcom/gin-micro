consul_api目录中准备重写服务注册和发现(consul)

根据源码:
    
    	reg := consul.NewRegistry(func(op *registry.Options) {
    		op.Addrs = []string{
    			"192.168.111.129:8500",
    		}
    	})
    	
    	
    	func NewRegistry(opts ...registry.Option) registry.Registry {
        	return registry.NewRegistry(opts...)
        }
        
        func NewRegistry(opts ...Option) Registry {
        	return newConsulRegistry(opts...)
        }
        
        
        func newConsulRegistry(opts ...Option) Registry {
        	cr := &consulRegistry{
        		opts:        Options{},
        		register:    make(map[string]uint64),
        		lastChecked: make(map[string]time.Time),
        		queryOptions: &consul.QueryOptions{
        			AllowStale: true,
        		},
        	}
        	configure(cr, opts...)
        	return cr
        }
        
        
        也就是NewRegistry--->newConsulRegistry--->
        
        type Server interface {
        	Options() Options
        	Init(...Option) error
        	Handle(Handler) error
        	NewHandler(interface{}, ...HandlerOption) Handler
        	NewSubscriber(string, interface{}, ...SubscriberOption) Subscriber
        	Subscribe(Subscriber) error
        	Register() error
        	Deregister() error
        	Start() error
        	Stop() error
        	String() string
        }
        
        但是还是有点迷糊,就算把这些重写了,go-micro底层还调用到自己实现的这部分吗
        
        
        
        自己实现的话
        client.Agent().ServiceRegister(reg)
        重写有实现这部分代码的函数就可以了理论上