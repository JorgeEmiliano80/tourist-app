import React, { useState } from 'react';
import { TextField, Button, Grid } from '@material-ui/core';
import { createPayment } from '../src/services/paymentService';

const PaymentForm = ({ bookingId, amount, onPaymentComplete }) => {
    const [paymentData, setPaymentData] = useState({
        cardNumber: '',
        expirationsDate: '',
        cvv: '',
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setPaymentData(prevState => ({ ...prevState, [name]: value }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await createPayment({ ...paymentData, bookingId, amount });
            onPaymentComplete();
        } catch (error) {
            console.error('Error processing payment:', error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <Grid container spacing={2}>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        label="Card Number"
                        name="cardNumber"
                        value={paymentData.cardNumber}
                        onChange={handleChange}
                    />
                </Grid>
                <Grid item xs={6}>
                    <TextField
                        fullWidth
                        label="Expiration Date"
                        name="expirationDate"
                        value={paymentData.expirationDate}
                        onChange={handleChange}
                    />
                </Grid>
                <Grid item xs={6}>
                    <TextField
                        fullWidth
                        label="CVV"
                        name="cvv"
                        value={paymentData.cvv}
                        onChange={handleChange}
                    />
                </Grid>
                <Grid item xs={12}>
                    <Button type="submit" variant="contained" color="primary">
                        Pay ${amount}
                    </Button>
                </Grid>
            </Grid>
        </form>
    );
};

export default PaymentForm;