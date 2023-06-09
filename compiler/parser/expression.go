package parser

type Expression interface {
}

type BlockExpression struct {
	Open Token
	Close Token
}

func parseExpression(source *[]string) Expression {
	var expression BlockExpression
	
	expression.Open = parseToken(source, "{")
	expression.Close = parseToken(source, "}")

	return expression
}

// let [id] = [ex]