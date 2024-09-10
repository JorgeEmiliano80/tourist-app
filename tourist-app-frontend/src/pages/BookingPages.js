import React, { useState, useEffect } from 'react';
import { Container, Typography, List, ListItem, ListItemText, Button } from '@material-ui/core';
import { getBookings, deleteBooking } from '../services/bookingService';
import { useAuthContext } from '../context/AuthContext';

const BookingPage = () => {
    const [bookings, setBookings] = useState([]);
    const { user } = useAuthContext();

    useEffect(() => {
        fetchBookings();
    }, []);

    const fetchBookings = async () => {
        try {
            const response = await getBookings();
            setBookings(response.data);
        } catch (error) {
            console.error('Errorr fetching bookings: ', error);
        }
    };

    const handleDeleteBooking = async (id) => {
        try {
            await deleteBooking(id);
            fetchBookings();
        } catch (error) {
            console.error('Error deleting booking: ', error);
        }
    };

    return (
        <Container>
            <Typography variant="h2" gutterBottom>
                Your Bookings
            </Typography>
            <List>
                {bookings.map((booking) => (
                    <ListItem key={booking.id}>
                        <ListItemText
                            primary={booking.tour.name}
                            secondary={`Date: ${booking.date}, People: ${booking.numberOfPeople}`}
                        />
                        <Button
                            variant="contained"
                            color="secondary"
                            onClick={() => handleDeleteBooking(booking.id)}
                        >
                            Cancel Booking
                        </Button>
                    </ListItem>
                ))}
            </List>
        </Container>
    );
};

export default BookingPage;
