package gossdb

import (
	"errors"
	"github.com/jolestar/go-commons-pool"
	"github.com/stitchcula/gossdb/conf"
)

//连接池
type Connectors struct {
	pool *pool.ObjectPool //连接池
	cfg  conf.Config      //配置
}

//用配置文件进行初始化
//
//  cfg 配置文件
func (c *Connectors) Init(cfg *conf.Config) {
	c.pool = pool.NewObjectPoolWithDefaultConfig(&clientFactory{
		MakeNew: func() (*SSDBClient, error) {
			return &SSDBClient{
				Host:     cfg.Host,
				Port:     cfg.Port,
				Password: cfg.Password,
			}, nil
		},
	})

	c.pool.Config.MaxTotal = cfg.MaxPoolSize
	c.pool.Config.MinIdle = cfg.MinPoolSize
}

// 测试&Ping
func (c *Connectors) Start() error {
	cc, err := c.NewClient()
	if err != nil {
		return err
	}
	_, err = cc.Info()
	return err
}

//关闭连接池
func (c *Connectors) Close() {
	c.pool.Close()
}

// GET
func (c *Connectors) NewClient() (*Client, error) {
	Object, err := c.pool.BorrowObject()
	if err != nil {
		return nil, err
	}
	ssdbClient, OK := Object.(*SSDBClient)
	if !OK {
		return nil, errors.New("PooledObject.Object Type error")
	}
	return &Client{
		db:   ssdbClient,
		pool: c,
	}, nil
}

// CLOSE
func (c *Connectors) closeClient(cc *Client) {
	// 回收ssdbClient，并抛弃Client
	c.pool.ReturnObject(cc.db)
}

/*----- factory BEGIN -----*/
type clientFactory struct {
	MakeNew func() (*SSDBClient, error)
}

func (f *clientFactory) MakeObject() (*pool.PooledObject, error) {
	client, _ := f.MakeNew()
	err := client.Start()
	return pool.NewPooledObject(client), err
}

func (f *clientFactory) DestroyObject(obj *pool.PooledObject) error {
	client, OK := obj.Object.(*SSDBClient)
	if !OK {
		return errors.New("PooledObject.Object Type error")
	}
	if client != nil {
		client.Close()
	}
	return nil
}

func (f *clientFactory) ValidateObject(obj *pool.PooledObject) bool {
	if obj.Object == nil {
		return false
	}
	client, OK := obj.Object.(*SSDBClient)
	if !OK {
		return false
	}
	return client.Ping()
}

func (f *clientFactory) ActivateObject(object *pool.PooledObject) error {
	// todo: do activate
	return nil
}

func (f *clientFactory) PassivateObject(object *pool.PooledObject) error {
	// todo: do passivate
	return nil
}

/*----- factory END -----*/
