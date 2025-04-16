package main

import "fmt"

func main() {
	// 创建一个存储商品和价格的map
	inventory := map[string]map[string]float64{
		"电子产品": {
			"手机": 1999.99,
			"电脑": 5999.99,
			"平板": 2399.99,
		},
		"书籍": {
			"Go程序设计": 79.8,
			"算法导论":   108.0,
			"数据结构":   56.5,
		},
	}

	// 打印所有商品
	fmt.Println("商品库存:")
	for category, products := range inventory {
		fmt.Printf("%s:\n", category)
		for name, price := range products {
			fmt.Printf("  - %s: ¥%.2f\n", name, price)
		}
	}

	// 添加新分类和商品
	inventory["服装"] = map[string]float64{
		"T恤":  99.9,
		"牛仔裤": 199.9,
	}

	// 向已有分类添加新商品
	inventory["电子产品"]["耳机"] = 299.9

	// 修改价格
	inventory["书籍"]["Go程序设计"] = 69.9

	fmt.Println("\n更新后的商品库存:")
	for category, products := range inventory {
		fmt.Printf("%s:\n", category)
		for name, price := range products {
			fmt.Printf("  - %s: ¥%.2f\n", name, price)
		}
	}
}
