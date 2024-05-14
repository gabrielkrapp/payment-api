package worker

import "fmt"

func RenderPaymentEmail(customerEmail, paymentID, amount string) (subject, body string) {
	subject = "Confirmação de Pagamento"
	body = fmt.Sprintf(`
        Prezado(a),

        Recebemos o seu pagamento com sucesso.

        Detalhes do pagamento:
        - ID: %s
        - Valor: %s

        Agradecemos a sua preferência.

        Atenciosamente,
        Equipe Financeira
    `, paymentID, amount)

	return subject, body
}
