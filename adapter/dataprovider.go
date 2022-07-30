package adapter

const XML_DATA = `<?xml version="1.0" encoding="UTF-8" ?>
<ProductList>
    <Product>
        <sku>ABC123</sku>
        <quantity>2</quantity>
    </Product>
    <Product>
        <sku>ABC321</sku>
        <quantity>5</quantity>
    </Product>
</ProductList>`

type DataProvider struct{} // Источник данных в формате xml

func (w DataProvider) GetDataXML() string {
	return XML_DATA
}
