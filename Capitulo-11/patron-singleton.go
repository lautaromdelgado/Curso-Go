/*
	Video: https://www.youtube.com/watch?v=5zQmFe2ijAg&ab_channel=RobertoMorais
*/

package singleton

import(
	"sync"
)

type singleton struct {}

var instance *singleton

var once sync.Once

func GetInstance() *singleton {
	once.Do(func(){
		instance = &singleton{}
	})

	return instance
}