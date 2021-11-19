//             ,%%%%%%%%,
//           ,%%/\%%%%/\%%
//          ,%%%\c "" J/%%%
// %.       %%%%/ o  o \%%%
// `%%.     %%%%    _  |%%%
//  `%%     `%%%%(__Y__)%%'
//  //       ;%%%%`\-/%%%'
// ((       /  `%%%%%%%'
//  \\    .'          |
//   \\  /       \  | |
//    \\/攻城狮保佑) | |
//     \         /_ | |__
//     (___________)))))))                   `\/'
/*
 * 修订记录:
 * long.qian 2021-11-19 09:39 创建
 */

/**
 * @author long.qian
 */

package channel_map_cache

// ChannelMapCache 高性能基于 channel 的缓存
// 示例用法：
// cache = channel_map_cache.ChannelMapCache{
//		// 指定缓存存储的 Key 策略
//		KeyLogic: func(obj interface{}) string {
//			u := obj.(model.User)
//			return fmt.Sprintf("%s", u.LoginName)
//		},
//	}
// cache.Put()
type ChannelMapCache struct {
	init                  bool
	KeyLogic              func(interface{}) string // 缓存存储的 Key 策略
	cache                 map[string]interface{}
	cacheChWrite          chan interface{}
	cacheChDelByKey       chan string
	cacheChDelByKeyLogic  chan interface{}
	cacheChReadNotice     chan byte
	cacheChRead           chan map[string]interface{}
	cacheChReadByKeyValue chan interface{}
	cacheChReadByKey      chan string
	cacheChClear          chan byte
}

// GetAll 获取所有的缓存
func (receiver *ChannelMapCache) GetAll() map[string]interface{} {
	if !receiver.init {
		receiver.cacheInit()
	}
	receiver.cacheChReadNotice <- 0
	cache := <-receiver.cacheChRead
	return cache
}

// GetByKey 获取指定 key 的缓存数据
func (receiver *ChannelMapCache) GetByKey(key string) interface{} {
	if !receiver.init {
		receiver.cacheInit()
	}
	receiver.cacheChReadByKey <- key
	value := <-receiver.cacheChReadByKeyValue
	return value
}

// Put 添加缓存
func (receiver *ChannelMapCache) Put(value interface{}) {
	if !receiver.init {
		receiver.cacheInit()
	}
	receiver.cacheChWrite <- value
}

// RemoveByKey 删除指定 Key 的缓存
func (receiver *ChannelMapCache) RemoveByKey(key string) {
	if !receiver.init {
		receiver.cacheInit()
	}
	receiver.cacheChDelByKey <- key
}

// RemoveByKeys 删除指定 Key 的缓存
func (receiver *ChannelMapCache) RemoveByKeys(keys ...string) {
	if !receiver.init {
		receiver.cacheInit()
	}
	for _, s := range keys {
		receiver.cacheChDelByKey <- s
	}
}

// RemoveByKeyLogic 根据指定 Value 删除缓存（将根据指定的缓存生存策略进行缓存的删除）
func (receiver *ChannelMapCache) RemoveByKeyLogic(value interface{}) {
	if !receiver.init {
		receiver.cacheInit()
	}
	receiver.cacheChDelByKeyLogic <- value
}

// Clear 清空缓存
func (receiver *ChannelMapCache) Clear() {
	receiver.cacheChClear <- 0
}

func (receiver *ChannelMapCache) cacheInit() {
	receiver.init = true
	receiver.cache = make(map[string]interface{})
	receiver.cacheChWrite = make(chan interface{})
	receiver.cacheChDelByKey = make(chan string)
	receiver.cacheChDelByKeyLogic = make(chan interface{})
	receiver.cacheChReadNotice = make(chan byte)
	receiver.cacheChRead = make(chan map[string]interface{})
	receiver.cacheChReadByKeyValue = make(chan interface{})
	receiver.cacheChReadByKey = make(chan string)
	receiver.cacheChClear = make(chan byte)
	go func() {
		for {
			select {
			case data := <-receiver.cacheChWrite:
				if receiver.KeyLogic == nil {
					panic("缓存 Key 逻辑规则未指定")
				}
				key := receiver.KeyLogic(data)
				receiver.cache[key] = data
			case data := <-receiver.cacheChDelByKey:
				delete(receiver.cache, data)
			case data := <-receiver.cacheChDelByKeyLogic:
				if receiver.KeyLogic == nil {
					panic("缓存 Key 逻辑规则未指定")
				}
				key := receiver.KeyLogic(data)
				delete(receiver.cache, key)
			case <-receiver.cacheChReadNotice:
				var tmp = make(map[string]interface{})
				for k, v := range receiver.cache {
					tmp[k] = v
				}
				receiver.cacheChRead <- tmp
			case key := <-receiver.cacheChReadByKey:
				receiver.cacheChReadByKeyValue <- receiver.cache[key]
			case <-receiver.cacheChClear:
				receiver.cache = make(map[string]interface{})
			}
		}
	}()
}
