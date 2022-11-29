## bubbly

Bubbly is a library written in go that use gographviz to create diagrams by just writing it and not wasting your time by dragging lines and reposition every time by hand all your schema at every modify.


Star schema:


```go
func main() {

	d := diagram.NewSTR()

	SaleFT := d.CreateDimension("SaleFT", "StoreID DateID ProductID", "quantity Receipts")
	
	WeekDT := d.CreateDimension("WeekDT", "DateID", "Date Month")
	StoreDT := d.CreateDimension("StoreDT", "StoreID", "Store City Country SalesManager")
	ProductDT := d.CreateDimension("ProductDT", "ProductID", "Product Type Category")

	d.JoinDimension(SaleFT, WeekDT, "DateID")
	d.JoinDimension(SaleFT, StoreDT, "StoreID")
	d.JoinDimension(SaleFT, ProductDT, "ProductID")

	d.RenderDiagram()
}
```

<img src="str_schema_look.png" width="50%" height="50%">


what bubbly needs to be complete:

DFM:
- [x] convergence
- [x] descriptive attributes
- [x] add fact
- [x] add node   
- [x] optional arc
- [x] hierarchies
- [ ] optional convergence arc
- [ ] optional arc
- [ ] role
- [ ] cycle

STR:
- [x] fact and table
- [x] colors

other diagrams:
- [ ] nn
- [ ] er

optional objectives:
- [ ] parser
- [ ] UI
