import React, { useState, useEffect } from 'react';
import { Container, Typography, List, ListItem, ListItemText } from '@material-ui/core';
import { getPayments } from '../services/paymentService';

const PaymentPage = () => {
    const [payments, setPayments] = useState([]);
    const { user } = useAuthContext();

    useEffect(() => {
        fetchPayments();
    }, []);

    const fetchPayments = async () => {
        try {
            const response = await getPayments();
            setPayments(response.data);
        } catch (error) {
            console.error('Error fetching payments: ', error);
        }
    };

    return (
        <Container>
            <Typography variant="h2" gutterBotom>
                Payment History
            </Typography>
            <List>
                {payments.map((payment) => (
                    <ListItem key={payment.id}>
                        <ListItemText
                            primary={`Payment for ${payment.booking.tour.name}`}
                            secondary={`Amount: $${payment.amount}, Date: ${new Date(payment.date).toLocaleDateString()}`}
                        />
                    </ListItem>
                ))}
            </List>
        </Container>
    );
};

export default PaymentPage;
