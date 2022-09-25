package etcd




//Self Discovery

/*

type discovery interface {
	Name() string
	Do(ctx context.Context, prefix string)
	GetInstance(ctx context.Context) *service_mng.Instance
}

type ETCDDiscovery struct {
	//loadBalancer is a
	service_mng.loadBalancer
	Instances map[string]*service_mng.Instance
	mu        sync.RWMutex
	cli       *clientv3.Client
	ttl       int64
	timeout   int64
}

func (d *ETCDDiscovery) Name() string {
	resolver.GetDefaultScheme()
	return "ETCD_DISCOVERY"
}

func NewETCDDiscovery(timeout int64) (*ETCDDiscovery, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 15 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	d := &ETCDDiscovery{
		loadBalancer: &service_mng.RandomLoadBalance{},
		cli:          cli,
		timeout:      timeout,
		Instances:    make(map[string]*service_mng.Instance, 16),
	}

	return d, nil
}

func (d *ETCDDiscovery) GetInstance(ctx context.Context) *service_mng.Instance {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.loadBalancer.Do(ctx, d.Instances)
}

//service_mng discovery
func (d *ETCDDiscovery) Do(ctx context.Context, prefix string) {
	kv := clientv3.NewKV(d.cli)
	resp, err := kv.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		slog.Panic("ETCDDiscovery: get kv err, ", err.Error())
	}

	for _, kv := range resp.Kvs {
		ins := &service_mng.Instance{}
		if err := json.Unmarshal(kv.Value, ins); err != nil {
			slog.Warn("ETCDDiscovery: json unmarshal err, ", err.Error())
			continue
		}
		d.Instances[string(kv.Key)] = ins
	}

	//watch service_mng change
	go d.watchService(prefix)
}

func (d *ETCDDiscovery) watchService(prefix string) {
	watcher := clientv3.NewWatcher(d.cli)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(d.timeout)*time.Second)
	watchChan := watcher.Watch(ctx, prefix, clientv3.WithPrefix())

	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			key := string(event.Kv.Key)
			value := string(event.Kv.Value)
			slog.Debugf("ETCDDiscovery: event, type:%s key:%d value:%d", event.Type.String(), key, value)
			switch event.Type {
			case mvccpb.PUT:
				d.addInstance(key, value)
			case mvccpb.DELETE:
				d.delInstance(string(event.Kv.Key))
			}
		}
	}
}

func (d *ETCDDiscovery) addInstance(key string, value string) {
	newIns := &service_mng.Instance{}
	if err := json.Unmarshal([]byte(value), newIns); err != nil {
		slog.Warn("ETCDDiscovery: json unmarshal err, ", err.Error())
		return
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	conn, err := grpc.Dial(newIns.Endpoint, grpc.WithInsecure())
	if err != nil {
		slog.Warn("ETCDDiscovery addInstance: grcp dial err, ", err.Error())
		return
	}
	newIns.clientConn = conn
	d.Instances[key] = newIns
}

func (d *ETCDDiscovery) delInstance(key string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	//close grpc connection
	d.Instances[key].GetGRPCConn().Close()
	delete(d.Instances, key)
}

*/
