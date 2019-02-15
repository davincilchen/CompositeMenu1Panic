package composite

import (
	"testing"
)

func TestComponent(t *testing.T) {

	menu1 := NewMenu("菲力6oz套餐", "供應時間 11:00--22:00")
	menu1.Add(NewMenuItem("主食", "Prime菲力6oz", 600))
	menu1.Add(NewMenuItem("前菜", "玉米沙拉", 35))
	menu1.Add(NewMenuItem("前菜", "水果沙拉", 35))
	menu1.Add(NewMenuItem("飲料(2選1)", "柳橙汁", 30))
	menu1.Add(NewMenuItem("飲料(2選1)", "蘋果汁", 30))
	menu1.Add(NewMenuItem("折價", "前菜折價", -35))
	menu1.Add(NewMenuItem("折價", "飲料折價", -30))
	menu1.Add(NewMenuItem("折價", "套餐折價", -10))

	menu2 := NewMenu("紐澳良烤雞腿飯套餐", "供應時間：11:00--22:00")
	menu2.Add(NewMenuItem("主食", "紐澳良烤雞腿飯", 545))
	menu2.Add(NewMenuItem("前菜", "玉米沙拉", 35))
	menu2.Add(NewMenuItem("前菜", "水果沙拉", 35))
	menu2.Add(NewMenuItem("飲料(2選1)", "柳橙汁", 30))
	menu2.Add(NewMenuItem("飲料(2選1)", "蘋果汁", 30))
	menu2.Add(NewMenuItem("折價", "前菜折價", -35))
	menu2.Add(NewMenuItem("折價", "飲料折價", -30))
	menu2.Add(NewMenuItem("折價", "套餐折價", -10))

	menu3 := NewMenu("雙人套餐總折價", "")
	menu3.Add(NewMenuItem("折價內容", "雙人套折價", -40))

	all := NewMenu("超值雙人午餐套餐", "周一至周五販賣 供應時間：11:00--14:00")
	all.Add(menu1)
	all.Add(menu2)
	all.Add(menu3)

	all.Print()

}
