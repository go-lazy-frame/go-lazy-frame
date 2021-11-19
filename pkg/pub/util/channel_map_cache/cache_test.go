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
 * long.qian 2021-11-19 10:21 创建
 */

/**
 * @author long.qian
 */

package channel_map_cache

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"
)

type V struct {
	Name string
}

func TestCache(t *testing.T) {
	cache := ChannelMapCache{KeyLogic: keyLogic}
	k := uuid.New().String()
	cache.Put(V{Name: k})
	cache.Put(V{Name: time.Now().String()})
	cache.Put(V{Name: time.Now().String()})
	cache.Put(V{Name: time.Now().String()})
	cache.Put(V{Name: time.Now().String()})
	fmt.Println(cache.GetAll())
	fmt.Println(cache.GetByKey(k))

	cache.Clear()
	fmt.Println(cache.GetAll())
}

func keyLogic(obj interface{}) string {
	return uuid.New().String()
}
