document.addEventListener('DOMContentLoaded', function() {
    var stripe = Stripe('public_key');
    var elements = stripe.elements();
    var style = {
        base: {
            color: "#32325d",
        }
    };

    var card = elements.create('card', { style: style });
    card.mount('#card-element');

    card.addEventListener('change', function(event) {
        var displayError = document.getElementById('card-errors');
        if (event.error) {
            displayError.textContent = event.error.message;
        } else {
            displayError.textContent = '';
        }
    });

    var form = document.getElementById('payment-form');
    form.addEventListener('submit', function(event) {
        event.preventDefault();

        var amount = document.getElementById('amount').value;
        var currency = document.getElementById('currency').value;

        stripe.createPaymentMethod({
            type: 'card',
            card: card
        }).then(function(result) {
            if (result.error) {
                var errorElement = document.getElementById('card-errors');
                errorElement.textContent = result.error.message;
            } else {
                fetch('/v1/payment', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        paymentMethodId: result.paymentMethod.id,
                        amount: parseInt(amount),
                        currency: currency.toUpperCase()
                    })
                }).then(function(response) {
                    response.json().then(function(json) {
                        console.log(json);
                    });
                });
            }
        });
    });
});
