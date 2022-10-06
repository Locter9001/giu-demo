package main

import (
	"github.com/AllenDang/giu"
	"time"
)

var rows []*giu.TableRowWidget
var personList []person

type person struct {
	name int
	age  int
}

func loop() {
	giu.SingleWindowWithMenuBar().Layout(
		giu.Table().Columns(
			giu.TableColumn("name"),
			giu.TableColumn("age"),
		).Rows(
			rows...,
		),
	)
}

func main() {
	w := giu.NewMasterWindow("Overview", 200, 500, 0)
	go refresh()
	w.Run(loop)
}

func refresh() {
	ticker := time.NewTicker(time.Millisecond * 100)
	name := 10
	age := 18
	for {
		var person = person{
			name: name + 1,
			age:  age + 1,
		}

		personList = append(personList, person)

		var rowList []*giu.TableRowWidget
		for i := 0; i < len(personList); i++ {
			rowList = append(rowList, giu.TableRow(
				giu.Selectablef("%v", personList[i].name),
				giu.Selectablef("%v", personList[i].age),
			))
		}
		name++
		age++
		rows = rowList

		giu.Update()

		<-ticker.C
	}
}
