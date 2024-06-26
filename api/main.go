package main

import (
    "net/http"
    "strconv"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    // Defino o objeto do Gin
    g := gin.Default()

    // Configurar o CORS
    g.Use(cors.Default())

    // Rota para um endpoint
    g.GET("/", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "message": "Hello World",
        })
    })

    // Calcular a média dos dividendos dos últimos 5 anos
    g.GET("/calcular", func(c *gin.Context) {
        soma := 0.0
        for i := 1; i <= 5; i++ {
            chave := "valor" + strconv.Itoa(i)
            valorStr := c.Query(chave)
            valor, err := strconv.ParseFloat(valorStr, 64)
            if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Todos os valores devem ser números válidos"})
                return
            }
            soma += valor
        }
        rentabilidadeStr := c.Query("rentabilidade")
        rentabilidade, err := strconv.ParseFloat(rentabilidadeStr, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Rentabilidade -> valor inválido"})
            return
        }
        cotacaoStr := c.Query("cotacao")
        cotacao, err := strconv.ParseFloat(cotacaoStr, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Cotação -> valor inválido"})
            return
        }

        media := soma / 5
        precoTeto := media / (rentabilidade / 100)

        if precoTeto >= cotacao {
            c.JSON(http.StatusOK, gin.H{"preço-teto": precoTeto, "Média de dividendos": media, "oportunidade": true})
        } else {
            c.JSON(http.StatusOK, gin.H{"preço-teto": precoTeto, "Média de dividendos": media, "oportunidade": false})
        }
    })

    g.Run(":3000")
}
