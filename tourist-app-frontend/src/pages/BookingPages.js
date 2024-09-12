import React, { useState, useEffect } from 'react';
import { Container, Typography, List, ListItem, ListItemText, Button, Stack } from '@mui/material';
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
            console.error('Error fetching bookings: ', error);
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
                    <ListItem
                        key={booking.id}
                        secondaryAction={
                            <Button
                                variant="contained"
                                color="error"
                                onClick={() => handleDeleteBooking(booking.id)}
                            >
                                Cancel Booking
                            </Button>
                        }
                    >
                        <ListItemText
                            primary={booking.tour.name}
                            secondary={
                                <React.Fragment>
                                    <Typography component="span" variant="body2" color="text.primary">
                                        Date: {booking.date}
                                    </Typography>
                                    {" — People: " + booking.numberOfPeople}
                                </React.Fragment>
                            }
                        />
                    </ListItem>
                ))}
            </List>
        </Container>
    );
};

export default BookingPage;
