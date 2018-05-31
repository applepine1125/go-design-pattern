package factory_method

type Product interface {
	Use() string
}

type Factory struct {
}

type creater interface {
	createProduct(owner string) Product
	registerProduct(product Product)
}

func (f *Factory) Create(c creater, owner string) Product {
	p := c.createProduct(owner)
	c.registerProduct(p)
	return p
}

type IDCard struct {
	owner string
}

func (ic *IDCard) Use() string {
	return ic.owner
}

type IDCardFactory struct {
	*Factory
	owners []*string
}

func (icf *IDCardFactory) createProduct(owner string) Product {
	return &IDCard{owner: owner}
}

func (icf *IDCardFactory) registerProduct(product Product) {
	owner := product.Use()
	icf.owners = append(icf.owners, &owner)
}

func (icf *IDCardFactory) getOwners() []*string {
	return icf.owners
}
