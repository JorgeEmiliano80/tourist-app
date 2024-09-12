import React, { useState } from 'react';
import { TextField, Button, Grid, Box } from '@mui/material';
import { createPayment } from '../services/paymentService';

const PaymentForm = ({ bookingId, amount, onPaymentComplete }) => {
    const [paymentData, setPaymentData] = useState({
        cardNumber: '',
        expirationDate: '',
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
        <Box component="form" onSubmit={handleSubmit} sx={{ mt: 3 }}>
            <Grid container spacing={2}>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        label="Card Number"
                        name="cardNumber"
                        value={paymentData.cardNumber}
                        onChange={handleChange}
                        variant="outlined"
                    />
                </Grid>
                <Grid item xs={6}>
                    <TextField
                        fullWidth
                        label="Expiration Date"
                        name="expirationDate"
                        value={paymentData.expirationDate}
                        onChange={handleChange}
                        variant="outlined"
                    />
                </Grid>
                <Grid item xs={6}>
                    <TextField
                        fullWidth
                        label="CVV"
                        name="cvv"
                        value={paymentData.cvv}
                        onChange={handleChange}
                        variant="outlined"
                    />
                </Grid>
                <Grid item xs={12}>
                    <Button 
                        type="submit" 
                        variant="contained" 
                        color="primary"
                        fullWidth
                    >
                        Pay ${amount}
                    </Button>
                </Grid>
            </Grid>
        </Box>
    );
};

export default PaymentForm;