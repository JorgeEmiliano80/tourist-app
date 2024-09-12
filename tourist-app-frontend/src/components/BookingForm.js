import React, { useState } from 'react';
import { TextField, Button, Grid } from '@mui/material';
import { createBooking } from '../services/bookingService';

const BookingForm = ({ tourId, onBookingComplete }) => {
    const [bookingData, setBookingData] = useState({
        date: '',
        numberOfPeople: '',
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setBookingData(prevState => ({ ...prevState, [name]: value }));
    };
    
    const handleSubmit = async (e) => {
        e.preventDefault(); // Corregido de prevenDefault a preventDefault
        try {
            await createBooking({ ...bookingData, tourId });
            onBookingComplete();
        } catch (error) {
            console.error('Error creating booking: ', error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <Grid container spacing={2}>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        label="Date"
                        name="date"
                        type="date"
                        value={bookingData.date}
                        onChange={handleChange}
                        InputLabelProps={{
                            shrink: true,
                        }}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        label="Number of People"
                        name="numberOfPeople"
                        type="number"
                        value={bookingData.numberOfPeople}
                        onChange={handleChange}
                    />
                </Grid>
                <Grid item xs={12}>
                    <Button type="submit" variant="contained" color="primary">
                        Book Now
                    </Button>
                </Grid>
            </Grid>
        </form>
    );
};

export default BookingForm;
